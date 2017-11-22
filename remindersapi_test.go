package main

import (
	"context"
	"log"
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

func (githubBridge testGHBridge) addIssue(t *pb.Reminder) (string, error) {
	if val, ok := githubBridge.issues[t.GetText()]; ok {
		return val, nil
	}
	return "added", nil
}

func InitTestServer(foldername string) *Server {
	server := &Server{data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}, ghbridge: testGHBridge{completes: make(map[string]bool), issues: make(map[string]string)}}
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
	log.Printf("BRIDGE: %v", s.ghbridge)
	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One"}, &pb.Reminder{Text: "This is Task Two"}}}})
	if err != nil {
		t.Fatalf("Error adding task list: %v", err)
	}

	log.Printf("BRIDGE IS %v", s.ghbridge)
	if s.last.Service != "issue1" {
		t.Errorf("Reminders were not created: %v", s.last)
	}

	s.refresh()
	s.ghbridge.(testGHBridge).completes["This is Task One"] = true
	s.refresh()

	if s.last.Service != "issue2" {
		t.Errorf("Reminders were not refreshed: %v", s.last)
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

	if len(rs.GetList().Reminders) != 1 {
		t.Errorf("Error getting reminders: %v", rs)
	}
}

func TestMonthlyReminder(t *testing.T) {
	s := InitTestServer(".testmonthlyreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", NextRunTime: time.Now().Unix(), RepeatPeriod: pb.Reminder_MONTHLY})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	t1 := time.Now().Add(time.Second)
	rs := s.getReminders(t1)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders")
	}

	t2 := rs[0].NextRunTime
	dt1 := time.Unix(t1.Unix(), 0)
	dt2 := time.Unix(t2, 0)

	if dt1.Month() == dt2.Month() {
		t.Errorf("Run time %v should be a month after %v", dt2, dt1)
	}
}

func TestSixMonthReminder(t *testing.T) {
	s := InitTestServer(".testmonthlyreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", NextRunTime: time.Now().Unix(), RepeatPeriod: pb.Reminder_HALF_YEARLY})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	t1 := time.Now().Add(time.Second)
	rs := s.getReminders(t1)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders")
	}

	t2 := rs[0].NextRunTime
	dt1 := time.Unix(t1.Unix(), 0)
	dt2 := time.Unix(t2, 0)

	if dt2.Sub(dt1).Hours() < 5*30*24 {
		t.Errorf("Run time %v should be a month after %v", dt2, dt1)
	}
}

func TestYearlyReminder(t *testing.T) {
	s := InitTestServer(".testmonthlyreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", NextRunTime: time.Now().Unix(), RepeatPeriod: pb.Reminder_YEARLY})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	t1 := time.Now().Add(time.Second)
	rs := s.getReminders(t1)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders")
	}

	t2 := rs[0].NextRunTime
	dt1 := time.Unix(t1.Unix(), 0)
	dt2 := time.Unix(t2, 0)

	if dt1.Year() == dt2.Year() {
		t.Errorf("Run time %v should be a year after %v", dt2, dt1)
	}
}

func TestBuildReminders(t *testing.T) {
	s := InitTestServer(".testaddlist")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", DayOfWeek: "Friday"})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	t1 := time.Now()
	rs := s.getReminders(t1)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders")
	}

	log.Printf("Running second pass")

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
