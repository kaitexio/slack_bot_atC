package Controller

import (
	"log"
	"net/http"
	"os"

	"github.com/kaitexio/slack_bot_atC/goquery"
	"github.com/kaitexio/slack_bot_atC/slack"
)

var (
	CrawlURL = os.Getenv("CrawlURL")
	SlackAPIToken = os.Getenv("SlackAPIToken")
	SlackChannel  = os.Getenv("SlackChannel")
)

func MessageController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	case http.MethodPost:
		tem, err := goquery.RequestGoquery(CrawlURL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			w.WriteHeader(http.StatusBadRequest) //400
		}
		if err := slack.PostMessages(SlackAPIToken, SlackChannel, tem); err != nil {
			log.Fatalf("failed to serve: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK) // 200
	default:
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
