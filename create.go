package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func createDir(dir string) (err error) {

	// check if dir exists, if not create it
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	} else {
		fmt.Printf("%s is already a directory. Please choose a different name.\n",
			dir,
		)
		return err
	}
	return nil
}

func CreateProject(rootDir string) {

	// load configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Creating %s with default configuration.\n",
			rootDir,
		)
	}

	// set default values
	defaults := []interface{}{"doc", "bin", "src", "results", "data"}

	// load defaults
	viper.SetDefault("directories", defaults)

	// create root directory
	if err := createDir(rootDir); err != nil {
		os.Exit(1)
	}

	// get list of sub directories from config.toml
	// TODO: set defaults in viper, and allow users
	// to pass a custom template
	directories := viper.Get("directories").([]interface{})

	for _, element := range directories {
		if err := createDir(rootDir + "/" + element.(string)); err != nil {
			os.Exit(1)
		}
	}
}
