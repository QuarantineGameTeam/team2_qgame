package game

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/drawers"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"log"
	"strconv"
)

func handleMainGameQueries(client *api.Client, query api.CallBackQuery) {
	switch query.CallBackData {
	case "map":
		handleMapQueries(client, query)
	case "interact":
		// todo
	case "castle":
		// todo
	case "up", "down", "left", "right":
		handleControlsQueries(client, query)
	}
}

func handleMapQueries(client *api.Client, query api.CallBackQuery) {
	// Gonna be working with MapView as soon as visited points are implemented

	_, err := client.SendMessage(api.Message {
		ChatID: query.FromUser.ID,
		Text: "This is all you've visited so far!",
	})
	if err != nil {
		log.Println(err)
	}

	gm := getPlayerGame(query.FromUser)

	photoLocation := fmt.Sprintf("temp/%d.png", query.FromUser.ID)
	err = drawers.CreateFullViewPhoto(gm.Locations, gm.Players, strconv.Itoa(query.FromUser.ID))
	if err != nil {
		log.Println(err)
	}

	err = client.SendPhoto(query.FromUser.ID, photoLocation)
	if err != nil {
		log.Println(err)
	}
}

func handleControlsQueries(client *api.Client, query api.CallBackQuery) {
	player := getPlayer(query.FromUser)
	game := getPlayerGame(query.FromUser)

	var x, y int // coords of an object to interact with

	if player != nil && game != nil {
		switch query.CallBackData {
		case "up":
			x = player.X
			y = player.Y - 1
		case "down":
			x = player.X
			y = player.Y + 1
		case "left":
			x = player.X - 1
			y = player.Y
		case "right":
			x = player.X + 1
			y = player.Y
		}

		for _, l := range game.Locations {
			haveX, haveY := l.GetLocation()
			if haveX == x && haveY == y {
				player.InteractWith(&l)
				break
			}
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
