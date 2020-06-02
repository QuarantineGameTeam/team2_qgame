package main

import (
	"log"
	"team2_qgame/api"
	"team2_qgame/handlers"
	"time"
)

var (
	client *api.Client
)

func main() {
	// Setting up telegram bot client
	var err error
	client, err = api.NewClient(botToken)
	if err != nil {
		log.Fatal(err)
	}

	firstUpdate := 0
	lastUpdate := 0
	var update api.Update

	for {
		updates := client.GetUpdates(lastUpdate + 1)
		if len(updates) != 0 {
			update = updates[0]

			if firstUpdate == 0 {
				firstUpdate = updates[0].UpdateID
			}
			lastUpdate = updates[0].UpdateID

			if update.UpdateID != 0 {
				// run handlers asynchronously
				go handlers.HandleUpdate(client, update)
			}
		}

		time.Sleep(time.Millisecond * 100)
	}

}
