package game

import (
	"encoding/json"
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"log"
	"strconv"
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
	LocatePlayers(gm)

	err = client.DeleteMessage(query.Message)
	if err != nil {
		log.Println("Unable to delete message: ", err)
	}

	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	player := getPlayer(query.FromUser)
	ind := indexPlayer(gm.Players, *player)
	player.Clan = data
	gm.Players[ind] = *player

	newJSON, err := json.Marshal(gm.Players)
	if err != nil {
		log.Println(err)
	}
	err = dbh.Update("players", "clan", data, "player_id", query.FromUser.ID)
	if err != nil {
		log.Println(err)
	}
	err = dbh.Update("games", "players_json", newJSON, "game_id", gm.GameID)

	_, err = client.SendMessage(api.Message {
		ChatID: query.FromUser.ID,
		Text: "Wait until other users find their clans.",
	})
	if err != nil {
		log.Println(err)
	}
}

func SendCurrentPhoto(client *api.Client, user api.User) {
	gm := getPlayerGame(user)

	photoLocation := fmt.Sprintf("temp/%d.png", user.ID)
	err := drawers.CreateFullViewPhoto(gm.Locations, gm.Players, strconv.Itoa(user.ID))
	if err != nil {
		log.Println(err)
	}

	err = client.SendPhoto(user.ID, photoLocation)
	if err != nil {
		log.Println(err)
	}
}

func SendMoveButtons(client *api.Client, user api.User) {
	_, err := client.SendMessage(api.Message{
		ChatID:       user.ID,
		Text:         "Your turn.",
		InlineMarkup: mainGameMarkup,
	})
	if err != nil {
		log.Println(err)
	}
}