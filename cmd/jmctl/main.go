package main

import (
	"os"
	"time"

	"github.com/jhonatanmacazana/jmctl/internal/commands"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetLevel(log.WarnLevel)

	templateCommand := commands.CreateTemplateCommand()

	app := &cli.App{
		Name:     "jmctl",
		Version:  "0.0.1",
		Compiled: time.Now(),
		Authors:  []*cli.Author{{Name: "Jhonatan Macazana"}},
		Usage:    "Set of useful tools for the CLI to work on various projects.",
		Commands: []*cli.Command{
			templateCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
