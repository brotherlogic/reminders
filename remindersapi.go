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
	config, err := s.loadReminders(ctx)
	if err != nil {
		return nil, err
	}
	in.Uid = time.Now().UnixNano()
	if config.List == nil {
		config.List = &pb.ReminderList{}
	}
	config.GetList().Reminders = append(config.GetList().GetReminders(), in)

	return &pb.Empty{}, s.save(ctx, config)
}

//ListReminders lists all the available reminders
func (s *Server) ListReminders(ctx context.Context, in *pb.Empty) (*pb.ReminderConfig, error) {
	config, err := s.loadReminders(ctx)
	if err != nil {
		return nil, err
	}

	return config, nil
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
	config, err := s.loadReminders(ctx)
	if err != nil {
		return nil, err
	}
	for i, reminder := range config.GetList().GetReminders() {
		if reminder.GetUid() == in.GetUid() {
			config.GetList().Reminders = append(config.GetList().Reminders[:i], config.GetList().Reminders[i+1:]...)
		}
	}

	for _, tasklist := range config.GetTasks() {
		for i, reminder := range tasklist.GetTasks().GetReminders() {
			if reminder.GetUid() == in.GetUid() {
				tasklist.GetTasks().Reminders = append(tasklist.GetTasks().Reminders[:i], tasklist.GetTasks().Reminders[i+1:]...)
			}
		}
	}

	s.save(ctx, config)
	return &pb.DeleteResponse{}, nil
}
