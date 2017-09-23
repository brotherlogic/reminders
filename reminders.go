package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

const (
	//KEY under which we store the config data
	KEY = "/github.com/brotherlogic/reminders/config"

	//How long to wait between running a reminder loop
	waitTime = time.Hour * 2
)

type gsGHBridge struct {
	getter func(servername string) (string, int)
}

func (s *Server) processLoop() {
	for true {
		s.refresh()
		log.Printf("GOT REFRESH")
		rs := s.getReminders(time.Now())
		s.Log("Got reminders (" + strconv.Itoa(len(rs)) + ")")
		for _, r := range rs {
			s.ghbridge.addIssue(r)
		}
		s.save()

		log.Printf("SLEEPING FOR %v", waitTime)
		time.Sleep(waitTime)
		log.Printf("SLEPT")
	}
}

func (g gsGHBridge) addIssue(r *pb.Reminder) string {
	ip, port := g.getter("githubcard")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial ghc: %v", err)
		return ""
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	resp, err := client.AddIssue(context.Background(), &pbgh.Issue{Service: r.GetGithubComponent(), Title: r.GetText()})
	if err != nil {
		log.Printf("Add issue failed: %v", err)
		return ""
	}

	return resp.GetService() + "/" + strconv.Itoa(int(resp.GetNumber()))
}

func (g gsGHBridge) isComplete(r *pb.Reminder) bool {
	ip, port := g.getter("githubcard")
	log.Printf("DIALLING: %v, %v", ip, port)
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial ghc: %v", err)
		return false
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	elems := strings.Split(r.GetGithubId(), "/")
	num, _ := strconv.Atoi(elems[1])
	log.Printf("GETTING NOW %v and %v", num, elems[0])
	if len(elems[0]) == 0 || num == 0 {
		//Can't process this, so just return true
		return true
	}
	resp, err := client.Get(context.Background(), &pbgh.Issue{Number: int32(num), Service: elems[0]})
	if err != nil {
		log.Printf("Failed to get issue: %v", err)
		return false
	}

	log.Printf("COME BACK: %v-> %v", &pbgh.Issue{Number: int32(num), Service: elems[0]}, resp)
	return resp.GetState() == pbgh.Issue_CLOSED
}

func (s *Server) save() {
	t := time.Now()
	s.KSclient.Save(KEY, s.data)
	s.LogFunction("save", t)
}

// InitServer builds an initial server
func InitServer() Server {
	server := Server{GoServer: &goserver.GoServer{}, data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}}
	server.ghbridge = gsGHBridge{getter: server.GetIP}
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
	server.RegisterServingTask(server.processLoop)

	log.Printf("Initial Config: %v", server.data)

	server.Serve()
}
