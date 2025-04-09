/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package model

type Config struct {
	Toto         string
	Repositories []Repository
}

type Repository struct {
	Name        string
	Active      bool
	BaseAddress string
	BasicToken  string
}
