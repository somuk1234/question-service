package commands

import (
	"context"

	"github.com/urfave/cli"

	"questions-keeper-service/internal/server"
)

// StartCommand is used to register the start cli command
var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts HTTP server",
	Action:  startAction,
}

// startAction start the http server and initializes the daemon
func startAction(ctx *cli.Context) error {
	backgroundCtx := context.Background()

	server.Start(backgroundCtx)
	return nil
}
