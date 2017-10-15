package bot

import (
	//"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/andersfylling/ccdb/bot/cmd"
	"github.com/andersfylling/ccdb/bot/service"
	"github.com/andersfylling/ccdb/config"
	"github.com/s1kx/unison"
)

// RunBot setup bot interface
func RunBot(conf *config.Config) {
	// Create bot structure
	settings := &unison.BotSettings{
		Token: conf.Discord.Token,

		Commands: []*unison.Command{
			cmd.BtcCommand,
			cmd.ServersCommand,
		},
		EventHooks: []*unison.EventHook{},
		Services: []*unison.Service{
			service.BTCBitfinexService,
		},

		CommandPrefix: conf.Bot.CommandPrefix,
	}

	// Start the bot
	err := unison.RunBot(settings)
	if err != nil {
		logrus.Error(err)
	}
}
