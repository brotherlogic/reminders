package main

import (
	"context"
	"testing"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/reminders/proto"
)

func InitTestServer(foldername string) Server {
	server := Server{reminders: make([]*pb.Reminder, 0)}
	server.GoServer = &goserver.GoServer{}
	server.SkipLog = true
	server.Register = server
	server.GoServer.KSclient = *keystoreclient.GetTestClient(foldername)

	return server
}

func TestAddList(t *testing.T) {
	s := InitTestServer(".testaddlist")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello"})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	rs, err := s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error listing reminders: %v", err)
	}

	if len(rs.Reminders) != 1 {
		t.Errorf("Error getting reminders: %v", rs)
	}
}

func TestBuildReminders(t *testing.T) {
	s := InitTestServer(".testaddlist")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", DayOfWeek: "Sunday"})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	t1 := time.Now()
	rs := s.getReminders(t1)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders")
	}

	t2 := t1.Add(time.Hour * 24)
	rs = s.getReminders(t2)
	if len(rs) != 0 {
		t.Fatalf("Wrong number of reminders on second call: %v", rs)
	}

	t3 := t2.Add(time.Hour * 24 * 7)
	rs = s.getReminders(t3)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders on third call: %v", rs)
	}

}
