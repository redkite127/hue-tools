package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type temperatureSensorState struct {
	Temperature int
	LastUpdated string
}

type temperatureSensorConfig struct {
	Battery int
}

type temperatureSensor struct {
	State  temperatureSensorState
	Config temperatureSensorConfig
}

func getMotionSensorTemperature(id int) (float32, error) {
	url := fmt.Sprintf("http://%v/api/%v/sensors/%v", viper.Get("hue_bridge_ip"), viper.Get("hue_bridge_key"), id)

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	var ts temperatureSensor
	if err := json.NewDecoder(resp.Body).Decode(&ts); err != nil {
		return 0.0, err
	}

	return float32(ts.State.Temperature) / 100, nil
}

func getMotionSensorTemperatureAndBattery(id int) (float32, int, error) {
	url := fmt.Sprintf("http://%v/api/%v/sensors/%v", viper.Get("hue_bridge_ip"), viper.Get("hue_bridge_key"), id)

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, 0, err
	}
	defer resp.Body.Close()

	var ts temperatureSensor
	if err := json.NewDecoder(resp.Body).Decode(&ts); err != nil {
		return 0.0, 0, err
	}

	return float32(ts.State.Temperature) / 100, ts.Config.Battery, nil
}
