package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/config"
	"github.com/QuarantineGameTeam/team2_qgame/database"
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

	err = client.AnswerCallBackQuery(query, "OK", false)
	if err != nil {
		log.Println(err)
	}

	err = client.SendMessage(api.Message{
		ChatID: query.FromUser.ID,
		Text: "Now choose the clan, you want to play in.",
		InlineMarkup: chooseClanMarkup,
	})
	if err != nil {
		log.Println(err)
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

	err = client.SendMessage(api.Message{
		ChatID: query.FromUser.ID,
		Text: "OK. Now send me your new nickname.",
	})

	err = dbh.Update("users", "state", config.StateChangingName, "telegram_id", query.FromUser.ID)
	if err != nil {
		log.Println(err)
	}
}


