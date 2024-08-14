package commands

import (
	"github.com/urfave/cli/v2"
)

type App struct{}

func CreateAppCommand() *cli.Command {
	return &cli.Command{
		Name:   "app",
		Usage:  "manage configured apps",
		Action: appAction,
	}
}

func appAction(cCtx *cli.Context) error {
	allTemplates, err := getAllTemplates()
	if err != nil {
		return nil
	}

	jsonFlag := cCtx.Bool("json")

	return displayTemplates(jsonFlag, allTemplates)
}
