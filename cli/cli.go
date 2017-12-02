package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/reminders/proto"

	_ "google.golang.org/grpc/encoding/gzip"
)

func findServer(name string) (string, int) {
	conn, err := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot reach discover server: %v (trying to discover %v)", err, name)
	}
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	r, err := registry.Discover(context.Background(), &pbdi.RegistryEntry{Name: name})

	if err != nil {
		log.Fatalf("Failure to list: %v", err)
		return "", -1
	}
	return r.Ip, int(r.Port)
}

func main() {

	host, port := findServer("reminders")
	conn, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRemindersClient(conn)

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: addlist\n")
	} else {
		switch os.Args[1] {
		case "list":
			rs, err := client.ListReminders(context.Background(), &pb.Empty{})
			if err != nil {
				log.Fatalf("Error adding task list %v", err)
			}
			fmt.Printf("%v", rs)
		case "add":
			reminder := os.Args[2]
			day := os.Args[3]
			_, err = client.AddReminder(context.Background(), &pb.Reminder{Text: reminder, DayOfWeek: day})
			if err != nil {
				log.Fatalf("Unable to add reminder: %v", err)
			}
		case "file":
			list := &pb.ReminderList{Reminders: make([]*pb.Reminder, 0)}
			file, err := os.Open(os.Args[2])
			if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				elems := strings.Split(scanner.Text(), "~")
				list.Reminders = append(list.Reminders, &pb.Reminder{Text: elems[0], GithubComponent: elems[1]})
			}

			_, err = client.AddTaskList(context.Background(), &pb.TaskList{Name: os.Args[2], Tasks: list})
			if err != nil {
				log.Fatalf("Error adding tasks: %v", err)
			}
		}
	}
}
