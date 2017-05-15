package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// subcommands
	// create initiates a new project
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)

	// create subcommands
	createNamePtr := createCommand.String("name", "", "Project name. Creates a new dir for your project (required).")
	createLangPtr := createCommand.String("language", "python", "Project language. Supported langauges are (python|r), defaults to python.")

	// run subcommands
	runAllPtr := runCommand.Bool("all", true, "Run all scripts in project.")
	// check for commands
	if len(os.Args) < 2 {
		fmt.Println("What do you want to do? Try 'help' for options")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCommand.Parse(os.Args[2:])
	case "run":
		runCommand.Parse(os.Args[2:])
	}

	// create command
	if createCommand.Parsed() {
		// required flags
		if *createNamePtr == "" {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		// optional flags
		languageOptions := map[string]bool{"python": true, "r": true}
		if _, validChoice := languageOptions[*createLangPtr]; !validChoice {
			createCommand.PrintDefaults()
			os.Exit(1)
		}
		CreateProject(*createNamePtr)
	}

	// run command
	if runCommand.Parsed() {
		fmt.Printf("Running %s",
			*runAllPtr,
		)
	}

}
