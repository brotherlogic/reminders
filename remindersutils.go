package main

import (
	"fmt"
	"log"
	"time"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

func (s *Server) refresh() {
	log.Printf("Refreshing")
	for _, tl := range s.data.GetTasks() {
		s.processTaskList(tl)
	}
}

func (s *Server) processTaskList(t *pb.TaskList) {
	for _, task := range t.Tasks.Reminders {
		log.Printf("Task = %v (%v)", task, task.GetCurrentState())
		switch task.GetCurrentState() {
		case pb.Reminder_UNASSIGNED:
			task.CurrentState = pb.Reminder_ASSIGNED
			task.GithubId = s.ghbridge.addIssue(task)
			s.last = &pbgh.Issue{Service: task.GithubId}
			return
		case pb.Reminder_ASSIGNED:
			log.Printf("COMPLETE? %v", s.ghbridge.isComplete(task))
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

	log.Printf("HERE")
	for _, r := range s.data.List.Reminders {
		s.Log(fmt.Sprintf("TESTING %v and %v", r, r.NextRunTime-t.Unix()))
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
	case pb.Reminder_WEEKLY:
		for ct.Weekday().String() != r.DayOfWeek || ct.Before(t) {
			ct = ct.AddDate(0, 0, 1)
			log.Printf("%v -> %v", ct, t)
			log.Printf("%v vs %v", ct.Weekday().String(), r.DayOfWeek)
		}
	case pb.Reminder_MONTHLY:
		ct = ct.AddDate(0, 1, 0)
	case pb.Reminder_YEARLY:
		ct = ct.AddDate(1, 0, 0)
	case pb.Reminder_HALF_YEARLY:
		ct = ct.AddDate(0, 6, 0)
	}

	log.Printf("Adjusted to: %v", ct)
	r.NextRunTime = ct.Unix()
}
