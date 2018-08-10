package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

func sendTemperature(room string, temp float32) error {
	u, err := url.Parse(fmt.Sprintf("http://%v:%v/sensors", viper.Get("home-hub_ip"), viper.Get("home-hub_port")))
	if err != nil {
		return err
	}

	q := u.Query()
	q.Set("room", room)
	q.Set("type", "temperature")
	u.RawQuery = q.Encode()

	body := []byte(fmt.Sprintf("%.2f", temp))

	resp, err := http.Post(u.String(), "text/plain", bytes.NewReader(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to send temperature to home-hub")
	}

	return nil
}
