package main

import (
	"fmt"
	"os"
	"simple-chat-app/server/server"

	"github.com/joho/godotenv"
)

/**** Functions ****/

// Load environment variables from ".env" files.
func init() {
	env := os.Args[1]
	cwd, _ := os.Getwd()
	path := cwd + "/app/env/" + env + ".env"
	// path := "env/" + env + ".env"
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Main()
func main() {
	server, err := server.InitializeServer()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	server.Start()
}
