package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/brotherlogic/goserver"
	keystoreclient "github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/reminders/proto"
)

type testSilence struct {
	failAdd bool
	failRem bool
}

func (t *testSilence) addSilence(ctx context.Context, silence, key string) error {
	if t.failAdd {
		return fmt.Errorf("Built to fail")
	}
	return nil
}

func (t *testSilence) removeSilence(ctx context.Context, key string) error {
	if t.failRem {
		return fmt.Errorf("Built to fail")
	}
	return nil
}

type testGHBridge struct {
	completes map[string]bool
	issues    map[string]string
	fail      bool
}

func (githubBridge testGHBridge) isComplete(ctx context.Context, t *pb.Reminder) bool {
	if val, ok := githubBridge.completes[t.GetText()]; ok {
		return val
	}
	return false
}

func (githubBridge testGHBridge) addIssue(ctx context.Context, t *pb.Reminder) (string, error) {
	log.Printf("ADDING: %v", githubBridge.fail)
	if githubBridge.fail {
		return "blah", fmt.Errorf("Failed to add issue")
	}
	if val, ok := githubBridge.issues[t.GetText()]; ok {
		return val, nil
	}

	githubBridge.issues[t.GetText()] = "blah"

	return "added", nil
}

func InitTestServer(foldername string) *Server {
	server := &Server{data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}, ghbridge: testGHBridge{completes: make(map[string]bool), issues: make(map[string]string)}}
	server.GoServer = &goserver.GoServer{}
	server.SkipLog = true
	server.SkipIssue = true
	server.Register = server
	server.GoServer.KSclient = *keystoreclient.GetTestClient(foldername)
	server.GoServer.KSclient.Save(context.Background(), KEY, &pb.ReminderConfig{})
	server.silence = &testSilence{}
	server.test = true
	return server
}

func TestAddTaskListWithFail(t *testing.T) {
	s := InitTestServer(".testaddtasklist")
	s.ghbridge = testGHBridge{fail: true}
	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One"}, &pb.Reminder{Text: "This is Task Two"}}}})
	s.refresh(context.Background())
	if err != nil {
		t.Fatalf("Error adding task list: %v", err)
	}

	if s.pushFail != 1 {
		t.Errorf("Fail did not increment fail count: %v", s.pushFail)
	}
}

func TestAddTaskList(t *testing.T) {
	s := InitTestServer(".testaddtasklist")
	s.AddReminder(context.Background(), &pb.Reminder{Text: "This is a regular reminder", DayOfWeek: "Sunday"})
	s.ghbridge.(testGHBridge).issues["This is Task One"] = "issue1"
	s.ghbridge.(testGHBridge).issues["This is Task Two"] = "issue2"
	log.Printf("BRIDGE: %v", s.ghbridge)
	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One", Silences: []string{"test"}}, &pb.Reminder{Text: "This is Task Two"}}}})
	if err != nil {
		t.Fatalf("Error adding task list: %v", err)
	}

	// Sleep to allow stuff to process
	time.Sleep(time.Second)

	log.Printf("BRIDGE IS %v", s.ghbridge)
	if s.last.Service != "issue1" {
		t.Errorf("Reminders were not created: %v", s.last)
	}

	s.refresh(context.Background())
	s.ghbridge.(testGHBridge).completes["This is Task One"] = true
	s.refresh(context.Background())

	if s.last.Service != "issue2" {
		t.Errorf("Reminders were not refreshed: %v", s.last)
	}
}

func TestAddTaskListWithSilenceRemoveFail(t *testing.T) {
	s := InitTestServer(".testaddtasklist")
	s.AddReminder(context.Background(), &pb.Reminder{Text: "This is a regular reminder", DayOfWeek: "Sunday"})
	s.ghbridge.(testGHBridge).issues["This is Task One"] = "issue1"
	s.ghbridge.(testGHBridge).issues["This is Task Two"] = "issue2"
	s.silence = &testSilence{failRem: true}
	log.Printf("BRIDGE: %v", s.ghbridge)
	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One", Silences: []string{"test"}}, &pb.Reminder{Text: "This is Task Two"}}}})
	if err != nil {
		t.Fatalf("Error adding task list: %v", err)
	}

	// Sleep to allow stuff to process
	time.Sleep(time.Second)

	log.Printf("BRIDGE IS %v", s.ghbridge)
	if s.last.Service != "issue1" {
		t.Errorf("Reminders were not created: %v", s.last)
	}

	s.refresh(context.Background())
	s.ghbridge.(testGHBridge).completes["This is Task One"] = true
	s.refresh(context.Background())

	if s.last.Service == "issue1" {
		t.Errorf("Reminders were not refreshed: %v", s.last)
	}
}

func TestAddTaskListWithSilenceFail(t *testing.T) {
	s := InitTestServer(".testaddtasklist")
	s.AddReminder(context.Background(), &pb.Reminder{Text: "This is a regular reminder", DayOfWeek: "Sunday"})
	s.ghbridge.(testGHBridge).issues["This is Task One"] = "issue1"
	s.ghbridge.(testGHBridge).issues["This is Task Two"] = "issue2"
	s.silence = &testSilence{failAdd: true}
	log.Printf("BRIDGE: %v", s.ghbridge)
	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One", Silences: []string{"test"}}, &pb.Reminder{Text: "This is Task Two"}}}})
	if err != nil {
		t.Fatalf("Error adding task list: %v", err)
	}

	// Sleep to allow stuff to process
	time.Sleep(time.Second)

	log.Printf("BRIDGE IS %v", s.ghbridge)
	if s.last != nil {
		t.Errorf("Reminders were created: %v", s.last)
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

func TestAddReminderWithLoadFail(t *testing.T) {
	s := InitTestServer(".testaddlist")
	s.GoServer.KSclient.Fail = true

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello"})
	if err == nil {
		t.Fatalf("Add reminder did not fail: %v", err)
	}
}

func TestListWithFail(t *testing.T) {
	s := InitTestServer(".testaddlist")
	s.GoServer.KSclient.Fail = true

	rs, err := s.ListReminders(context.Background(), &pb.Empty{})
	if err == nil {
		t.Fatalf("List reminders did not fail: %v", rs)
	}
}

func TestDeleteWithFail(t *testing.T) {
	s := InitTestServer(".testdeletelist")
	s.GoServer.KSclient.Fail = true

	rs, err := s.DeleteTask(context.Background(), &pb.DeleteRequest{})
	if err == nil {
		t.Fatalf("List reminders did not fail: %v", rs)
	}
}

func TestDeleteDaily(t *testing.T) {
	s := InitTestServer(".testmonthlyreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", NextRunTime: time.Now().Unix(), RepeatPeriod: pb.Reminder_DAILY})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	rems, err := s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error listing: %v", err)
	}

	if len(rems.GetList().GetReminders()) != 1 {
		t.Fatalf("Whaaaa: %v", rems)
	}

	s.DeleteTask(context.Background(), &pb.DeleteRequest{Uid: rems.GetList().GetReminders()[0].Uid})

	rems, err = s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error listing: %v", err)
	}

	if len(rems.GetList().GetReminders()) != 0 {
		t.Fatalf("Whaaaa: %v", rems)
	}
}
