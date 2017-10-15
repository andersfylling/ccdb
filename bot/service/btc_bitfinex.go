package service

import (
	"strconv"
	"time"
	//"github.com/bwmarrin/discordgo"
	"github.com/Sirupsen/logrus"
	"github.com/s1kx/unison"
)

// btc_bitfinexService updates the bot status to the latest btc price in usd
var BTC_bitfinexService = &unison.Service{
	Name:        "btc Bitfinex watcher",
	Description: "real time monitor of Bitfinex btc:usd rate.",
	Action:      BTC_bitfinexAction,
	Deactivated: false,
	Data: map[string]string{
		"btc_bitfinex_usd": "? USD",
	},
}

func BTC_bitfinexAction(ctx *unison.Context) error {
	//pull data in real time
	c := bitfinex.NewClient()

	// in case your proxy is using a non valid certificate set to TRUE
	c.WebSocketTLSSkipVerify = false

	err := c.WebSocket.Connect()
	if err != nil {
		logrus.Error("Error connecting to web socket : ", err)
	}
	defer c.WebSocket.Close()

	ticker_chan := make(chan []float64)

	c.WebSocket.AddSubscribe(bitfinex.ChanTicker, bitfinex.BTCUSD, ticker_chan)
	go updateStatus(ticker_chan, ctx)

	err = c.WebSocket.Subscribe()
	if err != nil {
		logrus.Error(err)
	}

	return err
}

// keep track of time since last sent update
var lastSent = time.Now().UTC()

func updateStatus(in chan []float64, ctx *unison.Context) {
	for {
		data := <-in
		status := strconv.FormatFloat(data[0], 'f', 2, 64) + " USD"

		// store to service data
		ctx.Bot.SetServiceData("btc Bitfinex watcher", "btc_bitfinex_usd", status)
		//BTC_bitfinexService.Data["btc_bitfinex_usd"] = status

		// `Clients may only update their game status 5 times per minute.`
		if time.Since(lastSent) > 12100 { // ms
			ctx.Bot.Discord.UpdateStatus(0, status)
			lastSent = time.Now().UTC()
		}
	}
}
