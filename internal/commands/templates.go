package commands

import (
	"fmt"

	"github.com/jhonatanmacazana/jmctl/pkg/utils/displayhelpers"
)

type Template struct {
}

type TemplatesCommand struct {
	AllTemplates bool `short:"a" long:"all" description:"Show builds for the all teams that user has access to"`
	Json         bool `long:"json" description:"Print command result as JSON"`
}

func (command *TemplatesCommand) Execute([]string) error {
	allTemplates, err := command.getAllTemplates()
	if err != nil {
		return nil
	}
	return command.displayTemplates(allTemplates)
}

func (command *TemplatesCommand) getAllTemplates() ([]Template, error) {
	allTemplates := []Template{}
	return allTemplates, nil
}

// func (command *TemplatesCommand) validateCurrentTemplate(currentTeam Template) error {
// 	return nil
// }

func (command *TemplatesCommand) displayTemplates(templates []Template) error {
	var err error
	if command.Json {
		err = displayhelpers.JsonPrint(templates)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println("TODO: display templates")
	return nil

}
