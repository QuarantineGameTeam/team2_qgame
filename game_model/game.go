package game_model

import (
	"encoding/json"
	"log"

	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/models"
)

const (
	StateMatchMaking = iota
	StateRunning     = iota
	StateEnded       = iota
)

const PlayersCount = 3

type Game struct {
	GameID int

	Locations []models.Location
	GameJSON  string

	PlayerID      int
	StartMoveTime int

	Players     []models.Player
	PlayersJSON string //for identifying a free place to add new gamer in current game

	RedSpawn   int
	GreenSpawn int
	BlueSpawn  int

	State int //for game status
}

func NewGame(starter *api.User) (*Game, error) {
	var err error

	game := new(Game)
	game.Locations = generateLocations(starter)

	jsonBytes, err := json.Marshal(game.Locations)
	game.GameJSON = string(jsonBytes)
	game.State = StateMatchMaking
	game.PlayerID = starter.ID
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
	return []models.Location{
		models.NewCandyFactory(3, 4),
		models.NewChest(3, 6),
	}
}
