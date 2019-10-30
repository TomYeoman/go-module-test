package main

import "github.com/fatih/color"

// Hi Returns hello world

func main() {
	color.Cyan("Main text in cyan.")
}

func Hi() string {
	color.Cyan("Prints text in cyan.")

	return "Hello world"
}
