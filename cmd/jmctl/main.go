package main

import (
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"github.com/jhonatanmacazana/jmctl/internal/commands"
	"github.com/jhonatanmacazana/jmctl/pkg/utils/ui"
)

func main() {
	log.SetLevel(log.WarnLevel)

	parser := goflags.NewParser(&commands.JmCtl, goflags.HelpFlag|goflags.PassDoubleDash)
	parser.NamespaceDelimiter = "-"

	helpParser := goflags.NewParser(&commands.JmCtl, goflags.HelpFlag)
	helpParser.NamespaceDelimiter = "-"

	args, err := parser.Parse()
	handleError(helpParser, err)

	if len(args) > 0 {
		fmt.Printf("Unknown argument '%s'.\n", args[0])
		os.Exit(1)
	}
}

func handleError(helpParser *goflags.Parser, err error) {
	if err != nil {
		if err == commands.ErrShowHelpMessage {
			showHelp(helpParser)
		} else if flagsErr, ok := err.(*goflags.Error); ok && flagsErr.Type == goflags.ErrCommandRequired {
			showHelp(helpParser)
		} else if flagsErr, ok := err.(*goflags.Error); ok && flagsErr.Type == goflags.ErrHelp {
			fmt.Println(err)
			os.Exit(0)
		} else {
			fmt.Fprintf(ui.Stderr, "error: %s\n", err)
		}

		flagError := err.(*goflags.Error)
		if flagError.Type == goflags.ErrHelp {
			os.Exit(0)
		}

		if flagError.Type == goflags.ErrUnknownFlag {
			fmt.Println("\nUse --help to view all available options.")
			os.Exit(1)
		}

		fmt.Println("\nUse --help to view all available options.")
		os.Exit(1)
	}
}

func showHelp(helpParser *goflags.Parser) {
	helpParser.ParseArgs([]string{"-h"})
	helpParser.WriteHelp(os.Stdout)
	os.Exit(0)
}
