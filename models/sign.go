package models

type Sign struct {
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
}

//GetLocation returns x and y
func (sn *Sign) GetLocation() (int, int) {
	return sn.X, sn.Y
}

//Interact just allows player to step onto this location
func (sn *Sign) Interact(player *Player) {
	if sn.OccupiedField == true {
		player.X = sn.X
		player.Y = sn.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (sn *Sign) Update(player *Player) {

}
