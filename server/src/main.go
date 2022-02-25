package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load environment variables from ".env" files.
func init() {
	env := os.Args[1]
	cwd, _ := os.Getwd()
	path := cwd + "/env/" + env + ".env"
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Main()
func main() {
	server, err := InitializeServer()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	server.Start()
}
