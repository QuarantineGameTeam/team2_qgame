package models

const (
	//Default parameters
	ChestName          = "Chest"
	ChestReusable      = false
	ChestOccupiedField = false
	ChestStopsMoving   = true
	ChestMessage       = "There is a chest"
	ChestActive        = true
	ChestAddMastery    = 10

	//Paths to pictures
	SmallChestPicPath = ""
	BigChestPicPath   = ""
)

type Chest struct {
	ObjectName      string `json:"object_name" mapstructure:"object_name"`
	Repeatable      bool   `json:"constantly" mapstructure:"constantly"`         //player can use repeatedly
	OccupiedField   bool   `json:"occupied_field" mapstructure:"occupied_field"` //If true -player can occupie this field
	StopMove        bool   `json:"stop_move" mapstructure:"stop_move"`           //Changes the Player speed parameter.
	Message         string `json:"message"`                                      //
	Active          bool   `json:"active"`                                       //if active = true, then drawing on the map and use the functional
	X               int
	Y               int
	SmallPic        string `json:"small_pic" mapstructure:"small_pic"` //path to pic
	BigPic          string `json:"big_pic"  mapstructure:"big_pic"`
	AddMasteryLevel int    `json:"add_mastery_level"  mapstructure:"add_mastery_level"`
}

//NewChest returns pointer to the default block
func NewChest(x, y int) *Chest {
	return &Chest{
		ObjectName:      ChestName,
		Repeatable:      ChestReusable,
		OccupiedField:   ChestOccupiedField,
		StopMove:        ChestStopsMoving,
		Message:         ChestMessage,
		Active:          ChestActive,
		AddMasteryLevel: ChestAddMastery,
		X:               x,
		Y:               y,
		SmallPic:        SmallChestPicPath,
		BigPic:          BigChestPicPath,
	}
}

//GetLocation returns x and y
func (ch *Chest) GetLocation() (int, int) {
	return ch.X, ch.Y
}

//Interact just allows player to step onto this location
func (ch *Chest) Interact(player *Player) {
	if !ch.OccupiedField {
		player.X = ch.X
		player.Y = ch.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (ch *Chest) Update(player *Player) {
	player.Mastery = player.Mastery + ch.AddMasteryLevel
}

//return SmallPic Path
func (ch *Chest) GetSmallPic() string {
	return ch.SmallPic
}

//return BigPic Path
func (ch *Chest) GetBigPic() string {
	return ch.BigPic
}
