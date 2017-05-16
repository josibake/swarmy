package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "swarmy"
	app.Usage = "the data science swiss army knife"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Josiah Baker",
			Email: "@josibake",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Create a project template inside a new directory",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Usage: "Load project configuration from `FILE`",
				},
				cli.StringFlag{
					Name:  "lang, l",
					Usage: "Default language for project. Supported values are python, r",
					Value: "python",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("created project: ", c.Args().First())
				CreateProject(c.Args().First())
				return nil
			},
		},
		{
			Name:  "run",
			Usage: "Run scripts. Can be used to run all or a subset of scripts in a project",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
