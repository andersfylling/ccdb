package bot

import (
	//"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/s1kx/unison"
	"github.com/sciencefyll/ccdb/bot/cmd"
	"github.com/sciencefyll/ccdb/bot/service"
	"github.com/sciencefyll/ccdb/config"
)

func RunBot(conf *config.Config) {
	// Create bot structure
	settings := &unison.BotSettings{
		Token: conf.Discord.Token,

		Commands: []*unison.Command{
			cmd.BtcCommand,
		},
		EventHooks: []*unison.EventHook{},
		Services: []*unison.Service{
			service.BTC_bitfinexService,
		},

		CommandPrefix: conf.Bot.CommandPrefix,
	}

	// Start the bot
	err := unison.RunBot(settings)
	if err != nil {
		logrus.Error(err)
	}
}
