package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/config"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"log"
)

func handleUpdateMessage(client *api.Client, update api.Update) {
	message := update.Message
	if message.Text == "/start" {
		handleStartMessage(client, message)
	} else if message.FromUser.State == config.StateChangingName {
		handleChangeNickNameMessage(client, message)
	}
}

func handleStartMessage(client *api.Client, message api.UpdateMessage) {
	dbh := database.NewDBHandler()

	var err error
	if dbh.ContainsUser(message.FromUser) {
		err = client.SendMessage(api.Message{
			ChatID:       message.FromUser.ID,
			Text:         fmt.Sprintf("Hello, %s! Welcome back!", message.FromUser.FirstName),
			InlineMarkup: startMarkup(message),
		})
	} else {
		err = client.SendMessage(api.Message{
			ChatID:       message.FromUser.ID,
			Text:         fmt.Sprintf("Hello, %s! Welcome to CandyWarGO!", message.FromUser.FirstName),
			InlineMarkup: startMarkup(message),
		})

		// adding user to database if it is not there
		dbh.InsertUser(message.FromUser)
	}

	if err != nil {
		log.Println(err)
	}
}

func handleChangeNickNameMessage(client *api.Client, message api.UpdateMessage) {
	dbh := database.NewDBHandler()

	dbh.Update("users", "nickname", message.Text, "telegram_id", message.FromUser.ID)
	dbh.Update("users", "state", config.StateNone, "telegram_id", message.FromUser.ID)
}
