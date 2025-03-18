/*
Copyright © 2025 Julien Creach julien.creach@pm.me
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jcreach/Leon/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println("Les variables d'environnement HOMEDRIVE ou HOMEPATH ne sont pas définies.")
		return
	}

	// Construction le chemin du répertoire personnel
	homeDir := filepath.Join(homeDrive, homePath, ".leon")

	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)
	viper.AddConfigPath(homeDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found;
			createMissingConfigFile(homeDir + "/" + configFileName + "." + configFileType)
		} else {
			// Config file was found but another error was produced
			fmt.Fprintln(os.Stderr, "Error reading config file:", err)
			os.Exit(1)
		}
	}
}

func createMissingConfigFile(fileName string) {

	dirPath := filepath.Dir(fileName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Créer le dossier s'il n'existe pas
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Erreur lors de la création du dossier :", err)
			return
		}
		fmt.Println("Dossier créé avec succès.")
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Créer le fichier s'il n'existe pas
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Erreur lors de la création du fichier :", err)
			return
		}
		defer file.Close()
		fmt.Println("Fichier créé avec succès.")
	}
}
