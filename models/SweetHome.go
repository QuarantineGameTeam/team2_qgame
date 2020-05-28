package models

type SweetHome struct {
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
	//bonus every day
	BonusCake  int `json:"bonus_cake"`
	BonusCandy int `json:"bonus_candy"`
	BonusGold  int `json:"bonus_gold"`
	//bonus on visit
	BonusHealth int `json:"bonus_health"`
	//Bonus can be bought
	BonusDexterity  int `json:"bonus_dexterity"`
	BonusMastery    int `json:"bonus_mastery"`
	BonusVisibility int `json:"bonus_visibility"`
	BonusSpeed      int `json:"bonus_speed"`
}

//GetLocation returns x and y
func (sh *SweetHome) GetLocation() (int, int) {
	return sh.X, sh.Y
}

//Interact just allows player to step onto this location
func (sh *SweetHome) Interact(player *Player) {
	if sh.OccupiedField == true {
		player.X = sh.X
		player.Y = sh.Y
	}
}
