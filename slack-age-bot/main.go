package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// os.Setenv("SLACK_BOT_TOKEN", "xoxb-9389310105058-9416755698960-LLSljeEB8CxVwlnb5PwY9NmQ")
	// os.Setenv("SLACK_APP_TOKEN", "xapp-1-A09BUMR0N8Z-9388035158949-d78f6f9b67afe7617ca09b49f4d37ac5c8311bf9e675115d9e1872dc06e7f5c1")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command(" my yob is <year>", &slacker.CommandDefinition{
		Description: "age calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2024 - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
			fmt.Println(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
