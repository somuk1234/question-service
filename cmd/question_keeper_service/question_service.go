package main

import (
	"log"
	"os"
	"questions-keeper-service/internal/commands"

	"github.com/urfave/cli"
)

// @title question service API
// @version 1.0
// @description service for adding coding questions and analysing
// @BasePath /api/
func main() {
	app := cli.NewApp()
	app.Name = "Question Keeper Service Backend"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		commands.StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Print(err)
	}
}
