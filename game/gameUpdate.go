package game

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/handlers"
)

func GameUpdate(client *api.Client) {
	handlers.StartGames(client)
}