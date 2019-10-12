package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/reminders/proto"

	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {

	host, port, _ := utils.Resolve("reminders", "reminders-cli")
	conn, err := grpc.Dial(host+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	ctx, cancel := utils.BuildContext("reminders_cli_"+os.Args[1], "reminders")
	defer cancel()

	client := pb.NewRemindersClient(conn)

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: addlist\n")
	} else {
		switch os.Args[1] {
		case "list":
			rs, err := client.ListReminders(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Error adding task list %v", err)
			}
			for i, reminder := range rs.List.Reminders {
				fmt.Printf("%v. %v\n", i, reminder)
			}
			for i, task := range rs.Tasks {
				fmt.Printf("%v. %v\n", i, task.Name)
				for j, item := range task.Tasks.Reminders {
					fmt.Printf("%v.%v. %v\n", i, j, item)
				}
			}
		case "add":
			reminder := os.Args[2]
			//day := os.Args[3]
			_, err = client.AddReminder(ctx, &pb.Reminder{Text: reminder, RepeatPeriod: pb.Reminder_MONTHLY})
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
			scanner.Scan()
			silences := strings.Split(scanner.Text(), ",")
			for scanner.Scan() {
				elems := strings.Split(scanner.Text(), "~")
				list.Reminders = append(list.Reminders, &pb.Reminder{Text: elems[0], GithubComponent: elems[1], Silences: silences})
			}

			_, err = client.AddTaskList(ctx, &pb.TaskList{Name: os.Args[2], Tasks: list})
			if err != nil {
				log.Fatalf("Error adding tasks: %v", err)
			}
		case "delete":
			uid, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatalf("Unable to convert UID: %v", err)
			}
			_, err = client.DeleteTask(ctx, &pb.DeleteRequest{Uid: int64(uid)})
			if err != nil {
				log.Fatalf("Delete failed: %v", err)
			}
		}
	}
}
