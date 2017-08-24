package main

import (
	"context"
	"testing"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/reminders/proto"
)

type testGHBridge struct {
	completes map[string]bool
	issues    map[string]string
}

func (githubBridge testGHBridge) isComplete(t *pb.Reminder) bool {
	if val, ok := githubBridge.completes[t.GetText()]; ok {
		return val
	}
	return false
}

func (githubBridge testGHBridge) addIssue(t *pb.Reminder) string {
	if val, ok := githubBridge.issues[t.GetText()]; ok {
		return val
	}
	return "added"
}

func InitTestServer(foldername string) Server {
	server := Server{data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}, ghbridge: testGHBridge{completes: make(map[string]bool), issues: make(map[string]string)}}
	server.GoServer = &goserver.GoServer{}
	server.SkipLog = true
	server.Register = server
	server.GoServer.KSclient = *keystoreclient.GetTestClient(foldername)

	return server
}

func TestAddTaskList(t *testing.T) {
	s := InitTestServer(".testaddtasklist")
	s.AddReminder(context.Background(), &pb.Reminder{Text: "This is a regular reminder", DayOfWeek: "Sunday"})
	s.ghbridge.(testGHBridge).issues["This is Task One"] = "issue1"
	s.ghbridge.(testGHBridge).issues["This is Task Two"] = "issue2"
	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One"}, &pb.Reminder{Text: "This is Task Two"}}}})
	if err != nil {
		t.Fatalf("Error adding task list: %v", err)
	}

	r, err := s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error getting reminders: %v", err)
	}

	if len(r.Reminders) != 2 || r.Reminders[1].GithubId != "issue1" {
		t.Errorf("Reminders were not created: %v", r)
	}

	s.refresh()
	s.ghbridge.(testGHBridge).completes["This is Task One"] = true
	s.refresh()

	r, err = s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error getting reminders: %v", err)
	}

	if len(r.Reminders) != 2 || r.Reminders[1].GithubId != "issue2" {
		t.Errorf("Reminders were not refreshed: %v", r)
	}
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
