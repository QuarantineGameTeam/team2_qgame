package handlers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"log"
)

func handleMainGameQueries(client *api.Client, query api.CallBackQuery) {
	switch query.CallBackData {
	case "map":

	case "interact":

	case "castle":

	case "up", "down", "left", "right":
		handleControlsQueries(client, query)
	}
}

func handleControlsQueries(client *api.Client, query api.CallBackQuery) {
	player := getPlayer(query.FromUser)
	game := getPlayerGame(query.FromUser)

	if player != nil && game != nil {
		switch query.CallBackData {
		case "up":
			player.Y--
		case "down":
			player.Y++
		case "left":
			player.X--
		case "right":
			player.X++
		}
	} else {
		if player == nil {
			fmt.Printf("Player %v does not exist", player)
		}
		if game == nil {
			fmt.Printf("Player %v is not in any game", player)
		}
	}

	updateGameAfterMove(game, player)

	// ------------------ //
	// TEST PURPOSES ONLY //
	// ------------------ //

	photoLocation := "temp/testpic.png"
	err := drawers.CreateFullViewPhoto(game.Locations, game.Players, "testpic")
	if err != nil {
		log.Println(err)
	}

	err = client.SendPhoto(query.FromUser.ID, photoLocation)
	if err != nil {
		log.Println(err)
	}

	// WEIRD ASYNC RIGHT NOW
	// Sending move buttons one by one for players
	_, err = client.SendMessage(api.Message{
		ChatID:       query.FromUser.ID,
		Text:         "Your turn.",
		InlineMarkup: mainGameMarkup,
	})
	if err != nil {
		log.Println(err)
	}

	// ----------------- //

	err = client.DeleteMessage(query.Message)
	if err != nil {
		log.Println("Unable to delete message: ", err)
	}
}
