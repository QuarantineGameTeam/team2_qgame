package models

import "github.com/QuarantineGameTeam/team2_qgame/api"

const (
	//Paths to pictures
	SmallPlayerPicPath = ""
	BigPlayerPicPath   = ""

	PlayerActive  = true
	PlayerMessage = "You've met the player"

	//Start characteristics for each player
	PlayerStartHealth     = 100
	PlayerStartDexterity  = 10
	PlayerStartMastery    = 10
	PlayerStartDamage     = 20
	PlayerStartSpeed      = 1
	PlayerStartVisibility = 1
	PlayerStartCakes      = 10
	PlayerStartGold       = 10
	PlayerStartCandy      = 10
)

//Player stores all information related to the ward of the user and its behaviour
type Player struct {
	X, Y       int
	ObjectName string `json:"object_name"`
	Message    string `json:"message"` //
	PlayerId   int    `json:"player_id"`
	SmallPic   string `json:"small_pic"` //path to pic
	BigPic     string `json:"big_pic"`
	Active     bool   `json:"active"` //if active = true, then drawing on the map and use the functional
	//
	Health     int `json:"health"`
	Dexterity  int `json:"dexterity"`
	Mastery    int `json:"mastery"`
	Damage     int `json:"damage"`
	Speed      int `json:"speed"`
	Visibility int `json:"visibility"`
	//score
	ScoreCake  int `json:"bonus_cake"`
	ScoreGold  int `json:"bonus_gold"`
	ScoreCandy int `json:"bonus_candy"`
}

//NewPlayer returns pointer to the default Player
func NewPlayer(owner api.User, x, y int) *Player {
	return &Player{
		Message:    PlayerMessage,
		Active:     PlayerActive,
		Health:     PlayerStartHealth,
		Dexterity:  PlayerStartDexterity,
		Mastery:    PlayerStartMastery,
		Damage:     PlayerStartDamage,
		Speed:      PlayerStartSpeed,
		Visibility: PlayerStartVisibility,
		ScoreCake:  PlayerStartCakes,
		ScoreGold:  PlayerStartGold,
		ScoreCandy: PlayerStartCandy,
		X:          x,
		Y:          y,
		ObjectName: owner.Username,
		PlayerId:   owner.ID,
		SmallPic:   SmallPlayerPicPath,
		BigPic:     BigPlayerPicPath,
	}
}

//GetLocation returns x and y
func (p *Player) GetLocation() (int, int) {
	return p.X, p.Y
}

func (p *Player) Interact(*Player){
	
}

//InteractWith makes player interact with game object on the location
func (p *Player) InteractWith(location *Location) {
	(*location).Interact(p)
}

//return SmallPic Path
func (p *Player) GetSmallPic() string {
	return p.SmallPic
}

//return BigPic Path
func (p *Player) GetBigPic() string {
	return p.BigPic
}
