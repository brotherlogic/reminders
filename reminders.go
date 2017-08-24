package main

import (
	"flag"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/reminders/proto"
)

const (
	//KEY under which we store the config data
	KEY = "/github.com/brotherlogic/reminders/config"
)

// InitServer builds an initial server
func InitServer() Server {
	server := Server{GoServer: &goserver.GoServer{}, data: &pb.ReminderConfig{}}
	server.PrepServer()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	return server
}

func (s *Server) loadReminders() error {
	config := &pb.ReminderConfig{}
	log.Printf("%v and %v", KEY, config)
	data, err := s.KSclient.Read(KEY, config)

	if err != nil {
		log.Printf("Unable to read collection: %v", err)
		return err
	}

	s.data = data.(*pb.ReminderConfig)
	return nil
}

// DoRegister does RPC registration
func (s Server) DoRegister(server *grpc.Server) {
	pb.RegisterRemindersServer(server, &s)
}

// Mote promotes/demotes this server
func (s Server) Mote(master bool) error {
	return nil
}

// ReportHealth alerts if we're not healthy
func (s Server) ReportHealth() bool {
	log.Printf("REPORTING HEALTH")
	return true
}

func main() {
	var quiet = flag.Bool("quiet", true, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	server := InitServer()
	err := server.loadReminders()
	if err != nil {
		log.Fatalf("Failed to load reminders: %v", err)
	}
	server.Register = server
	server.RegisterServer("reminders", false)

	server.Serve()
}
