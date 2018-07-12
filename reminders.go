package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

const (
	//KEY under which we store the config data
	KEY = "/github.com/brotherlogic/reminders/config"
)

type gsGHBridge struct {
	getter func(servername string) (string, int)
}

func (s *Server) processLoop(ctx context.Context) {
	s.lastBasicRun = time.Now()
	s.refresh()
	rs := s.getReminders(time.Now())
	s.Log("Got reminders (" + strconv.Itoa(len(rs)) + ")")
	for _, r := range rs {
		s.ghbridge.addIssue(r)
	}
	s.save()
}

func (g gsGHBridge) addIssue(r *pb.Reminder) (string, error) {
	ip, port := g.getter("githubcard")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	if r.GetGithubComponent() == "" {
		r.GithubComponent = "home"
	}
	resp, err := client.AddIssue(context.Background(), &pbgh.Issue{Service: r.GetGithubComponent(), Title: r.GetText()})
	if err != nil {
		return "", err
	}

	return resp.GetService() + "/" + strconv.Itoa(int(resp.GetNumber())), nil
}

func (g gsGHBridge) isComplete(r *pb.Reminder) bool {
	ip, port := g.getter("githubcard")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		return false
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	elems := strings.Split(r.GetGithubId(), "/")
	num, _ := strconv.Atoi(elems[1])
	if len(elems[0]) == 0 || num == 0 {
		//Can't process this, so just return true
		return true
	}
	resp, err := client.Get(context.Background(), &pbgh.Issue{Number: int32(num), Service: elems[0]})
	if err != nil {
		return false
	}

	return resp.GetState() == pbgh.Issue_CLOSED
}

func (s *Server) save() {
	s.KSclient.Save(KEY, s.data)
}

// InitServer builds an initial server
func InitServer() *Server {
	server := &Server{GoServer: &goserver.GoServer{}, data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}}
	server.ghbridge = gsGHBridge{getter: server.GetIP}
	server.PrepServer()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	return server
}

func (s *Server) loadReminders() error {
	config := &pb.ReminderConfig{}
	data, _, err := s.KSclient.Read(KEY, config)

	if err != nil {
		return err
	}

	s.data = data.(*pb.ReminderConfig)

	found := false
	for _, reminder := range s.data.GetList().GetReminders() {
		if reminder.GetUid() == 0 {
			reminder.Uid = time.Now().UnixNano()
			time.Sleep(time.Millisecond)
			found = true
		}
	}

	for _, tasklist := range s.data.GetTasks() {
		for _, reminder := range tasklist.GetTasks().GetReminders() {
			if reminder.GetUid() == 0 {
				reminder.Uid = time.Now().UnixNano()
				time.Sleep(time.Millisecond)
				found = true
			}
		}
	}

	if found {
		s.save()
	}

	return nil
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterRemindersServer(server, s)
}

// Mote promotes/demotes this server
func (s *Server) Mote(master bool) error {
	if master {
		return s.loadReminders()
	}
	return nil
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{&pbg.State{Key: "last_update_time", TimeValue: s.lastBasicRun.Unix()}}
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
	server.RegisterServer("reminders", false)

	//Update the tasks every 24 hours
	server.RegisterRepeatingTask(server.processLoop, time.Hour*6)

	server.Serve()
}
