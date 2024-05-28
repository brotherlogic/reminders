package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/net/context"

	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/reminders/proto"
)

func (s *Server) adjustRunTime(ctx context.Context, r *pb.Reminder) {
	s.CtxLog(ctx, fmt.Sprintf("Adjusting for %v", r.Text))
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
		blah, err := s.addIssue(ctx, reminder)
		s.CtxLog(ctx, fmt.Sprintf("Added reminder %v -> %v", blah, err))
	} else {
		err := s.pingServer(ctx, reminder.GetServer())
		s.CtxLog(ctx, fmt.Sprintf("Pinged %v -> %v", reminder.GetServer(), err))
	}
	return err
}

func (s *Server) runOnce() {
	ctx, cancel := utils.ManualContext("reminder-loop", time.Minute*30)
	defer cancel()

	key, err := s.RunLockingElection(ctx, "reminder-loop", "locking for reminders")
	if err != nil {
		s.CtxLog(ctx, fmt.Sprintf("Unable to lect: %v", err))
		return
	}

	defer s.ReleaseLockingElection(ctx, "reminder-loop", key)

	config, err := s.loadReminders(ctx)
	if err != nil {
		s.CtxLog(ctx, fmt.Sprintf("Unable to load reminders: %v", err))
		return
	}

	sort.SliceStable(config.List.Reminders, func(i, j int) bool {
		return config.List.Reminders[i].GetNextRunTime() < config.List.Reminders[j].GetNextRunTime()
	})

	if len(config.List.Reminders) > 0 && time.Now().After(time.Unix(config.List.Reminders[0].GetNextRunTime(), 0)) {
		err := s.writeReminder(ctx, config.List.Reminders[0])
		if err == nil {
			s.adjustRunTime(ctx, config.List.Reminders[0])
			s.save(ctx, config)
		}
	}
}
