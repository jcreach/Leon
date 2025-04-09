/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jcreach/Leon/cmd"
	"github.com/jcreach/Leon/util"
	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(initConfig)
	cmd.Execute()
}

func initConfig() {
	const configFileName string = "config"
	const configFileType string = "yaml"

	homeDrive := os.Getenv("HOMEDRIVE")
	homePath := os.Getenv("HOMEPATH")

	if homeDrive == "" || homePath == "" {
		log.Println("HOMEDRIVE or HOMEPATH environment variables are not defined.")
		return
	}

	// Construction le chemin du répertoire personnel
	homeDir := filepath.Join(homeDrive, homePath, ".leon")

	util.LoadConfig(homeDir, configFileName, configFileType)
}
