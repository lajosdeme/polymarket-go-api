package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Polymarket Go API SDK")
	fmt.Println("This is an SDK library. See examples/ directory for usage examples.")
	fmt.Println("Documentation: https://github.com/lajosdeme/polymarket-go-api")

	// If run with specific arguments, show examples
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "examples":
			fmt.Println("\nTo run examples:")
			fmt.Println("  go run examples/basic_usage.go")
		case "docs":
			fmt.Println("\nDocumentation:")
			fmt.Println("  README.md - Complete documentation")
			fmt.Println("  examples/basic_usage.go - Usage examples")
		case "version":
			fmt.Println("Version: 1.0.0")
		default:
			fmt.Println("\nUsage:")
			fmt.Println("  go run . [examples|docs|version]")
		}
	}
}
