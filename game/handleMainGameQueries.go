package game

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/models"
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

	players := game.Players
	if player != nil {
		ind := indexPlayer(players, *player)
		nextPlayer := models.Player{}
		if ind == len(players)-1 {
			nextPlayer = players[0]
		} else {
			nextPlayer = players[ind+1]
		}

		// sending update to the user, who did the move
		SendCurrentPhoto(client, query.FromUser)
		// sending move buttons to next player
		SendCurrentPhoto(client, api.User{ID: nextPlayer.PlayerId})
		SendMoveButtons(client, api.User{ID: nextPlayer.PlayerId})
	} else {
		fmt.Println("Player is <nil>")
	}

	err := client.DeleteMessage(query.Message)
	if err != nil {
		log.Println("Unable to delete message: ", err)
	}
}
