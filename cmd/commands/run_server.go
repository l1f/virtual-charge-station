package commands

import (
	"fmt"

	"github.com/l1f/virtual-charge-station/internal/app"
	"github.com/l1f/virtual-charge-station/internal/router"
	"github.com/urfave/cli/v2"
)

const (
	host = "localhost"
	port = 5000
)

func init() {
	Register(&command)
}

var command cli.Command = cli.Command{
	Name:    "start",
	Aliases: []string{"start_server", "run", "fly", "s"},

	Description: "Starts the server the webserver to control the virtual OCPP Stations",

	Flags: []cli.Flag{&cli.StringFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Required: true,
		Usage:    "The path to the config file",
	}},

	Action: startServer,
}

func startServer(ctx *cli.Context) error {
	configPath := ctx.String("config")
	fmt.Printf("TODO: load and read config: %s\n", configPath)

	logger, err := getLogger()
	if err != nil {
		return fmt.Errorf("Can't initialise webserver start: %w", err)
	}

	application := app.App{
		Logger: logger.Sugar().Named("web"),
	}

	return router.Start(&application)
}
