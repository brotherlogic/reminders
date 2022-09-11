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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
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
	running      bool
	test         bool
}

var (
	tasklistSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "reminders_tasklist",
		Help: "The size of the print queue",
	})
	taskSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "reminders_task",
		Help: "The size of the print queue",
	})
)

const (
	//KEY under which we store the config data
	KEY = "/github.com/brotherlogic/reminders/config"
)

type silence interface {
	addSilence(ctx context.Context, silence, key string) error
	removeSilence(ctx context.Context, key string) error
}

type prodSilence struct {
	dial func(ctx context.Context, server string) (*grpc.ClientConn, error)
}

func (ps *prodSilence) addSilence(ctx context.Context, silence, key string) error {
	conn, err := ps.dial(ctx, "githubcard")
	if err != nil {
		return err
	}

	client := pbgh.NewGithubClient(conn)
	_, err = client.Silence(ctx, &pbgh.SilenceRequest{Silence: silence, Origin: key, State: pbgh.SilenceRequest_SILENCE})
	return err
}

func (ps *prodSilence) removeSilence(ctx context.Context, key string) error {
	conn, err := ps.dial(ctx, "githubcard")
	if err != nil {
		return err
	}

	client := pbgh.NewGithubClient(conn)
	err = nil
	for err == nil {
		_, err = client.Silence(ctx, &pbgh.SilenceRequest{Origin: key, State: pbgh.SilenceRequest_UNSILENCE})
	}
	return nil
}

type gsGHBridge struct {
	dial func(ctx context.Context, server string) (*grpc.ClientConn, error)
	log  func(ctx context.Context, logs string)
}

func (s *Server) pingServer(ctx context.Context, server string) error {
	if s.test {
		return nil
	}
	conn, err := s.FDialServer(ctx, server)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewReminderReceiverClient(conn)
	_, err = client.Receive(ctx, &pb.ReceiveRequest{})
	return err
}

func (s *Server) run() {
	for s.running {
		ctx, cancel := utils.ManualContext("reminderrun", time.Minute)
		nextRunTime := s.runFull()
		sleepTime := nextRunTime.Sub(time.Now())
		s.CtxLog(ctx, fmt.Sprintf("Sleeping for %v", sleepTime))
		cancel()
		time.Sleep(sleepTime)
	}
}

func (s *Server) runFull() time.Time {
	s.runOnce()
	time.Sleep(time.Second * 5)
	return time.Now().Add(time.Minute)
}

func (g gsGHBridge) addIssue(ctx context.Context, r *pb.Reminder) (string, error) {
	conn, err := g.dial(ctx, "githubcard")
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
	conn, err := g.dial(ctx, "githubcard")
	if err != nil {
		g.log(ctx, fmt.Sprintf("DIAL FAIL: %v", err))
		return false
	}
	defer conn.Close()

	client := pbgh.NewGithubClient(conn)
	elems := strings.Split(r.GetGithubId(), "/")
	num, _ := strconv.Atoi(elems[1])
	if len(elems[0]) == 0 || num == 0 {
		//Can't process this, so just return true
		g.log(ctx, fmt.Sprintf("UNPROCESSABLE: %v %v", elems[0], num))
		return true
	}
	resp, err := client.Get(ctx, &pbgh.Issue{Number: int32(num), Service: elems[0]})
	if err != nil {
		g.log(ctx, fmt.Sprintf("ERRORED: %v", err))
		return false
	}

	g.log(ctx, fmt.Sprintf("GOT RESPONSE: %v", resp))
	return resp.GetState() == pbgh.Issue_CLOSED
}

func (s *Server) save(ctx context.Context, config *pb.ReminderConfig) error {
	return s.KSclient.Save(ctx, KEY, config)
}

// InitServer builds an initial server
func InitServer() *Server {
	server := &Server{GoServer: &goserver.GoServer{}, data: &pb.ReminderConfig{List: &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}}}
	server.ghbridge = gsGHBridge{dial: server.FDialServer, log: server.CtxLog}
	server.silence = &prodSilence{dial: server.FDialServer}

	server.PrepServer("reminders")

	return server
}

func (s *Server) loadReminders(ctx context.Context) (*pb.ReminderConfig, error) {
	config := &pb.ReminderConfig{}
	data, _, err := s.KSclient.Read(ctx, KEY, config)

	if err != nil {
		return nil, err
	}

	config = data.(*pb.ReminderConfig)

	taskSize.Set(float64(len(config.GetList().GetReminders())))

	return config, nil
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterRemindersServer(server, s)
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{}
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

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	server.DiskLog = true
	server.running = true
	go server.run()

	server.Serve()
}
