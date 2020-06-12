package game

import "github.com/QuarantineGameTeam/team2_qgame/game_model"

//LocatePlayers sets position for every player according to pre-generated spawn points
func LocatePlayers(game *game_model.Game) {
	for player := 0; player < len(game.Players); player++ {
		switch game.Players[player].Clan {
		case Clans[0]:
			game.Players[player].X = game.RedSpawn % Width
			game.Players[player].Y = game.RedSpawn / Width
			break
		case Clans[1]:
			game.Players[player].X = game.GreenSpawn % Width
			game.Players[player].Y = game.GreenSpawn / Width
			break
		case Clans[2]:
			game.Players[player].X = game.BlueSpawn % Width
			game.Players[player].Y = game.BlueSpawn / Width
			break
		}
	}
}
