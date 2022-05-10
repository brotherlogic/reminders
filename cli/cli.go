package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver/utils"
	pb "github.com/brotherlogic/reminders/proto"

	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	ctx, cancel := utils.BuildContext("reminders_cli_"+os.Args[1], "reminders")
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "reminders")
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
			rs, err := client.ListReminders(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Error adding task list %v", err)
			}
			for i, reminder := range rs.List.Reminders {
				fmt.Printf("%v. %v\n", i, reminder)
			}
		case "server":
			reminder := os.Args[2]
			_, err = client.AddReminder(ctx, &pb.Reminder{Server: reminder, RepeatPeriodInSeconds: int64((time.Hour * 24 * 14).Seconds())})
			if err != nil {
				log.Fatalf("Unable to add reminder: %v", err)
			}
		case "add":
			reminder := os.Args[2]
			_, err = client.AddReminder(ctx, &pb.Reminder{Text: reminder, RepeatPeriodInSeconds: int64((time.Hour * 24).Seconds())})
			if err != nil {
				log.Fatalf("Unable to add reminder: %v", err)
			}

		case "delete":
			uid, err := strconv.ParseInt(os.Args[2], 10, 64)
			if err != nil {
				log.Fatalf("Unable to convert UID: %v", err)
			}
			resp, err := client.DeleteTask(ctx, &pb.DeleteRequest{Uid: int64(uid)})
			if err != nil {
				log.Fatalf("Delete failed: %v", err)
			}
			log.Printf("Got here: %v", resp)
		}
	}
}
