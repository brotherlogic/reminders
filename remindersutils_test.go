package main

import (
	"testing"
	"time"

	pb "github.com/brotherlogic/reminders/proto"
	"golang.org/x/net/context"
)

func TestBasicReminder(t *testing.T) {
	s := InitTestServer(".testbasicreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{NextRunTime: time.Now().Unix()})

	if err != nil {
		t.Fatalf("Unable to add reminder: %v", err)
	}

	s.runOnce()

	if len(s.ghbridge.(testGHBridge).issues) != 1 {
		t.Errorf("Issue not added")
	}
}
