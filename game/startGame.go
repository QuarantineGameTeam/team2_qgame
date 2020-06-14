package game

import (
	"encoding/json"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/game_model"
	"log"
)

func StartGames(client *api.Client) {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	games := dbh.GetGames()

	for _, gm := range games {
		if gm.State != game_model.StateChoosingClan {
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
			for _, p := range players {
				_, err = client.SendMessage(api.Message{
					ChatID: p.PlayerId,
					Text: "Game is starting. Be ready.",
				})
			}

			gm.State = game_model.StateRunning
			err = dbh.Update("games", "state", gm.State, "game_id", gm.GameID)
			if err != nil {
				log.Println(err)
			}

			game_model.LocatePlayers(gm)
			playersJSON, err := json.Marshal(gm.Players)
			if err != nil {
				log.Println(err)
			}

			err = dbh.Update("games", "players_json", playersJSON, "game_id", gm.GameID)
			if err != nil {
				log.Println(err)
			}
			for _, p := range gm.Players {
				err = dbh.Update("players", "x", p.X, "player_id", p.PlayerId)
				if err != nil {
					log.Println(err)
				}
				err = dbh.Update("players", "y", p.Y, "player_id", p.PlayerId)
				if err != nil {
					log.Println(err)
				}
			}

			sendFirstGameMessage(client, gm)
		}
	}
}

func sendFirstGameMessage(client *api.Client, gm *game_model.Game) {
	players := gm.Players
	for _, player := range players {
		SendCurrentPhoto(client, api.User{ID: player.PlayerId})
	}

	SendMoveButtons(client, api.User{ID: gm.PlayerID})
}

