/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package cmd

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/jcreach/Leon/model"
	"github.com/jcreach/Leon/util"
	"github.com/spf13/cobra"
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
	basicToken := "Basic " + encodedBasicToken

	isLoginValid := checkCredentialsValidity(baseAddress, basicToken)
	if isLoginValid {
		// store new repository
		newRepository := model.Repository{
			Name:        repositoryName,
			BaseAddress: baseAddress,
			BasicToken:  basicToken,
		}

		util.AddOrUpdateRepository(newRepository)
	}
}

func checkCredentialsValidity(baseAddress string, basicToken string) bool {
	statusUrl := baseAddress + "/service/rest/v1/status"
	req, err := http.NewRequest("GET", statusUrl, nil)
	if err != nil {
		log.Fatalln("Error during query creation:", err)
		return false
	}
	req.Header.Set("Authorization", basicToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error during query execution:", err)
		return false
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		{
			log.Println("Logged in successfully!")
			return true
		}
	case http.StatusUnauthorized:
		{
			log.Fatalln("Logged denied! Please check your credetials!")
			return false
		}
	default:
		{
			log.Fatalln("An error as occured!")
			return false
		}
	}
}
