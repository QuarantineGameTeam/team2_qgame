package game

import (
	"encoding/json"
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"github.com/QuarantineGameTeam/team2_qgame/game_model"
	"log"
	"strconv"
)

var (
	clans = []string{"red", "green", "blue"}
)

func handleChooseGameQueries(client *api.Client, query api.CallBackQuery) {
	data := query.CallBackData

	gm := getPlayerGame(query.FromUser)
	if gm != nil {
		players := gm.Players
		for _, p := range players {
			if p.Clan == data {
				refuseJoining(client, query)
				return
			}
		}
	}

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

func refuseJoining(client *api.Client, query api.CallBackQuery) {
	err := client.AnswerCallBackQuery(query, fmt.Sprintf("This clan is full. Please, choose another one."), true)
	if err != nil {
		log.Println(err)
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

	if ind == -1 {
		_, err = client.SendMessage(api.Message{
			ChatID: query.FromUser.ID,
			Text: "You are not found in game.",
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	player.Clan = data
	gm.Players[ind] = *player
	gm.Players[ind].SmallPic = fmt.Sprintf("photos/player-%s.png", data)

	err = dbh.Update("players", "clan", data, "player_id", query.FromUser.ID)
	if err != nil {
		log.Println(err)
	}
	err = dbh.Update("players", "smallPic", fmt.Sprintf("photos/player-%s.png", data), "player_id", query.FromUser.ID)
	if err != nil {
		log.Println(err)
	}

	newJSON, err := json.Marshal(gm.Players)
	if err != nil {
		log.Println(err)
	}

	err = dbh.Update("games", "players_json", newJSON, "game_id", gm.GameID)
	if err != nil {
		log.Println(err)
	}


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
	player := getPlayer(user)

	if player == nil {
		log.Println("this player does not exist")
		return
	}

	photoLocation := fmt.Sprintf("temp/%d.png", user.ID)
	err := drawers.CreatePartViewPhoto(gm.Locations, gm.Players, player.X, player.Y, game_model.Horizon, strconv.Itoa(user.ID))
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