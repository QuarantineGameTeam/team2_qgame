package models

type CandyFactory struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"`     //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"` //If true -player can occupie this field
	StopMove      bool   `json:"stop_move"`      //Changes the Player speed parameter.
	Message       string `json:"message"`        //
	Active        bool   `json:"active"`         //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	Owner         int    `json:"owner_id"`  //player Id
	SmallPic      string `json:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"`
	//
	Friendly   bool `json:"friendly"` //if true, the player and ally can visit. If false must fight
	Health     int  `json:"health"`
	Dexterity  int  `json:"dexterity"`
	Mastery    int  `json:"mastery"`
	Damage     int  `json:"damage"`
	Visibility int  `json:"visibility"`
	//bonus
	BonusCandy int `json:"bonus_candy"`
	BonusGold  int `json:"bonus_gold"`
}

//GetLocation returns x and y
func (cf *CandyFactory) GetLocation() (int, int) {
	return cf.X, cf.Y
}

//Interact just allows player to step onto this location
func (cf *CandyFactory) Interact(player *Player) {
	if cf.OccupiedField == true {
		player.X = cf.X
		player.Y = cf.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (cf *CandyFactory) Update(player *Player) {
	player.ScoreCandy = player.ScoreCandy + cf.BonusCandy
	player.ScoreGold = player.ScoreGold + cf.BonusGold
}

//return SmallPic Path
func (cf *CandyFactory) GetSmallPic() string {
	return cf.SmallPic
}

//return BigPic Path
func (cf *CandyFactory) GetBigPic() string {
	return cf.BigPic
}
