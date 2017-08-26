package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/reminders/proto"
)

func findServer(name string) (string, int) {
	conn, err := grpc.Dial("192.168.86.64:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot reach discover server: %v (trying to discover %v)", err, name)
	}
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceClient(conn)
	r, err := registry.Discover(context.Background(), &pbdi.RegistryEntry{Name: name})

	if err != nil {
		log.Printf("Failure to list: %v", err)
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
			_, err = client.AddTaskList(context.Background(), &pb.TaskList{Name: "Testing", Tasks: &pb.ReminderList{Reminders: []*pb.Reminder{&pb.Reminder{Text: "This is task one"}, &pb.Reminder{Text: "This is task two"}}}})
			if err != nil {
				log.Fatalf("Error adding task list %v", err)
			}
		case "add":
			reminder := os.Args[2]
			day := os.Args[3]
			_, err = client.AddReminder(context.Background(), &pb.Reminder{Text: reminder, DayOfWeek: day})
			if err != nil {
				log.Fatalf("Unable to add reminder: %v", err)
			}
		}
	}
}
