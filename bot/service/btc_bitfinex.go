package service

import (
	"strconv"
	"time"
	//"github.com/bwmarrin/discordgo"
	"github.com/Sirupsen/logrus"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/s1kx/unison"
)

// BTCBitfinexService updates the bot status to the latest btc price in usd
var BTCBitfinexService = &unison.Service{
	Name:        "btc Bitfinex watcher",
	Description: "real time monitor of Bitfinex btc:usd rate.",
	Action:      BTCBitfinexAction,
	Deactivated: false,
	Data: map[string]string{
		"btc_bitfinex_usd": "? USD",
	},
}

// BTCBitfinexAction service action
func BTCBitfinexAction(ctx *unison.Context) error {
	//pull data in real time
	c := bitfinex.NewClient()

	// in case your proxy is using a non valid certificate set to TRUE
	c.WebSocketTLSSkipVerify = false

	err := c.WebSocket.Connect()
	if err != nil {
		logrus.Error("Error connecting to web socket : ", err)
	}
	defer c.WebSocket.Close()

	tickerChan := make(chan []float64)

	c.WebSocket.AddSubscribe(bitfinex.ChanTicker, bitfinex.BTCUSD, tickerChan)
	go updateData(tickerChan, ctx)

	go updateBotPresence(ctx)

	err = c.WebSocket.Subscribe()
	if err != nil {
		logrus.Error(err)
	}

	return err
}

// update the bot status every 12s
func updateBotPresence(ctx *unison.Context) {
	for {
		time.Sleep(12 * time.Second)
		ctx.Bot.Discord.UpdateStatus(0, ctx.Bot.GetServiceData("btc Bitfinex watcher", "btc_bitfinex_usd"))
	}
}

func updateData(in chan []float64, ctx *unison.Context) {
	for {
		data := <-in
		status := strconv.FormatFloat(data[0], 'f', 2, 64) + " USD"

		// store to service data
		ctx.Bot.SetServiceData("btc Bitfinex watcher", "btc_bitfinex_usd", status)
	}
}
