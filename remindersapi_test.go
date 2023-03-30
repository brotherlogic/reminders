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
	server := &Server{data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}}, ghbridge: testGHBridge{completes: make(map[string]bool), issues: make(map[string]string)}}
	server.GoServer = &goserver.GoServer{}
	server.SkipLog = true
	server.SkipIssue = true
	server.SkipElect = true
	server.Register = server
	server.GoServer.KSclient = *keystoreclient.GetTestClient(foldername)
	server.GoServer.KSclient.Save(context.Background(), KEY, &pb.ReminderConfig{})
	server.silence = &testSilence{}
	server.test = true
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
