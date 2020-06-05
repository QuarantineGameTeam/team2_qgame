package handlers

import (
	"fmt"
	"team2_qgame/api"
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