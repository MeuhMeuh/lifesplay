package main

import (
	"flag"
	"fmt"

	"github.com/meuhmeuh/lifesplay/internal/app/lifesplay"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("lifesplay")
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Could not read the configuration file : %s", err))
	}

	lifesplay := &lifesplay.App{
		Debug: flag.Bool("d", false, "enables the debug mode"),
	}

	lifesplay.Initialize()

	lifesplay.Start()
}
