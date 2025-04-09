/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package model

type NexusPackageResponse struct {
	Items             []NexusPackage `json:"items"`
	ContinuationToken string         `json:"continuationToken"`
}
