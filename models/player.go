package models

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

//GetLocation returns x and y
func (p *Player) GetLocation() (int, int) {
	return p.X, p.Y
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
