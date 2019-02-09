package main

import (
	"encoding/json"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

func getBitfinexRate() (value float64, err error) {
	price := BitfinexJSON{}
	const u = "https://api.bitfinex.com/v1/pubticker/btcusd"
	if err = getJSON(u, &price); err != nil {
		return 0, err
	}

	if value, err = strconv.ParseFloat(price.LastPrice, 64); err != nil {
		return 0, err
	}

	return
}

func formatValue(value float64) string {
	var t string
	if value > 1000 {
		value /= 1000
		t = "k"
	}

	if value > 1000 {
		value /= 1000
		t = "m"
	}

	return fmt.Sprintf("%.2f%s USD", value, t)
}

func getStatusUpdateBody(price string) interface{} {
	return &disgord.UpdateStatusCommand{
		Since: nil,
		Game: &disgord.Activity{
			Name: price,
			Type: 0,
		},
		Status: disgord.StatusOnline,
		AFK:    false,
	}
}

type bitcoinValueFetcher func() (float64, error)

func statusUpdateScheduler(session disgord.Session, fetch bitcoinValueFetcher, stop chan interface{}) {

	previous := "?"
	for {
		var price string
		value, err := fetch()
		if err != nil && previous[0] != '(' {
			price = "(" + previous + ")"
		} else {
			price = formatValue(value)
		}

		if price != previous {
			err = session.UpdateStatusString(price)
			if err != nil {
				logrus.Error(err)
			} else {
				previous = price
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
