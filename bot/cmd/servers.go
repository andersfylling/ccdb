package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/s1kx/unison"
	"strconv"
)

var ServersCommand = &unison.Command{
	Name:        "servers",
	Description: "Get the number of servers using this bot",
	Action:      serversCommandAction,
	Deactivated: false,
	Permission:  unison.NewCommandPermission(),
}

func serversCommandAction(ctx *unison.Context, m *discordgo.Message, content string) error {
	_, err := ctx.Bot.Discord.ChannelMessageSend(m.ChannelID, strconv.FormatInt(int64(len(ctx.Bot.Discord.State.Guilds)), 10))

	return err
}
