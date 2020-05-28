package models

type CoffeePoint struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"` //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"`
	StopMove      bool   `json:"stop_move"` //Changes the Player speed parameter.
	Message       string `json:"message"`   //
	X             int
	Y             int
	SmallPic      string `json:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"`
	BonusHealth   int    `json:"bonus_health"`
}

//GetLocation returns x and y
func (cp *CoffeePoint) GetLocation() (int, int) {
	return cp.X, cp.Y
}

//Interact just allows player to step onto this location
func (cp *CoffeePoint) Interact(player *Player) {
	if cp.OccupiedField == true {
		player.X = cp.X
		player.Y = cp.Y
	}
}
