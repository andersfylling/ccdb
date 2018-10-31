package main

import (
	"fmt"
	"github.com/andersfylling/disgord"
)

func about(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	if !equalCommand(msg.Content, "about") {
		return
	}

	aboutBot := fmt.Sprintln() + `
Discord bot to show current Bitcoin value based on Bitfinex statistics
This bot is open source and can be found at GitHub:
	https://github.com/andersfylling/ccdb

This bot was built using Disgord: https://github.com/andersfylling/disgord
`

	msg.RespondString(session, aboutBot)
}

func servers(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	if !equalCommand(msg.Content, "servers") {
		return
	}

	// TODO: Disgord must implement a way to get all guilds
}