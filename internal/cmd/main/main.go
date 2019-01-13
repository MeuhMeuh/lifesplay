package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/meuhmeuh/lifesplay/internal/app/lifesplay"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Very optimistic way to start the app.
	var err error

	box := packr.NewBox("../../../config")

	if _, err := os.Stat("../../../resources/logs"); os.IsNotExist(err) {
		os.MkdirAll("../../../resources/logs", 0777)
	}

	// Logs
	var log = logrus.New()
	logFile, err := os.OpenFile("../../../resources/logs/lifesplay.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Errorf("Something went wrong when opening the log file"))
	}
	log.Out = logFile
	log.Info("Hello world")

	// Config file
	configurationBytes, err := box.Find("lifesplay.json")
	if err != nil {
		panic(fmt.Errorf("Could not find the json bytes in file : %s", err))
	}

	viper.SetConfigType("json")
	viper.ReadConfig(bytes.NewBuffer(configurationBytes))
	// err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Could not read the configuration file : %s", err))
	}

	lifesplay := &lifesplay.Lifesplay{
		Debug:  flag.Bool("d", false, "enables the debug mode"),
		Logger: log,
	}

	lifesplay.Initialize()
	lifesplay.Start()
}
