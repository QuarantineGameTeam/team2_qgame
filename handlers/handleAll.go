package handlers

import (
	"fmt"
	"log"
	"team2_qgame/api"
)

func HandleUpdate(client *api.Client, update api.Update) {
	if update.HasMessage() {
		handleUpdateMessage(client, update)
	}
}

func handleUpdateMessage(client *api.Client, update api.Update) {
	message := update.Message
	if message.Text == "/start" {
		err := client.SendMessage(api.Message{
			ChatID: message.FromUser.ID,
			Text:   fmt.Sprintf("Hello, %s!", message.FromUser.FirstName),
		})
		if err != nil{
			log.Println(err)
		}
	}
}
