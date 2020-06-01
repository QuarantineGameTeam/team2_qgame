package handlers

import (
	"fmt"
	"log"
	"team2_qgame/api"
	"team2_qgame/database"
)

func handleUpdateMessage(client *api.Client, update api.Update) {
	message := update.Message
	if message.Text == "/start" {
		handleStartMessage(client, message)
	}
}

func handleStartMessage(client *api.Client, message api.UpdateMessage) {
	// Setting up database handler
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
