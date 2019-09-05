package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

func (s *Server) refresh(ctx context.Context) {
	for _, tl := range s.data.GetTasks() {
		s.Log(fmt.Sprintf("Processing %v", tl.GetName()))
		s.processTaskList(ctx, tl)
	}
}

func (s *Server) processTaskList(ctx context.Context, t *pb.TaskList) {
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
						s.pushFailure = fmt.Sprintf("%v", err)
						return
					}
				}
				task.CurrentState = pb.Reminder_ASSIGNED
				s.pushCount++
				task.GithubId = t
				s.last = &pbgh.Issue{Service: task.GithubId}
				s.save(ctx)
			} else {
				s.pushFail++
				s.pushFailure = fmt.Sprintf("%v", err)
			}

			return
		case pb.Reminder_ASSIGNED:
			if s.ghbridge.isComplete(ctx, task) {
				err := s.silence.removeSilence(ctx, fmt.Sprintf("%v", task.Uid))
				if err != nil {
					s.Log(fmt.Sprintf("Unable to silence"))
				}
				task.CurrentState = pb.Reminder_COMPLETE
			} else {
				return
			}
		}
	}
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
