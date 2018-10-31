package main

import (
	"encoding/json"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type BitfinexJSON struct {
	// https://api.bitfinex.com/v1/pubticker/btcusd
	LastPrice string `json:"last_price"`
}
var myClient = &http.Client{Timeout: 7 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getBitfinexRate() (value string) {
	price := BitfinexJSON{}
	err := getJSON("https://api.bitfinex.com/v1/pubticker/btcusd", &price)
	if err != nil {
		logrus.Error(err)
		value = "?"
	} else {
		value = price.LastPrice + " USD"
	}
	return
}

func getStatusUpdateBody(price string) interface{} {
	return &disgord.UpdateStatusCommand{
		Since: nil,
		Game: &disgord.Activity{
			Name: price,
			Type: 0,
		},
		Status: disgord.StatusOnline,
		AFK: false,
	}
}

type bitcoinValueFetcher func() string
func statusUpdateScheduler(session disgord.Session, fetch bitcoinValueFetcher,stop chan interface{}) {

	previous := "?"
	for {
		price := fetch()
		if price != previous {
			if price == "?" {
				price = "(" + previous + ")"
			} else {
				previous = price
			}

			newStatus := getStatusUpdateBody(price)
			data, _ := json.Marshal(newStatus)
			fmt.Println(string(data))
			err := session.Emit(disgord.CommandUpdateStatus, newStatus)
			if err != nil {
				logrus.Error(err)
			}
		}

		select {
		case <-stop:
			logrus.Info("Stopping CCDB service")
			return
		case <-time.After(12 * time.Second):
		}
	}
}