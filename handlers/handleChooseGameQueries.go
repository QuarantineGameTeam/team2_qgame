package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"github.com/QuarantineGameTeam/team2_qgame/game"
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

	err = client.AnswerCallBackQuery(query, fmt.Sprintf("You entered game in clan of %s!", data), true)
	if err != nil {
		log.Println(err)
	}

	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	// Go for determining player's game
	gm, err := dbh.GetGameByID(7)
	if err != nil || gm == new(game.Game) {
		log.Println("Error getting game with index 7.\n",err)
	}

	gm.LocatePlayers()

	photoLocation := "temp/testpic.png"
	err = drawers.CreateFullViewPhoto(gm.Locations, "testpic")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Game is: ", *gm)

	for _, p := range gm.Players {
		err = client.SendPhoto(p.PlayerId, photoLocation)
		if err != nil {
			log.Println(err)
		}
	}

	// Sending move buttons one by one for players
	//_, err = client.SendMessage(api.Message {
	//	ChatID: query.FromUser.ID,
	//	Text: "Your turn.",
	//	InlineMarkup: mainGameMarkup,
	//})
	//if err != nil {
	//	log.Println(err)
	//}
}
