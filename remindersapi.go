package main

import (
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/reminders/proto"
)

//Server is the main server
type Server struct {
	*goserver.GoServer
	data     *pb.ReminderConfig
	ghbridge githubBridge
}

type githubBridge interface {
	isComplete(t *pb.Reminder) bool
	addIssue(t *pb.Reminder) string
}

//AddReminder adds a reminder into the system
func (s *Server) AddReminder(ctx context.Context, in *pb.Reminder) (*pb.Empty, error) {
	t := time.Now()
	s.data.List.Reminders = append(s.data.List.Reminders, in)
	s.saveReminders()
	s.LogFunction("AddReminder", t)
	return &pb.Empty{}, nil
}

//ListReminders lists all the available reminders
func (s *Server) ListReminders(ctx context.Context, in *pb.Empty) (*pb.ReminderList, error) {
	return s.data.List, nil
}

//AddTaskList adds a task list into the system
func (s *Server) AddTaskList(ctx context.Context, in *pb.TaskList) (*pb.Empty, error) {
	t := time.Now()

	//Ensure all tasks in the list are unassigned
	for _, task := range in.GetTasks().GetReminders() {
		task.CurrentState = pb.Reminder_UNASSIGNED
	}

	s.data.Tasks = append(s.data.Tasks, in)
	s.processTaskList(in)

	s.saveReminders()
	s.LogFunction("AddTaskList", t)
	return &pb.Empty{}, nil
}
