package main

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/config"
	"github.com/QuarantineGameTeam/team2_qgame/game"
	"github.com/QuarantineGameTeam/team2_qgame/handlers"
	"log"
	"time"
)

var (
	client *api.Client
)

func main() {
	// Setting up telegram bot client
	var err error
	client, err = api.NewClient(config.BotToken)
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

		game.GameUpdate()

		time.Sleep(time.Millisecond * 100)
	}

}
