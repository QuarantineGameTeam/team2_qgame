package models

type Monster struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"`     //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"` //If true -player can occupie this field
	StopMove      bool   `json:"stop_move"`      //Changes the Player speed parameter.
	Message       string `json:"message"`        //
	Active        bool   `json:"active"`         //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	SmallPic      string `json:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"`
	//
	Health     int `json:"health"`
	Dexterity  int `json:"dexterity"`
	Mastery    int `json:"mastery"`
	Damage     int `json:"speed"`
	Visibility int `json:"visibility"`
	//bonus
	BonusDexterity int `json:"bonus_dexterity"`
	BonusMastery   int `json:"bonus_mastery"`
}

//GetLocation returns x and y
func (mn *Monster) GetLocation() (int, int) {
	return mn.X, mn.Y
}

//Interact just allows player to step onto this location
func (mn *Monster) Interact(player *Player) {
	if mn.OccupiedField == true {
		player.X = mn.X
		player.Y = mn.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (mn *Monster) Update(player *Player) {
	player.Dexterity = player.Dexterity + mn.BonusDexterity
	player.Mastery = player.Mastery + mn.BonusMastery
}

//return SmallPic Path
func (mn *Monster) GetSmallPic() string {
	return mn.SmallPic
}

//return BigPic Path
func (mn *Monster) GetBigPic() string {
	return mn.BigPic
}
