package commands

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const version = "DEV ((REPLACE AT BUILD))"

var commands []*cli.Command = []*cli.Command{}

func Register(command *cli.Command) {
	commands = append(commands, command)
}

var commander cli.App = cli.App{
	Name:        "vcss",
	Usage:       "Virtual OCPP Charge Station Server",
	Description: "This server provides a function to create virtual OCPP charge stations",
	Version:     version,
	Authors:     []*cli.Author{{Name: "Nina Siegel", Email: "contact@l1f.de"}},
}

func Run() {
	commander.Commands = commands
	commander.Run(os.Args)
}

func getLogger() (*zap.Logger, error) {
	// TODO: return correct logger for dev and prod

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("Can't create logger: %w", err)
	}

	return logger, nil
}
