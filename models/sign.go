package models

const (
	//Default parameters
	SignName          = "Sign"
	SignReusable      = false
	SignOccupiedField = false
	SignStopsMoving   = true
	SignMessage       = "There is a sign"
	SignActive        = true

	//Paths to pictures
	SmallSignPicPath = ""
	BigSignPicPath   = ""
)

type Sign struct {
	ObjectName    string `json:"object_name" mapstructure:"object_name"`
	Repeatable    bool   `json:"constantly" mapstructure:"constantly"`     //player can use repeatedly
	OccupiedField bool   `json:"occupied_field" mapstructure:"occupied_field"` //If true -player can occupie this field
	StopMove      bool   `json:"stop_move" mapstructure:"stop_move"`      //Changes the Player speed parameter.
	Message       string `json:"message"`        //
	Active        bool   `json:"active"`         //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	SmallPic      string `json:"small_pic" mapstructure:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"  mapstructure:"big_pic"`
}

//NewSign returns pointer to the default block
func NewSign(x, y int) *Sign {
	return &Sign{
		ObjectName:    SignName,
		Repeatable:    SignReusable,
		OccupiedField: SignOccupiedField,
		StopMove:      SignStopsMoving,
		Message:       SignMessage,
		Active:        SignActive,
		X:             x,
		Y:             y,
		SmallPic:      SmallSignPicPath,
		BigPic:        BigSignPicPath,
	}
}

//GetLocation returns x and y
func (sn *Sign) GetLocation() (int, int) {
	return sn.X, sn.Y
}

//Interact just allows player to step onto this location
func (sn *Sign) Interact(player *Player) {
	if !sn.OccupiedField {
		player.X = sn.X
		player.Y = sn.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (sn *Sign) Update(player *Player) {

}

//return SmallPic Path
func (sn *Sign) GetSmallPic() string {
	return sn.SmallPic
}

//return BigPic Path
func (sn *Sign) GetBigPic() string {
	return sn.BigPic
}
