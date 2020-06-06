package game

import (
	"encoding/json"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"log"
)

type Game struct {
	GameID        int
	Locations     []models.Location
	PlayerID      int
	StartMoveTime int
	GameJSON      string
}

func NewGame() (*Game, error) {
	var err error

	game := new(Game)
	game.Locations = generateLocations()
	jsonBytes, err := json.Marshal(game.Locations)
	game.GameJSON = string(jsonBytes)
	if err != nil {
		log.Println("Error creating the game.\n", err)
	}

	return game, err
}

func generateLocations() []models.Location {
	// has to perform generating. Right now its giving out some of actual nothing ...
	return []models.Location{
		&models.CandyFactory{
			ObjectName: "cf1",
			SmallPic:   "photos/candy_factory.png",
			X:          3, Y: 4,
		},
		&models.Chest{
			ObjectName: "ch1",
			SmallPic:   "photos/chest.png",
			X:          4, Y: 6,
		},
		&models.Player{
			ObjectName: "p1",
			X:          5, Y: 5,
			SmallPic: "photos/enemy.png",
		},
	}
}
