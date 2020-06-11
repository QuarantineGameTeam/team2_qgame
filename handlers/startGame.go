package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/game"
	"log"
)

func StartGames(client *api.Client) {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	games := dbh.GetGames()

	for _, gm := range games {
		if gm.State != game.StateMatchMaking {
			continue
		}

		players := gm.Players
		ready := len(players) > 0
		for _, player := range players {
			if player.Clan == "" {
				ready = false
			}
		}

		if ready {
			gm.State = game.StateRunning
			sendFirstGameMessage(client, gm)
		}
	}
}

func sendFirstGameMessage(client *api.Client, gm *game.Game) {
	players := gm.Players
	for _, player := range players {
		SendCurrentPhoto(client, api.User{ID: player.PlayerId})
	}

	SendMoveButtons(client, api.User{ID: gm.PlayerID})
}

