package game

import (
	"encoding/json"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"log"
)

const (
	StateMatchMaking = iota
	StateRunning     = iota
	StateEnded       = iota
)

const PlayersCount = 2

type Game struct {
	GameID int

	Locations []models.Location
	GameJSON  string

	PlayerID      int
	StartMoveTime int

	Players     []models.Player
	PlayersJSON string //for identifying a free place to add new gamer in current game

	State int //for game status
}

func NewGame(starter *api.User) (*Game, error) {
	var err error

	game := new(Game)
	game.Locations = generateLocations(starter)

	jsonBytes, err := json.Marshal(game.Locations)
	game.GameJSON = string(jsonBytes)
	game.State = StateMatchMaking
	game.Players = []models.Player{
		*models.NewPlayer(*starter, 0, 0),
	}
	bytePlayers, err := json.Marshal(game.Players)
	if err != nil {
		log.Println(err)
	}
	game.PlayersJSON = string(bytePlayers)

	if err != nil {
		log.Println("Error creating the game.\n", err)
	}

	return game, err
}

func generateLocations(starter *api.User) []models.Location {
	// needs to be relative with castle
	starterPlayer := models.NewPlayer(*starter, 0, 0)

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
		starterPlayer,
	}
}
