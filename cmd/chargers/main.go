package main

import (
	"github.com/urfave/cli/v2"
	"github.com/zigapk/prpo-chargers/internal/cmd"
	"github.com/zigapk/prpo-chargers/internal/config"
	"github.com/zigapk/prpo-chargers/internal/database"
	"github.com/zigapk/prpo-chargers/internal/logger"
	"log"
	"os"
)

// @title        PRPO Chargers API
// @version      1.0
// @description  PRPO Chargers service api.

// @contact.name   Žiga Patačko Koderman
// @contact.url    https://zerodays.dev
// @contact.email  ziga@zerodays.dev
func main() {
	// Load config
	config.Load()

	// Init logger
	logger.Init()

	// Init database.
	database.Init()
	defer database.Close()

	// Set public key for tokens.
	err := config.Login.LoadKeys()
	if err != nil {
		log.Fatal(err)
	}

	// Create cli app.
	app := cli.NewApp()
	app.Name = "PRPO chargers microservice."
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{{Name: "Žiga Patačko Koderman", Email: "ziga.patacko@gmail.com"}}

	app.Commands = []*cli.Command{
		cmd.Serve,
		cmd.CreateCharger,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
