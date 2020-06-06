package game

import (
	"github.com/QuarantineGameTeam/team2_qgame/models"
)

type Game struct {
	GameID        int
	Locations     []models.Location
	PlayerID      int
	StartMoveTime int
	GameJSON      string
	Players       string //for identifying a free place to add new gamer in current game
	State         int    //for game status
}
