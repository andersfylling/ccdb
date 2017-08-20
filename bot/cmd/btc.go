package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/s1kx/unison"
)

var BtcCommand = &unison.Command{
	Name:        "btc",
	Description: "Get the latest BTC:USD rate",
	Action:      btcCommandAction,
	Deactivated: false,
	Permission:  unison.NewCommandPermission(),
}

func btcCommandAction(ctx *unison.Context, m *discordgo.Message, content string) error {
	c := ctx.Bot.GetServiceData("btc Bitfinex watcher", "btc_bitfinex_usd")
	_, err := ctx.Bot.Discord.ChannelMessageSend(m.ChannelID, c)

	return err
}
