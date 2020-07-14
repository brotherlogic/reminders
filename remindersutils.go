package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/net/context"

	pbgh "github.com/brotherlogic/githubcard/proto"
	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/reminders/proto"
)

func (s *Server) refresh(ctx context.Context) {
	complete := false
	for _, tl := range s.data.GetTasks() {
		s.Log(fmt.Sprintf("Processing %v", tl.GetName()))
		complete = complete || s.processTaskList(ctx, tl)
	}
}

func (s *Server) processTaskList(ctx context.Context, t *pb.TaskList) bool {
	for _, task := range t.Tasks.Reminders {

		//Reassign a task with an empty id
		if task.GetGithubId() == "" {
			task.CurrentState = pb.Reminder_UNASSIGNED
		}

		switch task.GetCurrentState() {
		case pb.Reminder_UNASSIGNED:
			t, err := s.ghbridge.addIssue(ctx, task)
			if err == nil {
				for _, sil := range task.Silences {
					err2 := s.silence.addSilence(ctx, sil, fmt.Sprintf("%v", task.Uid))
					if err2 != nil {
						s.pushFail++
						s.pushFailure = fmt.Sprintf("silence - %v", err)
						s.Log(fmt.Sprintf("Error adding silence %v", err2))
						return true
					}
				}
				task.CurrentState = pb.Reminder_ASSIGNED
				s.pushCount++
				task.GithubId = t
				s.last = &pbgh.Issue{Service: task.GithubId}
				//				s.save(ctx)
			} else {
				s.pushFail++
				s.pushFailure = fmt.Sprintf("add issue - %v", err)
				s.Log(fmt.Sprintf("Error adding issue %v", err))
			}

			return true
		case pb.Reminder_ASSIGNED:
			if s.ghbridge.isComplete(ctx, task) {
				err := s.silence.removeSilence(ctx, fmt.Sprintf("%v", task.Uid))
				if err != nil {
					s.Log(fmt.Sprintf("Unable to unsilence"))
				}
				task.CurrentState = pb.Reminder_COMPLETE
			} else {
				return true
			}
		}
	}

	return false
}

func (s *Server) getReminders(t time.Time) []*pb.Reminder {
	reminders := make([]*pb.Reminder, 0)

	if s.data.List != nil && s.data.List.Reminders != nil {
		for _, r := range s.data.List.Reminders {
			if r.NextRunTime < t.Unix() {
				s.adjustRunTime(r)
				s.Log(fmt.Sprintf("Adjusted %v", r.Text))
				reminders = append(reminders, r)
			}
		}
	}

	return reminders
}

func (s *Server) adjustRunTime(r *pb.Reminder) {
	s.Log(fmt.Sprintf("Adjusting for %v", r.Text))
	t := time.Now()
	ct := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	_, week := ct.ISOWeek()

	if r.RepeatPeriodInSeconds > 0 {
		r.NextRunTime = t.Unix() + r.RepeatPeriodInSeconds
	} else {
		switch r.RepeatPeriod {
		case pb.Reminder_DAILY:
			ct = ct.AddDate(0, 0, 1)
		case pb.Reminder_WEEKLY:
			for (r.DayOfWeek != "" && ct.Weekday().String() != r.DayOfWeek) || ct.Before(t) {
				ct = ct.AddDate(0, 0, 1)
			}
		case pb.Reminder_BIWEEKLY:
			for (r.DayOfWeek != "" && ct.Weekday().String() != r.DayOfWeek) || week%2 == 0 || ct.Before(t) {
				ct = ct.AddDate(0, 0, 1)
				_, week = ct.ISOWeek()
			}

		case pb.Reminder_MONTHLY:
			ct = ct.AddDate(0, 1, 0)
		case pb.Reminder_YEARLY:
			ct = ct.AddDate(1, 0, 0)
		case pb.Reminder_HALF_YEARLY:
			ct = ct.AddDate(0, 6, 0)
		}

		r.NextRunTime = ct.Unix()
	}
}

func (s *Server) runOnce() {
	ctx, cancel := utils.ManualContext("reminder-loop", "reminder-loop", time.Minute, true)
	defer cancel()

	config, err := s.loadReminders(ctx)
	if err != nil {
		s.Log(fmt.Sprintf("Unable to load reminders: %v", err))
		return
	}

	sort.SliceStable(config.List.Reminders, func(i, j int) bool {
		return config.List.Reminders[i].GetNextRunTime() < config.List.Reminders[j].GetNextRunTime()
	})

	if len(config.List.Reminders) > 0 && time.Now().After(time.Unix(config.List.Reminders[0].GetNextRunTime(), 0)) {
		s.Log(fmt.Sprintf("Adding reminder: %v", config.List.Reminders[0]))
		//s.ghbridge.addIssue(ctx, config.List.Reminders[0])
		//s.adjustRunTime(config.List.Reminders[0])
		//s.save(ctx, config)
	}
}
