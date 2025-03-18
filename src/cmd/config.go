/*
Copyright © 2025 Julien Creach julien.creach@pm.me
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show configuration",
	Run:   ShowConfiguration,
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ShowConfiguration(cmd *cobra.Command, args []string) {
	// Lire toutes les configurations
	configs := viper.AllSettings()

	if len(configs) > 0 {
		// Afficher toutes les configurations
		for key, value := range configs {
			fmt.Printf("Key: %s, Value: %v\n", key, value)
		}
	} else {
		fmt.Println("Nexcln, no configuration ...")
	}

}
