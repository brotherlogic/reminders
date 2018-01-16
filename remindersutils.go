package main

import (
	"time"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

func (s *Server) refresh() {
	for _, tl := range s.data.GetTasks() {
		s.processTaskList(tl)
	}
}

func (s *Server) processTaskList(t *pb.TaskList) {
	for _, task := range t.Tasks.Reminders {

		//Reassign a task with an empty id
		if task.GetGithubId() == "" {
			task.CurrentState = pb.Reminder_UNASSIGNED
		}

		switch task.GetCurrentState() {
		case pb.Reminder_UNASSIGNED:
			task.CurrentState = pb.Reminder_ASSIGNED
			t, err := s.ghbridge.addIssue(task)
			if err == nil {
				task.GithubId = t
				s.last = &pbgh.Issue{Service: task.GithubId}
				s.save()
			}

			return
		case pb.Reminder_ASSIGNED:
			if s.ghbridge.isComplete(task) {
				task.CurrentState = pb.Reminder_COMPLETE
			} else {
				return
			}
		}
	}
}

func (s *Server) getReminders(t time.Time) []*pb.Reminder {
	reminders := make([]*pb.Reminder, 0)

	for _, r := range s.data.List.Reminders {
		if r.NextRunTime < t.Unix() {
			adjustRunTime(r)
			reminders = append(reminders, r)
		}
	}

	return reminders
}

func adjustRunTime(r *pb.Reminder) {
	t := time.Now()
	ct := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	switch r.RepeatPeriod {
	case pb.Reminder_DAILY:
		ct = ct.AddDate(0, 0, 1)
	case pb.Reminder_WEEKLY:
		for (r.DayOfWeek != "" && ct.Weekday().String() != r.DayOfWeek) || ct.Before(t) {
			ct = ct.AddDate(0, 0, 1)
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
