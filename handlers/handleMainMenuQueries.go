package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/config"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/game"
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

	err = client.AnswerCallBackQuery(query, "Wait until we find a game for you.", true)
	if err != nil {
		log.Println(err)
	}

	games := dbh.GetGames()
	fmt.Println(games)

	if len(games) == 0 {
		// Creating new game if there are no games at all
		gm, err := game.NewGame(&query.FromUser)
		if err != nil {
			log.Println(err)
		}

		err = dbh.InsertGame(*gm)
		if err != nil {
			log.Println("Unable to insert game.\n", err)
		}
	} else {
		for _, gm := range games {
			if gm.State == game.StateMatchMaking {
				// adding players to 0x0 location to move after further clan choosing
				gm.Players = append(gm.Players, *models.NewPlayer(query.FromUser, 0, 0))
				if len(gm.Players) >= game.PlayersCount {
					gm.State = game.StateRunning
					sendChooseClanMarkup(client, gm)
				}
			} else {
				fmt.Printf("Game %v is full", *gm)
			}
		}
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
}
