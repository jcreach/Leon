/*
Copyright © 2025 Julien Creach julien.creach@pm.me
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	basicToken := viper.GetString("basictoken")
	baseAddress := viper.GetString("baseaddress")

	idToDelete, _ := cmd.Flags().GetString("id")

	deleteUrl := baseAddress + "/service/rest/v1/assets/" + idToDelete
	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return
	}
	req.Header.Set("Authorization", basicToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête:", err)
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusNoContent:
		fmt.Printf("Package id : %s deleted !\n", idToDelete)

	}
}
