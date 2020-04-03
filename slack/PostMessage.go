package slack

import (
	"log"

	"github.com/nlopes/slack"
)

func PostMessages(slackAPIToken, slackChannel string, tmp map[int]string) (err error) {
	client := slack.New(slackAPIToken)
	for i, _ := range tmp {
		text := tmp[i]
		_, _, err := client.PostMessage(slackChannel,
			slack.MsgOptionText(text, false),
		)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			panic(err)
		}
	}

	return err
}
