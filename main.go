package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("go run main.go scan <directory>")
		return
	}

	command := os.Args[1]
	path := os.Args[2]

	if command != "scan" {
		fmt.Println("Unknown command")
		return
	}

	report := ScanDirectory(path)

	PrintReport(report)
}
