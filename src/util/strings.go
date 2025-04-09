/*
Copyright © 2025 Julien Creach github.com/jcreach
*/
package util

import "strings"

func IsEmpty(value string) bool {
	return len(value) == 0
}

func IsWhiteSpace(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func IsEmptyOrWhiteSpace(value string) bool {
	return IsEmpty(value) || IsWhiteSpace(value)
}
