/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package model

type NexusPackage struct {
	Name   string              `json:"name"`
	Assets []NexusPackageAsset `json:"assets"`
}

type NexusPackageAsset struct {
	LastModified string `json:"lastModified"`
	Id           string `json:"id"`
}
