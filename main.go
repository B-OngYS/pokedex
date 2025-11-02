package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	output := strings.Fields(text)
	return output
}

func main() {
	fmt.Println("Hello, World!")
}
