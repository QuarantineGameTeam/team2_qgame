package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"log"
)

var (
	clans = []string{"red", "green", "blue"}
)

func handleChooseGameQueries(client *api.Client, query api.CallBackQuery) {
	data := query.CallBackData

	pass := true
	for _, clan := range clans {
		if data == clan {
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

	// Go for determining player's game
	gm := getPlayerGame(query.FromUser)

	// Positioning them near the castles
	gm.LocatePlayers()

	err = client.DeleteMessage(query.Message)
	if err != nil {
		log.Println("Unable to delete message: ", err)
	}

	photoLocation := "temp/testpic.png"
	err = drawers.CreateFullViewPhoto(gm.Locations, gm.Players, "testpic")
	if err != nil {
		log.Println(err)
	}

	err = client.SendPhoto(query.FromUser.ID, photoLocation)
	if err != nil {
		log.Println(err)
	}

	// WEIRD ASYNC RIGHT NOW
	// Sending move buttons one by one for players
	_, err = client.SendMessage(api.Message{
		ChatID:       query.FromUser.ID,
		Text:         "Your turn.",
		InlineMarkup: mainGameMarkup,
	})
	if err != nil {
		log.Println(err)
	}
}
