package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"github.com/QuarantineGameTeam/team2_qgame/game"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"log"
)

var (
	clans = []string{"red", "green", "blue"}
)

func handleChooseGameQueries(client *api.Client, query api.CallBackQuery) {
	data := query.CallBackData

	pass := true
	for _, clan := range clans {
		if startsWith(data, clan) {
			pass = false
		}
	}

	if !pass {
		joinClan(client, query, data)
	}
}

func joinClan(client *api.Client, query api.CallBackQuery, data string) {
	var err error
	currGame, err := game.NewGame()
	if err != nil {
		log.Println(err)
	}

	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	err = dbh.InsertGame(currGame.GameJSON, *models.NewPlayer(query.FromUser, 5, 5))
	if err != nil {
		log.Println(err)
	}

	err = client.AnswerCallBackQuery(query, "OK.", false)
	if err != nil {
		log.Println(err)
	}

	err = client.SendMessage(api.Message{
		ChatID: query.FromUser.ID,
		Text: fmt.Sprintf("You entered game in clan of %s!", data),
	})
	if err != nil {
		log.Println(err)
	}

	photoLocation := "../drawers/temp/testpic"
	err = drawers.CreateFullViewPhoto(currGame.Locations, "testpic")
	if err != nil {
		log.Println(err)
	}
	err = client.SendPhoto(query.FromUser.ID, photoLocation)
	if err != nil {
		log.Println(err)
	}
}
