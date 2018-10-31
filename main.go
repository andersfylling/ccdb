package main

import (
	//"github.com/sirupsen/logrus"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/event"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	// EnvVarPrefix is the prefix for environment variables
	EnvVarPrefix = "CCDB_"
	BotTokenKey = EnvVarPrefix + "TOKEN"

	CommandPrefix = "ccdb!"
)

func equalCommand(input, command string) bool {
	return strings.HasPrefix(input, CommandPrefix + command)
}

func main() {
	// create a Disgord session
	session, err := disgord.NewSession(&disgord.Config{
		Token: os.Getenv(BotTokenKey),
	})
	if err != nil {
		panic(err)
	}

	// register commands
	session.On(event.MessageCreate, about)
	session.On(event.MessageCreate, servers)
	//session.On(event.MessageCreate, bitfinex)

	// connect to the discord gateway to receive events
	err = session.Connect()
	if err != nil {
		panic(err)
	}

	stop := make(chan interface{})
	go statusUpdateScheduler(session, getBitfinexRate, stop)

	// Keep the socket connection alive, until you terminate the application
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-termSignal
	close(stop)
	session.Disconnect()
}