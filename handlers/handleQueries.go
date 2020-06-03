package handlers

import (
	"log"
	"strings"
	"team2_qgame/api"
)

var (
	startsWith = strings.HasPrefix
)

func handleUpdateCallBackQuery(client *api.Client, update api.Update) {
	handleMainMenuQueries(client, update.CallBackQuery)
}

func handleMainMenuQueries(c *api.Client, q api.CallBackQuery) {
	var err error

	if startsWith(q.CallBackData, "joingame") {
		err = c.AnswerCallBackQuery(q, "Joining the game, please wait...", false)
	} else if startsWith(q.CallBackData, "stats") {
		err = c.AnswerCallBackQuery(q, "Stats...", true)
	} else if startsWith(q.CallBackData, "changenickname") {
		err = c.AnswerCallBackQuery(q, "OK. Type in your nickname.", false)
	}

	if err != nil {
		log.Println(err)
	}
}
