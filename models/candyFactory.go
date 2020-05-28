package models

type CandyFactory struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"` //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"`
	StopMove      bool   `json:"stop_move"` //Changes the Player speed parameter.
	Message       string `json:"message"`   //
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
	Damage     int  `json:"speed"`
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
