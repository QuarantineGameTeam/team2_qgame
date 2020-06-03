package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
)

func HandleUpdate(client *api.Client, update api.Update) {
	if update.HasMessage() {
		handleUpdateMessage(client, update)
	} else if update.HasCallBackQuery() {
		handleUpdateCallBackQuery(client, update)
	}
}