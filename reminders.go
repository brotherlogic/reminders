package main

import (
	"flag"
	"fmt"
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

//Server is the main server
type Server struct {
	*goserver.GoServer
	data         *pb.ReminderConfig
	ghbridge     githubBridge
	last         *pbgh.Issue
	lastBasicRun time.Time
	pushCount    int64
	pushFail     int64
	pushFailure  string
	silence      silence
}

const (
	//KEY under which we store the config data
	KEY = "/github.com/brotherlogic/reminders/config"
)

type silence interface {
	addSilence(ctx context.Context, silence, key string) error
	removeSilence(ctx context.Context, key string) error
}

type prodSilence struct {
	dial func(server string) (*grpc.ClientConn, error)
}

func (ps *prodSilence) addSilence(ctx context.Context, silence, key string) error {
	conn, err := ps.dial("githubcard")
	if err != nil {
		return err
	}

	client := pbgh.NewGithubClient(conn)
	_, err = client.Silence(ctx, &pbgh.SilenceRequest{Silence: silence, Origin: key, State: pbgh.SilenceRequest_SILENCE})
	return err
}

func (ps *prodSilence) removeSilence(ctx context.Context, key string) error {
	conn, err := ps.dial("githubcard")
	if err != nil {
		return err
	}

	client := pbgh.NewGithubClient(conn)
	_, err = client.Silence(ctx, &pbgh.SilenceRequest{Origin: key, State: pbgh.SilenceRequest_UNSILENCE})
	return err
}

type gsGHBridge struct {
	getter func(servername string) (string, int)
}

func (s *Server) processLoop(ctx context.Context) error {
	s.lastBasicRun = time.Now()
	s.refresh(ctx)
	rs := s.getReminders(time.Now())
	s.Log("Got reminders (" + strconv.Itoa(len(rs)) + ")")
	for _, r := range rs {
		s.ghbridge.addIssue(ctx, r)
	}
	s.save(ctx)
	return nil
}

func (g gsGHBridge) addIssue(ctx context.Context, r *pb.Reminder) (string, error) {
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
	resp, err := client.AddIssue(ctx, &pbgh.Issue{Service: r.GetGithubComponent(), Title: r.GetText(), Body: "From your reminders"})
	if err != nil {
		return "", err
	}

	return resp.GetService() + "/" + strconv.Itoa(int(resp.GetNumber())), nil
}

func (g gsGHBridge) isComplete(ctx context.Context, r *pb.Reminder) bool {
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
	resp, err := client.Get(ctx, &pbgh.Issue{Number: int32(num), Service: elems[0]})
	if err != nil {
		return false
	}

	return resp.GetState() == pbgh.Issue_CLOSED
}

func (s *Server) save(ctx context.Context) {
	s.KSclient.Save(ctx, KEY, s.data)
}

// InitServer builds an initial server
func InitServer() *Server {
	server := &Server{GoServer: &goserver.GoServer{}, data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}, Tasks: make([]*pb.TaskList, 0)}}
	server.ghbridge = gsGHBridge{getter: server.GetIP}
	server.PrepServer()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	server.silence = &prodSilence{dial: server.DialMaster}
	return server
}

func (s *Server) loadReminders(ctx context.Context) error {
	config := &pb.ReminderConfig{}
	data, _, err := s.KSclient.Read(ctx, KEY, config)

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
		s.save(ctx)
	}

	return nil
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterRemindersServer(server, s)
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.save(ctx)
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	if master {
		err := s.loadReminders(ctx)
		if err != nil {
			return err
		}
		if s.data.List != nil && s.data.List.Reminders != nil && len(s.data.List.Reminders) == 0 {
			s.Log(fmt.Sprintf("No reminders loaded: %v", s.data))
			return fmt.Errorf("Unable to load reminders")
		}
	}
	return nil
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{
		&pbg.State{Key: "last_update_time", TimeValue: s.lastBasicRun.Unix()},
		&pbg.State{Key: "push_count", Value: s.pushCount},
		&pbg.State{Key: "push_fail", Value: s.pushFail},
		&pbg.State{Key: "push_fail_reason", Text: s.pushFailure},
	}
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
	server.RegisterRepeatingTask(server.processLoop, "process_loop", time.Hour*6)

	server.Serve()
}
