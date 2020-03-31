package main

import (
	"log"
	"net/http"

	"github.com/kaitexio/slack_bot_atC/Controller"
)

func main() {
	log.Println("Server start...")
	http.HandleFunc("/message", Controller.MessageController)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
