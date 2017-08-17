package main

import (
	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/reminders/proto"
)

//Server is the main server
type Server struct {
	*goserver.GoServer
	reminders []*pb.Reminder
}

//AddReminder adds a reminder into the system
func (s *Server) AddReminder(ctx context.Context, in *pb.Reminder) (*pb.Empty, error) {

	s.reminders = append(s.reminders, in)

	return &pb.Empty{}, nil
}

//ListReminders lists all the available reminders
func (s *Server) ListReminders(ctx context.Context, in *pb.Empty) (*pb.ReminderList, error) {
	return &pb.ReminderList{Reminders: s.reminders}, nil
}
