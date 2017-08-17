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

// InitServer builds an initial server
func InitServer() Server {
	server := Server{GoServer: &goserver.GoServer{}, reminders: make([]*pb.Reminder, 0)}
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	return server
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

	server.Register = server
	server.PrepServer()
	server.RegisterServer("reminders", false)

	server.Serve()
}
