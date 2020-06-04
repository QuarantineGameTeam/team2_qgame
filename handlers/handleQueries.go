package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/config"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"log"
	"strings"
)

var (
	startsWith = strings.HasPrefix
)

func handleUpdateCallBackQuery(client *api.Client, update api.Update) {
	handleMainMenuQueries(client, update.CallBackQuery)
}

func handleMainMenuQueries(client *api.Client, query api.CallBackQuery) {
	var err error

	if startsWith(query.CallBackData, "joingame") {
		err = client.AnswerCallBackQuery(query, "Joining the game, please wait...", false)
	} else if startsWith(query.CallBackData, "stats") {
		err = client.AnswerCallBackQuery(query, "Stats...", true)
	} else if startsWith(query.CallBackData, "changenickname") {
		handleChangeNickNameQuery(client, query)
	}

	if err != nil {
		log.Println(err)
	}

}

func handleChangeNickNameQuery(client *api.Client, query api.CallBackQuery) {
	dbh := database.NewDBHandler()

	err := client.AnswerCallBackQuery(query, "OK. Type in your nickname.", false)
	if err != nil {
		log.Println(err)
	}

	dbh.Update("users", "state", config.StateChangingName, "telegram_id", query.FromUser.ID)
}
