package handlers

import (
	"fmt"
<<<<<<< HEAD
	"github.com/QuarantineGameTeam/team2_qgame/api"
=======
	"team2_qgame/api"
>>>>>>> f42b3f3afd86bec62aa8bc6a094df7066142ab24
)

func startMarkup(message api.UpdateMessage) api.InlineKeyboardMarkup{
	return api.InlineKeyboardMarkup{
		Buttons: [][]api.InlineKeyboardButton{
			{
				{
					Text:     "Join Game",
					Callback: fmt.Sprintf("joingame-%d", message.FromUser.ID),
				},
			},
			{
				{
					Text:     "View Stats",
					Callback: fmt.Sprintf("stats-%d", message.FromUser.ID),
				},
			},
			{
				{
					Text:     "Change Nickname",
					Callback: fmt.Sprintf("changenickname-%d", message.FromUser.ID),
				},
			},
		},
	}
}