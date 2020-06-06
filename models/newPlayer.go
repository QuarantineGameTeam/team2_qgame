package models

import "github.com/QuarantineGameTeam/team2_qgame/api"

const (
	//Paths to pictures
	SmallPicPath = ""
	BigPicPath   = ""

	Active  = true
	Message = "You've met the player"

	//Start characteristics for each player
	StartHealth     = 100
	StartDexterity  = 100
	StartMastery    = 100
	StartDamage     = 100
	StartSpeed      = 1
	StartVisibility = 1
	StartCakes      = 10
	StartGold       = 10
	StartCandy      = 10
)

//NewPlayer returns pointer to the default Player
func NewPlayer(owner api.User, x, y int) *Player {
	player := &Player{
		X:          x,
		Y:          y,
		ObjectName: owner.Username,
		Message:    Message,
		PlayerId:   owner.ID,
		SmallPic:   SmallPicPath,
		BigPic:     BigPicPath,
		Active:     Active,
		Health:     StartHealth,
		Dexterity:  StartDexterity,
		Mastery:    StartMastery,
		Damage:     StartDamage,
		Speed:      StartSpeed,
		Visibility: StartVisibility,
		ScoreCake:  StartCakes,
		ScoreGold:  StartGold,
		ScoreCandy: StartCandy,
	}
	return player
}
