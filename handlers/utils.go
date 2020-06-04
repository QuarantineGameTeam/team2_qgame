package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"log"
)

func fitsState(user api.User, state int) bool{
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	res := dbh.GetField("users", "state", "telegram_id", user.ID)

	return res == state
}
