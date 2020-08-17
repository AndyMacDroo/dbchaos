package main

import (
	. "dbchaos/dbchaos/config"
	. "dbchaos/dbchaos/strategy"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	databaseConfig, err := DatabaseConfiguration()
	if err != nil {
		panic(err)
	}
	options, err := ConfigurationOptions()
	if err != nil {
		panic(err)
	}

	chaos := &ChaosStrategy{
		StrategyName: ConnectionLeak,
		ChaosOptions: options,
		DatabaseConfiguration: databaseConfig,
	}
	for {
		chaos.Execute()
	}
}
