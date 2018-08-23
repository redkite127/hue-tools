package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type motionSensor struct {
	Name string
}

var motionSensors map[int]motionSensor

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/usr/local/etc")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	log.Infoln("using config:", viper.ConfigFileUsed())

	viper.SetDefault("frequency_minutes", 1)

	viper.UnmarshalKey("hue_motion_sensors", &motionSensors)
	for k, v := range motionSensors {
		log.Infoln("new motion sensor:", k, v)
	}

	log.SetLevel(log.DebugLevel)
}

func main() {

	// Prepare ticker
	tick := time.Tick(time.Duration(viper.GetInt("frequency_minutes")) * time.Minute)

	// Prepare graceful shutdown
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	log.Debugln("waiting")
Loop:
	for {
		select {
		case <-gracefulStop:
			break Loop
		case <-tick:
			for k, v := range motionSensors {
				// Get temperatures from Hue motion sensors
				t, p, err := getMotionSensorTemperatureAndBattery(k)
				if err != nil {
					log.WithError(err).Infoln("failed to get sensor data:", k)
					continue
				}

				log.Debugf("%v: %2.2f°C %2.2d%%", v.Name, t, p)

				// Send temperatures to home-hub
				if err = sendTemperatureAndBattery(v.Name, t, p); err != nil {
					log.WithError(err).Warningln("failed to send sensor data:", k)
				}
			}
		}
	}

	log.Infoln("application stopped gracefully")
}
