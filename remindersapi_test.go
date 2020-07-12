package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
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
	return "added", nil
}

func InitTestServer(foldername string) *Server {
	server := &Server{data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}, ghbridge: testGHBridge{completes: make(map[string]bool), issues: make(map[string]string)}}
	server.GoServer = &goserver.GoServer{}
	server.SkipLog = true
	server.Register = server
	server.GoServer.KSclient = *keystoreclient.GetTestClient(foldername)
	server.silence = &testSilence{}
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
		t.Errorf("Fail did not increment fail count")
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

func TestDailyReminder(t *testing.T) {
	s := InitTestServer(".testmonthlyreminder")

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", NextRunTime: time.Now().Unix(), RepeatPeriod: pb.Reminder_DAILY})
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

	if dt2.Day() != dt1.Day()+1 {
		t.Errorf("Run time %v should be the day after %v", dt2, dt1)
	}
}

func TestBiweeklyReminder(t *testing.T) {
	s := InitTestServer(".testbiweklyreminder")

	ti := time.Now()
	_, w := ti.ISOWeek()
	if w%2 == 0 {
		ti = ti.AddDate(0, 0, -7)
	}

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", NextRunTime: ti.Unix(), RepeatPeriod: pb.Reminder_BIWEEKLY, DayOfWeek: "Thursday"})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	t1 := time.Now().Add(time.Second)
	rs := s.getReminders(t1)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders")
	}

	t2 := rs[0].NextRunTime
	dt1 := time.Unix(ti.Unix(), 0)
	dt2 := time.Unix(t2, 0)

	if dt2.Weekday() != time.Thursday {
		t.Errorf("Run time %v should be the second week after %v -> %v", dt2.Weekday(), time.Hour*24*7, dt2.Sub(dt1))
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

	_, err := s.AddReminder(context.Background(), &pb.Reminder{Text: "Hello", DayOfWeek: "Monday"})
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
	if len(rs) == 0 {
		t.Fatalf("Wrong number of reminders on second call: %v with %v", rs, t2)
	}

	t3 := t2.Add(time.Hour * 24 * 7)
	rs = s.getReminders(t3)
	if len(rs) != 1 {
		t.Fatalf("Wrong number of reminders on third call: %v", rs)
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

func TestDeleteTask(t *testing.T) {
	s := InitTestServer(".testmonthlyreminder")

	_, err := s.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is Task One"}}}})
	if err != nil {
		t.Fatalf("Error adding reminder: %v", err)
	}

	rems, err := s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error listing: %v", err)
	}

	if len(rems.GetTasks()[0].GetTasks().GetReminders()) != 1 {
		t.Fatalf("Whaaaa: %v", rems)
	}

	s.DeleteTask(context.Background(), &pb.DeleteRequest{Uid: rems.GetTasks()[0].GetTasks().GetReminders()[0].Uid})

	rems, err = s.ListReminders(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Error listing: %v", err)
	}

	if len(rems.GetTasks()[0].GetTasks().GetReminders()) != 0 {
		t.Fatalf("Whaaaa: %v", rems)
	}
}
