package main

import (
	"database/sql"
	"fmt"
	"github.com/metaleaf-io/gator/conf"
	"github.com/metaleaf-io/gator/database"
	"github.com/metaleaf-io/gator/services"
	"github.com/metaleaf-io/log"
	"os"
	"os/signal"
	"syscall"
)

const (
	configPath  = "./gator.yaml"
	migratePath = "./db/migrate"
)

var (
	appConfig   *conf.AppConfig
	db          *sql.DB
)

func main() {
	appConfig = conf.LoadConfig(configPath)
	db = database.Connect(appConfig)
	database.Migrate(db, migratePath)
	services.StartListener(appConfig, db)
	waitForExit()
}

func waitForExit() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signals
	db.Close()
	log.Info(fmt.Sprintf("Exiting on signal: %v", sig))
}
