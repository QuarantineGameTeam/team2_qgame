package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/api"
)

var (
	startMarkup = api.InlineKeyboardMarkup{
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

	mainGameMarkup = api.InlineKeyboardMarkup{
		Buttons: [][]api.InlineKeyboardButton{
			// Row 1
			{
				{
					Text:     "Map",
					Callback: "map",
				},
				{
					Text:     "⬆️",
					Callback: "up",
				},
				{
					Text:     "Interact",
					Callback: "interact",
				},
			},
			// Row 2
			{
				{
					Text:     "⬅️",
					Callback: "left",
				},
				{
					Text:     "⬇️",
					Callback: "down",
				},
				{
					Text:     "➡️",
					Callback: "right",
				},
			},
			// Row 3
			{
				{
					Text:     "Castle",
					Callback: "castle",
				},
			},
		},
	}
)
