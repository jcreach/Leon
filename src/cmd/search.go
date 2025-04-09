/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jcreach/Leon/model"
	"github.com/jcreach/Leon/util"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search packages on a nexus repository",
	Run:   searchPackages,
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("name", "n", "", "Value to find in name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func searchPackages(cmd *cobra.Command, args []string) {
	util.CheckConfig()

	basicToken := util.ActiveRepository.BasicToken
	baseAddress := util.ActiveRepository.BaseAddress
	repositoryName := util.ActiveRepository.Name

	nameTofind, _ := cmd.Flags().GetString("name")

	searchUrl := baseAddress + "/service/rest/v1/search?repository=" + repositoryName + "&q=" + nameTofind

	req, err := http.NewRequest("GET", searchUrl, nil)
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
	case http.StatusOK:
		{
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln("Error when reading the query answer:", err)
				return
			}

			var response model.NexusPackageResponse
			errjson := json.Unmarshal([]byte(string(body)), &response)
			if errjson != nil {
				log.Fatalf("Error when deserializing JSON: %v", errjson)
			}

			if len(response.Items) > 0 {
				for _, nexusPackage := range response.Items {
					fmt.Printf("Name: %s, ", nexusPackage.Name)
					for _, asset := range nexusPackage.Assets {
						fmt.Printf("Asset Id: %s, Modification date: %s\n", asset.Id, asset.LastModified)
					}
				}

				fmt.Printf("Next page: %s\n", response.ContinuationToken)
			} else {
				fmt.Println("No entry found...")
			}
		}
	}
}
