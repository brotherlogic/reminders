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

func (s *Server) adjustRunTime(r *pb.Reminder) {
	s.Log(fmt.Sprintf("Adjusting for %v", r.Text))
	t := time.Now()
	ct := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	if r.RepeatPeriodInSeconds > 0 {
		r.NextRunTime = t.Unix() + r.RepeatPeriodInSeconds
	} else {
		switch r.RepeatPeriod {
		case pb.Reminder_DAILY:
			ct = t.Add(time.Hour * 24)
		case pb.Reminder_WEEKLY:
			ct = t.Add(time.Hour * 24 * 7)
		case pb.Reminder_BIWEEKLY:
			ct = t.Add(time.Hour * 24 * 14)
		case pb.Reminder_MONTHLY:
			ct = t.AddDate(0, 1, 0)
		case pb.Reminder_YEARLY:
			ct = t.AddDate(1, 0, 0)
		case pb.Reminder_HALF_YEARLY:
			ct = t.AddDate(0, 6, 0)
		}

		r.NextRunTime = ct.Unix()
	}
}

func (s *Server) writeReminder(ctx context.Context, reminder *pb.Reminder) (err error) {
	if len(reminder.GetServer()) == 0 {
		blah, err := s.ghbridge.addIssue(ctx, reminder)
		s.Log(fmt.Sprintf("Added reminder %v -> %v", blah, err))
	} else {
		err := s.pingServer(ctx, reminder.GetServer())
		s.Log(fmt.Sprintf("Pinged %v -> %v", reminder.GetServer(), err))
	}
	return err
}

func (s *Server) runOnce() {
	ctx, cancel := utils.ManualContext("reminder-loop", "reminder-loop", time.Minute*30, true)
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
		err := s.writeReminder(ctx, config.List.Reminders[0])
		if err == nil {
			s.adjustRunTime(config.List.Reminders[0])
			err := s.save(ctx, config)
			if err != nil {
				time.Sleep(time.Second * 2)
				s.Log(fmt.Sprintf("Unable to save reminders: %v", err))
			}
		}
	}
}
