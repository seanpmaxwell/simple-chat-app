package main

import (
	"log"

	"github.com/seanpmaxwell/simple-chat-app/server/simplechat"
)

func main() {
	simplechat.Init()
	log.Printf("%s %v listening...",
		simplechat.Name,
		simplechat.Version,
	)
	err := simplechat.NewServer().Start()
	if err != nil {
		log.Fatalf("%s exited with error: %+v", simplechat.Name, err)
	}
	log.Printf("%s exited.", simplechat.Name)
}
