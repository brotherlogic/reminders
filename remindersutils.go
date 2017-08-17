package main

import (
	"log"
	"time"

	pb "github.com/brotherlogic/reminders/proto"
)

func adjustRunTime(r *pb.Reminder) {
	t := time.Now()
	ct := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	for ct.Weekday().String() != r.DayOfWeek || ct.Before(t) {
		ct = ct.AddDate(0, 0, 1)

	}

	log.Printf("Adjusted to: %v", ct)
	r.NextRunTime = ct.Unix()
}

func (s *Server) getReminders(t time.Time) []*pb.Reminder {
	reminders := make([]*pb.Reminder, 0)

	for _, r := range s.reminders {
		if r.NextRunTime < t.Unix() {
			adjustRunTime(r)
			reminders = append(reminders, r)
		}
	}

	return reminders
}
