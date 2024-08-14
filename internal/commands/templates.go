package commands

import (
	"fmt"

	"github.com/jhonatanmacazana/jmctl/pkg/utils/display"
	"github.com/urfave/cli/v2"
)

type Template struct{}

func CreateTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:    "template",
		Aliases: []string{"te"},
		Usage:   "list available templates",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"A"},
				Usage:   "show builds for the all teams that user has access to",
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "Print command result as JSON",
			},
		},
		Action: templateAction,
	}
}

func templateAction(cCtx *cli.Context) error {
	allTemplates, err := getAllTemplates()
	if err != nil {
		return nil
	}

	jsonFlag := cCtx.Bool("json")

	return displayTemplates(jsonFlag, allTemplates)
}

func getAllTemplates() ([]Template, error) {
	allTemplates := []Template{}
	return allTemplates, nil
}

// func validateCurrentTemplate(currentTeam Template) error {
// 	return nil
// }

func displayTemplates(jsonFlag bool, templates []Template) error {
	var err error
	if jsonFlag {
		err = display.JsonPrint(templates)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println("TODO: display templates")
	return nil

}
