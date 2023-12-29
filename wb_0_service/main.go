package main

import (
	"fmt"
	"github.com/andrey-tsyn/wb_0/app"
	"github.com/andrey-tsyn/wb_0/app/database"
	"github.com/andrey-tsyn/wb_0/configuration"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	cfg := configuration.Config{}
	if err := env.Parse(&cfg); err != nil {
		panic("Can't read config.")
	}

	initLogger(cfg)

	db, err := database.ConnectToDb(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName),
	)
	if err != nil {
		log.Fatal(err)
	}

	app.Start(cfg, db)
}

func initLogger(cfg configuration.Config) {
	log.SetOutput(os.Stdout)

	if cfg.Environment == "dev" {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if lvl, err := log.ParseLevel(cfg.LogLevel); err != nil {
		log.SetLevel(log.InfoLevel)
		log.Warn("Log level is incorrect, switching to info log mode.")
	} else {
		log.SetLevel(lvl)
	}
}
