package handlers

import (
	"team2_qgame/api"
)

func HandleUpdate(client *api.Client, update api.Update) {
	if update.HasMessage() {
		handleUpdateMessage(client, update)
	}
}

