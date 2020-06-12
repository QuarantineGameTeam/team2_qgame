package game

import (
	"encoding/json"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/config"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/game_model"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"log"
)

func handleMainMenuQueries(client *api.Client, query api.CallBackQuery) {
	var err error

	if startsWith(query.CallBackData, "joingame") {
		handleJoinGameQuery(client, query)
	} else if startsWith(query.CallBackData, "stats") {
		err = client.AnswerCallBackQuery(query, "Stats...", true)
	} else if startsWith(query.CallBackData, "changenickname") {
		handleChangeNickNameQuery(client, query)
	}

	if err != nil {
		log.Println(err)
	}

}

func handleJoinGameQuery(client *api.Client, query api.CallBackQuery) {
	var err error

	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println("Failed to connect to database.\n", err)
	}

	game := getPlayerGame(query.FromUser)
	if game != nil {
		err = client.AnswerCallBackQuery(query, "You are already in a game! GRRR.", true)
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = client.AnswerCallBackQuery(query, "Wait until we find a game for you.", true)
	if err != nil {
		log.Println(err)
	}

	games := dbh.GetGames()

	if len(games) == 0 {
		// Creating new game if there are no games at all
		gm, err := game_model.NewGame(&query.FromUser)
		if err != nil {
			log.Println(err)
		}
		
		p := models.NewPlayer(query.FromUser, 0, 0)
		err = dbh.InsertPlayer(*p)
		if err != nil {
			log.Println(err)
		}

		err = dbh.InsertGame(*gm)
		if err != nil {
			log.Println("Unable to insert game.\n", err)
		}
	} else {
		// else joining the game which is in Matchmaking state
		joined := false
		for _, gm := range games {
			if gm.State == game_model.StateMatchMaking {
				// adding players to 0x0 location to move after further clan choosing
				p := models.NewPlayer(query.FromUser, 1, 0)
				err = dbh.InsertPlayer(*p)
				if err != nil {
					log.Println(err)
				}

				gm.Players = append(gm.Players, *p)

				bytes, err := json.Marshal(gm.Players)
				if err != nil {
					log.Println(err)
				}
				gm.PlayersJSON = string(bytes)
				err = dbh.Update("games", "players_json", gm.PlayersJSON, "game_id", gm.GameID)
				if err != nil {
					log.Println("Unable to update table games: ", err)
				}

				if len(gm.Players) >= game_model.PlayersCount {
					gm.State = game_model.StateRunning
					sendChooseClanMarkup(client, gm)
				}
				joined = true
				break
			}
		}
		if !joined {
			// Creating new game if there are no opened games
			gm, err := game_model.NewGame(&query.FromUser)
			if err != nil {
				log.Println(err)
			}

			err = dbh.InsertGame(*gm)
			if err != nil {
				log.Println("Unable to insert game.\n", err)
			}
		}
	}

	err = client.DeleteMessage(query.Message)
	if err != nil {
		log.Println("Unable to delete message: ", err)
	}
}

func handleChangeNickNameQuery(client *api.Client, query api.CallBackQuery) {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	err = client.AnswerCallBackQuery(query, "OK.", false)
	if err != nil {
		log.Println(err)
	}

	_, err = client.SendMessage(api.Message{
		ChatID: query.FromUser.ID,
		Text:   "OK. Now send me your new nickname.",
	})

	err = dbh.Update("users", "state", config.StateChangingName, "telegram_id", query.FromUser.ID)
	if err != nil {
		log.Println(err)
	}

	err = client.DeleteMessage(query.Message)
	if err != nil {
		log.Println("Unable to delete message: ", err)
	}
}
