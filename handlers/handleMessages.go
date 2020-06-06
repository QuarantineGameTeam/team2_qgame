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
	} else if fitsState(message.FromUser, config.StateChangingName) {
		handleChangeNickNameMessage(client, message)
	}
}

func handleStartMessage(client *api.Client, message api.UpdateMessage) {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	contains, err := dbh.ContainsUser(message.FromUser)
	if err != nil {
		log.Println(err)
	}

	if contains {
		err = client.SendMessage(api.Message{
			ChatID:       message.FromUser.ID,
			Text:         fmt.Sprintf("Hello, %s! Welcome back!", message.FromUser.FirstName),
			InlineMarkup: startMarkup,
		})
	} else {
		err = client.SendMessage(api.Message{
			ChatID:       message.FromUser.ID,
			Text:         fmt.Sprintf("Hello, %s! Welcome to CandyWarGO!", message.FromUser.FirstName),
			InlineMarkup: startMarkup,
		})

		// adding user to database if it is not there
		err = dbh.InsertUser(message.FromUser)
		if err != nil {
			log.Println(err)
		}
	}

	if err != nil {
		log.Println(err)
	}
}

func handleChangeNickNameMessage(client *api.Client, message api.UpdateMessage) {
	success := true
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
		success = false
	}

	err = dbh.Update("users", "nickname", message.Text, "telegram_id", message.FromUser.ID)
	if err != nil {
		log.Println(err)
		success = false
	}

	err = dbh.Update("users", "state", config.StateNone, "telegram_id", message.FromUser.ID)
	if err != nil {
		log.Println(err)
		success = false
	}

	var msg string
	if success {
		msg = "Nickname changed successfully."
	} else {
		msg = "Sorry. Some error happened."
	}

	err = client.SendMessage(
		api.Message{
			ChatID: message.FromUser.ID,
			Text:   msg,
			InlineMarkup: startMarkup,
		})
}
