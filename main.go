package main

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/event"
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
	client, err := disgord.NewClient(&disgord.Config{
		BotToken:     os.Getenv(BotTokenKey),
		Logger:       disgord.DefaultLogger(true),
		DisableCache: true, // don't need it
	})
	if err != nil {
		panic(err)
	}

	// register commands
	client.On(event.MessageCreate, about)
	client.On(event.MessageCreate, servers)
	client.On(event.GuildCreate, func(session disgord.Session, evt *disgord.GuildCreate) {
		fmt.Println("joined guild", evt.Guild.Name)
	})

	// connect to the discord gateway to receive events
	if err = client.Connect(); err != nil {
		panic(err)
	}

	stop := make(chan interface{})
	go statusUpdateScheduler(client, getBitfinexRate, stop)

	client.DisconnectOnInterrupt()
	close(stop)
}
