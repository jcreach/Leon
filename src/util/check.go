package util

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func CheckConfig() {
	if !viper.IsSet("basictoken") || !viper.IsSet("baseaddress") || !viper.IsSet("repository") {
		fmt.Println("Veuillez éxecuter la commande login avant de pouvoir continuer")
		os.Exit(1)
	}
}
