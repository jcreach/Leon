/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package util

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func CheckConfig() {
	if !viper.IsSet("repositories") || !AnyActiveRepository() {
		log.Fatalln("Please execute the login command before continuing..")
		os.Exit(1)
	}
}
