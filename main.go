package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("qrop-coordinator starting...")
	fmt.Printf("version: %s\n", version())
	os.Exit(0)
}

func version() string {
	return "0.1.0-dev"
}
