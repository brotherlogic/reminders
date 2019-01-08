package main

import (
	"testing"
	"time"

	pb "github.com/brotherlogic/reminders/proto"
	"golang.org/x/net/context"
)

func TestAddOpenEndedReminder(t *testing.T) {
	s := InitTestServer(".testopenended")

	tim := time.Now()
	s.AddReminder(context.Background(), &pb.Reminder{Text: "OpenEnded", NextRunTime: tim.Unix() - 1, RepeatPeriodInSeconds: 100})
	s.processLoop(context.Background())

	reminders, err := s.ListReminders(context.Background(), &pb.Empty{})

	if err != nil {
		t.Fatalf("Error listing reminders: %v", err)
	}

	if len(reminders.List.Reminders) != 1 {
		t.Fatalf("Wrong number of reminders: %v", len(reminders.List.Reminders))
	}

	if time.Unix(reminders.List.Reminders[0].NextRunTime, 0).Sub(tim).Truncate(time.Second) == time.Second*100 {
		t.Errorf("Wrong next run time: %v", time.Unix(reminders.List.Reminders[0].NextRunTime-1, 0).Sub(tim).Truncate(time.Second))
	}
}
