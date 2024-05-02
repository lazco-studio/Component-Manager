package main

import (
	"embed"
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"Component-Manager/command"
	"Component-Manager/module"
)

//go:embed .cmrc.official.json
var configFile embed.FS

var GITHUB_TOKEN string

func main() {
	os.Setenv("GITHUB_TOKEN", GITHUB_TOKEN)

	officialConfigBytes, err := configFile.ReadFile(".cmrc.official.json")
	if err != nil {
		color.Redln(err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:     "Component-Manager",
		HelpName: "cm",
		Usage:    "A tool for managing JS/TS components and modules.",
		Version:  "v1.3.0",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show the version of cm",
				Action:  command.Version,
			},
			{
				Name:  "init",
				Usage: "initialize a new project",
				Before: func(c *cli.Context) error {
					return module.LoadAppConfig(c, officialConfigBytes)
				},
				Action: command.Init,
			},
			{
				Name:    "add",
				Aliases: []string{"a", "get", "download"},
				Usage:   "add a new component",
				Before: func(c *cli.Context) error {
					return module.LoadAppConfig(c, officialConfigBytes)
				},
				Action: command.Add,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		var errorString = err.Error()

		if errorNumber, err := strconv.Atoi(errorString); err == nil {
			os.Exit(errorNumber)
		}

		color.Redln(errorString)
		os.Exit(1)
	}
}
