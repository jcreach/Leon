/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/jcreach/Leon/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single package by id",
	Run:   deletePackage,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
	deleteCmd.Flags().StringP("id", "i", "", "Identifier of the package to be deleted")

	deleteCmd.MarkFlagRequired("id")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deletePackage(cmd *cobra.Command, args []string) {
	util.CheckConfig()

	basicToken := util.ActiveRepository.BasicToken
	baseAddress := util.ActiveRepository.BaseAddress

	idToDelete, _ := cmd.Flags().GetString("id")

	deleteUrl := baseAddress + "/service/rest/v1/assets/" + idToDelete
	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		log.Fatalln("Error during query creation:", err)
		return
	}
	req.Header.Set("Authorization", basicToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error during query execution:", err)
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusNoContent:
		log.Fatalf("Package id : %s deleted !\n", idToDelete)

	}
}
