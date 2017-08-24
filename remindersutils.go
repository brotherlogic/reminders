package main

import (
	"log"
	"time"

	pb "github.com/brotherlogic/reminders/proto"
)

func (s *Server) refresh() {
	log.Printf("Refreshing")
	for _, tl := range s.data.GetTasks() {
		s.processTaskList(tl)
	}
}

func (s *Server) removeReminder(r *pb.Reminder) {
	i := -1
	for j, rem := range s.data.List.GetReminders() {
		if rem.GetText() == r.GetText() {
			i = j
		}
	}

	if i >= 0 {
		s.data.List.Reminders = append(s.data.List.Reminders[0:i], s.data.List.Reminders[i+1:len(s.data.List.Reminders)]...)
	}
}

func (s *Server) processTaskList(t *pb.TaskList) {
	for _, task := range t.Tasks.Reminders {
		log.Printf("Task = %v (%v)", task, task.GetCurrentState())
		switch task.GetCurrentState() {
		case pb.Reminder_UNASSIGNED:
			task.CurrentState = pb.Reminder_ASSIGNED
			task.GithubId = s.ghbridge.addIssue(task)
			s.data.List.Reminders = append(s.data.List.Reminders, task)
			return
		case pb.Reminder_ASSIGNED:
			if s.ghbridge.isComplete(task) {
				task.CurrentState = pb.Reminder_COMPLETE
				s.removeReminder(task)
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

	for ct.Weekday().String() != r.DayOfWeek || ct.Before(t) {
		ct = ct.AddDate(0, 0, 1)

	}

	log.Printf("Adjusted to: %v", ct)
	r.NextRunTime = ct.Unix()
}
