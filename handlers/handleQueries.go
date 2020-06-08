package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"strings"
)

var (
	startsWith = strings.HasPrefix
)

func handleUpdateCallBackQuery(client *api.Client, update api.Update) {
	fmt.Println("Update with query:", update)
	handleMainMenuQueries(client, update.CallBackQuery)
	handleChooseGameQueries(client, update.CallBackQuery)
	handleMainGameQueries(client, update.CallBackQuery)
}