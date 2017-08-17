package main

import (
	"context"
	"testing"

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
