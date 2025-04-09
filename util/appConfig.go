/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jcreach/Leon/model"
	"github.com/spf13/viper"
)

var appConfig model.Config
var ActiveRepository model.Repository

func AddOrUpdateRepository(repository model.Repository) {
	index := doesRepositoryAlreadyExist(repository.Name)

	if index > -1 {
		appConfig.Repositories[index] = repository
	} else {
		if len(appConfig.Repositories) == 0 {
			repository.Active = true
		}

		appConfig.Repositories = append(appConfig.Repositories, repository)
	}

	saveRepositories()
	loadActiveRepository()
}

func AnyActiveRepository() bool {
	var result = false

	for _, repository := range appConfig.Repositories {
		if repository.Active {
			return true
		}
	}

	return result
}

func LoadConfig(homeDir string, filename string, configType string) {
	viper.AddConfigPath(homeDir)
	viper.SetConfigName(filename)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found;
			createMissingConfigFile(homeDir + "/" + filename + "." + configType)
		} else {
			// Config file was found but another error was produced
			log.Fatalln("Error reading config file:", err)
			os.Exit(1)
		}
	}

	viper.Unmarshal(&appConfig)
	loadActiveRepository()
}

func SetActiveRepository(name string) bool {

	if doesRepositoryAlreadyExist(name) == -1 {
		return false
	}

	for i, repository := range appConfig.Repositories {
		if repository.Name == name {
			appConfig.Repositories[i].Active = true
		} else {
			appConfig.Repositories[i].Active = false
		}
	}

	saveRepositories()
	loadActiveRepository()
	return true
}

func ShowLoggedInRepositories() {
	fmt.Println("")
	fmt.Println("Configured repositories :")
	fmt.Println("🪴  : active config")
	fmt.Println("🪟  : inactive config")
	fmt.Println("")

	status := "🪟"
	for _, repository := range appConfig.Repositories {
		if repository.Active {
			status = "🪴"
		}
		fmt.Println("name :", repository.Name)
		fmt.Println("  url :", repository.BaseAddress)
		fmt.Println("  Status :", status)
		fmt.Println("")
	}
}

func createMissingConfigFile(fileName string) {

	dirPath := filepath.Dir(fileName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Créer le dossier s'il n'existe pas
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalln("Folder creation error :", err)
			return
		}
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Créer le fichier s'il n'existe pas
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalln("Error during file creation :", err)
			return
		}
		defer file.Close()
	}
}

// Check if a Repository already exist in the config and return its index
// otherwise it return -1
func doesRepositoryAlreadyExist(name string) int {
	index := -1

	if len(appConfig.Repositories) == 0 {
		return index
	}

	return getRepositoryIndex(appConfig.Repositories, name)
}

func getRepositoryIndex(repositories []model.Repository, name string) int {
	index := -1
	for i, repository := range repositories {
		if repository.Name == name {
			index = i
			break
		}
	}

	return index
}

func loadActiveRepository() {
	for _, repository := range appConfig.Repositories {
		if repository.Active {
			ActiveRepository = repository
			return
		}
	}
}

func saveRepositories() {
	viper.Set("repositories", appConfig.Repositories)
	viper.WriteConfig()
}
