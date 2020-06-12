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
const Horizon = 1

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

	gm := new(Game)
	GenerateMap(gm)

	jsonBytes, err := json.Marshal(gm.Locations)
	gm.GameJSON = string(jsonBytes)
	gm.State = StateMatchMaking
	gm.PlayerID = starter.ID
	gm.Players = []models.Player{
		*models.NewPlayer(*starter, 0, 0),
	}
	bytePlayers, err := json.Marshal(gm.Players)
	if err != nil {
		log.Println(err)
	}
	gm.PlayersJSON = string(bytePlayers)

	if err != nil {
		log.Println("Error creating the gm.\n", err)
	}

	return gm, err
}