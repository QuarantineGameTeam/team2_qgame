package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
)

var (
	startMarkup =  api.InlineKeyboardMarkup{
		Buttons: [][]api.InlineKeyboardButton{
			{
				{
					Text:     "Join Game",
					Callback: "joingame",
				},
			},
			{
				{
					Text:     "View Stats",
					Callback: "stats",
				},
			},
			{
				{
					Text:     "Change Nickname",
					Callback: "changenickname",
				},
			},
		},
	}

	chooseClanMarkup = api.InlineKeyboardMarkup{
		Buttons: [][]api.InlineKeyboardButton{
			{
				{
					Text:     "Red",
					Callback: "red",
				},
			},
			{
				{
					Text:     "Green",
					Callback: "green",
				},
			},
			{
				{
					Text:     "Blue",
					Callback: "blue",
				},
			},
		},
	}

)