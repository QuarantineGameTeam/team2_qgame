package models

type Block struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"` //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"`
	StopMove      bool   `json:"stop_move"` //Changes the Player speed parameter.
	Message       string `json:"message"`   //You tried to overcome this obstacle, but the forces left you and you remained in the same place
	Active        bool   `json:"active"`    //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	SmallPic      string `json:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"`
}

//GetLocation returns x and y
func (bl *Block) GetLocation() (int, int) {
	return bl.X, bl.Y
}

//Interact just allows player to step onto this location
func (bl *Block) Interact(player *Player) {
	if bl.OccupiedField == true {
		player.X = bl.X
		player.Y = bl.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (bl *Block) Update(player *Player) {

}
