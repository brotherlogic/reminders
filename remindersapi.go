package main

import (
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/reminders/proto"
)

type githubBridge interface {
	isComplete(ctx context.Context, t *pb.Reminder) bool
	addIssue(ctx context.Context, t *pb.Reminder) (string, error)
}

//AddReminder adds a reminder into the system
func (s *Server) AddReminder(ctx context.Context, in *pb.Reminder) (*pb.Empty, error) {
	in.Uid = time.Now().UnixNano()
	s.data.List.Reminders = append(s.data.List.Reminders, in)
	//s.save(ctx)
	return &pb.Empty{}, nil
}

//ListReminders lists all the available reminders
func (s *Server) ListReminders(ctx context.Context, in *pb.Empty) (*pb.ReminderConfig, error) {
	return s.data, nil
}

//AddTaskList adds a task list into the system
func (s *Server) AddTaskList(ctx context.Context, in *pb.TaskList) (*pb.Empty, error) {
	//Ensure all tasks in the list are unassigned
	for _, task := range in.GetTasks().GetReminders() {
		task.CurrentState = pb.Reminder_UNASSIGNED
		task.Uid = time.Now().UnixNano()
	}

	s.data.Tasks = append(s.data.Tasks, in)
	//s.save(ctx)
	go s.processTaskList(ctx, in)

	return &pb.Empty{}, nil
}

//DeleteTask deletes a task
func (s *Server) DeleteTask(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
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

	//s.save(ctx)
	return &pb.DeleteResponse{}, nil
}
