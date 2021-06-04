package main

import (
	"testing"
	"time"

	pb "github.com/brotherlogic/reminders/proto"
	"golang.org/x/net/context"
)

func TestBasicReminder(t *testing.T) {
	s := InitTestServer(".testbasicreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{NextRunTime: time.Now().Unix(), RepeatPeriodInSeconds: 10})

	if err != nil {
		t.Fatalf("Unable to add reminder: %v", err)
	}

	s.runOnce()

	if len(s.ghbridge.(testGHBridge).issues) != 1 {
		t.Errorf("Issue not added")
	}
}

func TestBasicReminderDaily(t *testing.T) {
	s := InitTestServer(".testbasicreminder")

	t1 := time.Now()
	_, err := s.AddReminder(context.Background(), &pb.Reminder{NextRunTime: t1.Unix(), RepeatPeriod: pb.Reminder_DAILY})

	if err != nil {
		t.Fatalf("Unable to add reminder: %v", err)
	}

	s.runOnce()

	if len(s.ghbridge.(testGHBridge).issues) != 1 {
		t.Errorf("Issue not added")
	}

	next, err := s.ListReminders(context.Background(), &pb.Empty{})
	if len(next.GetList().GetReminders()) == 0 {
		t.Errorf("No tasks listed")
	}

	if time.Unix(next.GetList().GetReminders()[0].GetNextRunTime(), 0).YearDay() != t1.YearDay()+1 {
		t.Errorf("Time gap is too small: %v", time.Unix(next.GetList().GetReminders()[0].GetNextRunTime(), 0).Sub(t1))
	}
}

func TestBasicReminderBiWeekly(t *testing.T) {
	s := InitTestServer(".testbasicreminder")

	t1 := time.Now()
	_, err := s.AddReminder(context.Background(), &pb.Reminder{NextRunTime: t1.Unix(), RepeatPeriod: pb.Reminder_BIWEEKLY, DayOfWeek: "Tues"})

	if err != nil {
		t.Fatalf("Unable to add reminder: %v", err)
	}

	s.runOnce()

	if len(s.ghbridge.(testGHBridge).issues) != 1 {
		t.Errorf("Issue not added")
	}

	next, err := s.ListReminders(context.Background(), &pb.Empty{})
	if len(next.GetList().GetReminders()) == 0 {
		t.Errorf("No tasks listed")
	}

	if time.Unix(next.GetList().GetReminders()[0].GetNextRunTime(), 0).Sub(t1) < time.Hour*24*13 {
		t.Errorf("Time gap is too small: %v", time.Unix(next.GetList().GetReminders()[0].GetNextRunTime(), 0).Sub(t1))
	}
}
