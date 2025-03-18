/*
Copyright © 2025 Julien Creach julien.creach@pm.me
*/
package model

type NexusPackageResponse struct {
	Items             []NexusPackage `json:"items"`
	ContinuationToken string         `json:"continuationToken"`
}
