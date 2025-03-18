/*
Copyright © 2025 Julien Creach julien.creach@pm.me
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login in to your nexus instance.",
	Run:   Login,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	loginCmd.Flags().StringP("password", "p", "", "Password for login")
	loginCmd.Flags().StringP("address", "a", "", "Url of the nexus repository")
	loginCmd.Flags().StringP("repository", "r", "", "Repository name")

	// Mark flags as required
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
	loginCmd.MarkFlagRequired("address")
	loginCmd.MarkFlagRequired("repository")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Login(cmd *cobra.Command, args []string) {

	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	baseAddress, _ := cmd.Flags().GetString("address")
	repositoryName, _ := cmd.Flags().GetString("repository")

	encodedBasicToken := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	statusUrl := baseAddress + "/service/rest/v1/status"

	req, err := http.NewRequest("GET", statusUrl, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return
	}

	basicToken := "Basic " + encodedBasicToken
	req.Header.Set("Authorization", basicToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête:", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		{
			viper.Set("basictoken", basicToken)
			viper.Set("baseaddress", baseAddress)
			viper.Set("repository", repositoryName)

			err := viper.WriteConfig()
			if err != nil {
				log.Fatalf("Erreur lors de l'écriture du fichier de configuration : %v\n", err)
			}
			fmt.Println("Logged in successfully!")
		}
	case http.StatusUnauthorized:
		fmt.Println("Logged denied! Please check your credetials!")
	default:
		fmt.Println("An error as occured!")
	}
}
