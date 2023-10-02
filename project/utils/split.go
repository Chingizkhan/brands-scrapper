package utils

import (
	"log"
	"strings"
)

func GetProductString(str string) []string {
	prod := make([]string, 0, 5)

	text := strings.Split(str, "\n")
	for _, t := range text {
		if str := strings.TrimSpace(t); str != "" {
			prod = append(prod, str)
		}
	}

	log.Printf("prod %#v", prod)

	return prod
}
