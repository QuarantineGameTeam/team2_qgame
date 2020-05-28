package models

type Chest struct {
	ObjectName      string `json:"object_name"`
	Repeatable      bool   `json:"constantly"` //player can use repeatedly
	OccupiedField   bool   `json:"occupied_field"`
	StopMove        bool   `json:"stop_move"` //Changes the Player speed parameter.
	Message         string `json:"message"`   //
	X               int
	Y               int
	SmallPic        string `json:"small_pic"` //path to pic
	BigPic          string `json:"big_pic"`
	AddMasteryLevel int    `json:"add_mastery_level"`
}

//GetLocation returns x and y
func (ch *Chest) GetLocation() (int, int) {
	return ch.X, ch.Y
}

//Interact just allows player to step onto this location
func (ch *Chest) Interact(player *Player) {
	if ch.OccupiedField == true {
		player.X = ch.X
		player.Y = ch.Y
	}
}
