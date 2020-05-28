package models

type Monster struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"` //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"`
	StopMove      bool   `json:"stop_move"` //Changes the Player speed parameter.
	Message       string `json:"message"`   //
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
