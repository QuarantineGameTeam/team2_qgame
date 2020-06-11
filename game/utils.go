package game

import (
	"encoding/json"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/game_model"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"log"
)

func fitsState(user api.User, state int) bool {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	res := dbh.GetField("users", "state", "telegram_id", user.ID)

	return res == state
}

func getPlayer(user api.User) *models.Player {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	player, err := dbh.GetPlayerByID(user.ID)
	if err != nil {
		return nil
	}

	return player
}

func getPlayerGame(user api.User) *game_model.Game {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	games := dbh.GetGames()
	for _, gm := range games {
		for _, p := range gm.Players {
			if p.PlayerId == user.ID {
				return gm
			}
		}
	}

	return nil
}

func updateGameAfterMove(game *game_model.Game, player *models.Player) {
	players := game.Players
	for i, p := range players {
		if p.PlayerId == player.PlayerId {
			game.Players[i] = *player
		}
	}

	bytes, err := json.Marshal(players)
	if err != nil {
		log.Println(err)
	}
	game.PlayersJSON = string(bytes)

	updateDBGame(game)
	updateDBPlayer(player)
}

func updateDBPlayer(player *models.Player) {
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	err = dbh.Update("players", "x", player.X, "player_id", player.PlayerId)
	if err != nil {
		log.Println(err)
	}
	err = dbh.Update("players", "y", player.Y, "player_id", player.PlayerId)
	if err != nil {
		log.Println(err)
	}
}

func updateDBGame(game *game_model.Game) {
	// writing changes to database
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}
	err = dbh.Update("games", "game_json", game.GameJSON, "game_id", game.GameID)
	if err != nil {
		log.Println(err)
	}
	err = dbh.Update("games", "players_json", game.PlayersJSON, "game_id", game.GameID)
	if err != nil {
		log.Println(err)
	}
}

func indexPlayer(players []models.Player, toFind models.Player) int {
	for i, v := range players {
		if v.PlayerId == toFind.PlayerId {
			return i
		}
	}
	return -1
}