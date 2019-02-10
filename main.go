package main

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/event"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const (
	// EnvVarPrefix is the prefix for environment variables
	EnvVarPrefix = "CCDB_"
	BotTokenKey  = EnvVarPrefix + "TOKEN"

	CommandPrefix = "ccdb!"
)

func equalCommand(input, command string) bool {
	return strings.HasPrefix(input, CommandPrefix+command)
}

func main() {
	// create a Disgord client
	client := disgord.New(&disgord.Config{
		BotToken:     os.Getenv(BotTokenKey),
		Logger:       disgord.DefaultLogger(true),
		DisableCache: true, // don't need it
	})

	// register commands
	client.On(event.MessageCreate, about)
	client.On(event.MessageCreate, servers)
	client.On(event.GuildCreate, func(session disgord.Session, evt *disgord.GuildCreate) {
		fmt.Println("joined guild", evt.Guild.Name)
	})

	stop := make(chan interface{})
	defer close(stop)
	go statusUpdateScheduler(client, getBitfinexRate, stop)

	// connect to the discord gateway to receive events
	if err := client.StayConnectedUntilInterrupted(); err != nil {
		logrus.Error(err)
	}
}
