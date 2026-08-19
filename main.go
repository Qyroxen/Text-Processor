package main

import (
	"fmt"
	"os"
)

// text_processor - Text processing utilities
func text_processor(path string) {
	fmt.Println("========================================")
	fmt.Println("  Text-Processor")
	fmt.Println("  Text processing utilities")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	text_processor(path)
}
