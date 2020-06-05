package models

type CoffeePoint struct {
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
	BonusHealth   int    `json:"bonus_health"` //maximum allowed value = Health of Player at start
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

//score update (adds some resources and bonuses from game objects)
func (cp *CoffeePoint) Update(player *Player) {
	player.Health = cp.BonusHealth //maximum allowed value = Health of Player at start
}

//return SmallPic Path
func (cp *CoffeePoint) GetSmallPic() string {
	return cp.SmallPic
}

//return BigPic Path
func (cp *CoffeePoint) GetBigPic() string {
	return cp.BigPic
}
