package main

import (
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	pbd "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

//Server is the main server
type Server struct {
	*goserver.GoServer
	data         *pb.ReminderConfig
	ghbridge     githubBridge
	last         *pbd.Issue
	lastBasicRun time.Time
}

type githubBridge interface {
	isComplete(t *pb.Reminder) bool
	addIssue(t *pb.Reminder) (string, error)
}

//AddReminder adds a reminder into the system
func (s *Server) AddReminder(ctx context.Context, in *pb.Reminder) (*pb.Empty, error) {
	t := time.Now()
	in.Uid = time.Now().UnixNano()
	if s.data.List == nil {
		s.data.List = &pb.ReminderList{}
	}
	if s.data.List.Reminders == nil {
		s.data.List.Reminders = []*pb.Reminder{}
	}
	s.data.List.Reminders = append(s.data.List.Reminders, in)
	s.save()
	s.LogFunction("AddReminder", t)
	return &pb.Empty{}, nil
}

//ListReminders lists all the available reminders
func (s *Server) ListReminders(ctx context.Context, in *pb.Empty) (*pb.ReminderConfig, error) {
	return s.data, nil
}

//AddTaskList adds a task list into the system
func (s *Server) AddTaskList(ctx context.Context, in *pb.TaskList) (*pb.Empty, error) {
	t := time.Now()

	//Ensure all tasks in the list are unassigned
	for _, task := range in.GetTasks().GetReminders() {
		task.CurrentState = pb.Reminder_UNASSIGNED
		task.Uid = time.Now().UnixNano()
	}

	s.data.Tasks = append(s.data.Tasks, in)
	s.save()
	go s.processTaskList(in)

	s.LogFunction("AddTaskList", t)
	return &pb.Empty{}, nil
}

//DeleteTask deletes a task
func (s *Server) DeleteTask(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	t := time.Now()

	for i, reminder := range s.data.GetList().GetReminders() {
		if reminder.GetUid() == in.GetUid() {
			s.data.GetList().Reminders = append(s.data.GetList().Reminders[:i], s.data.GetList().Reminders[i+1:]...)
		}
	}

	for _, tasklist := range s.data.GetTasks() {
		for i, reminder := range tasklist.GetTasks().GetReminders() {
			if reminder.GetUid() == in.GetUid() {
				tasklist.GetTasks().Reminders = append(tasklist.GetTasks().Reminders[:i], tasklist.GetTasks().Reminders[i+1:]...)
			}
		}
	}

	s.save()
	s.LogFunction("Delete", t)
	return &pb.DeleteResponse{}, nil
}
