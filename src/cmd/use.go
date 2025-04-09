/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jcreach/Leon/util"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Select the active repository",
	Args:  cobra.ExactArgs(1),
	Run:   SelectActiveRepository,
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SelectActiveRepository(cmd *cobra.Command, args []string) {
	if !util.SetActiveRepository(args[0]) {
		log.Fatalln("Repository doesn't exist! Please use login command before ...")
		os.Exit(1)
	}

	fmt.Println(args[0], "is now the active repository 🪴")
}
