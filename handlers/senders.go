package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/game"
	"log"
)

func sendChooseClanMarkup(client *api.Client, gm *game.Game) {
	for _, p := range gm.Players {
		_, err := client.SendMessage(api.Message {
			ChatID: p.PlayerId,
			Text: `Game is Starting. You have 1 minute to choose a clan.
			Otherwise, you will join a random clan.`,
			InlineMarkup: chooseClanMarkup,
		})
		if err != nil {
			log.Println(err)
		}
	}
}