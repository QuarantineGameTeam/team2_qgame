package handlers

import (
<<<<<<< HEAD
	"github.com/QuarantineGameTeam/team2_qgame/api"
=======
	"team2_qgame/api"
>>>>>>> f42b3f3afd86bec62aa8bc6a094df7066142ab24
)

func HandleUpdate(client *api.Client, update api.Update) {
	if update.HasMessage() {
		handleUpdateMessage(client, update)
<<<<<<< HEAD
	} else if update.HasCallBackQuery() {
		handleUpdateCallBackQuery(client, update)
	}
}
=======
	}
}

>>>>>>> f42b3f3afd86bec62aa8bc6a094df7066142ab24
