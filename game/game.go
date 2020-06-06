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
}