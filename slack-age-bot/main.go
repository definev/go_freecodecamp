package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvent(commandEvents <-chan *slacker.CommandEvent) {
	for event := range commandEvents {
		fmt.Println("Command event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
	} 
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3340086094245-3328411645719-I9jcLWjEbFdDsS86bS0mhU30")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A039W9ER78E-3328375428327-b7111bc3823207e4ed09611d4f3c3ee7a954685b75d2097494608e873ffcfcbd")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	
	go printCommandEvent(bot.CommandEvents())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bot.Command("My dob is <year>", &slacker.CommandDefinition{
		Description: "Get your dob",
		Example: "My dob is 2003",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Println(err)
				response.Reply("You must type number for dob")
				return
			}

			age := 2022 - yob
			if age < 0 {
				response.Reply("You are not born yet")
				return
			}
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}