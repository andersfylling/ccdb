package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

type BitfinexJSON struct {
	// https://api.bitfinex.com/v1/pubticker/btcusd
	LastPrice string `json:"last_price"`
}

var myClient = &http.Client{Timeout: 5 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getLatestPrice() string {
	price := BitfinexJSON{}
	err := getJSON("https://api.bitfinex.com/v1/pubticker/btcusd", &price)

	if err != nil {
		fmt.Println(err.Error())
		return "? USD"
	}

	return price.LastPrice
}

// BTCBitfinexAction service action
func BTCBitfinexAction(ctx *unison.Context) error {
	go func() {
		// update the status
		ctx.Bot.Discord.UpdateStatus(0, getLatestPrice()+" USD")
		for {
			select {
			case <-ctx.SystemInteruptChan:
				fmt.Println("\tStopped Service: BTCBitfinex")
				return
			default:
				time.Sleep(12 * time.Second)
				// update the status again
				ctx.Bot.Discord.UpdateStatus(0, getLatestPrice()+" USD")
			}
		}
	}()

	return nil
}
