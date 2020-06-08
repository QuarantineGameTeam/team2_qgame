package models

const (
	//Default parameters
	EmptyName          = "Empty Field"
	EmptyReusable      = true
	EmptyOccupiedField = false

	//Paths to pictures
	SmallEmptyPicPath = ""
	BigEmptyPicPath   = ""
)

//EmptyField is a location, which does not contain any game objects
type EmptyField struct {
	X, Y          int
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"`     //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"` //If true -player can occupie this field
	SmallPic      string `json:"small_pic"`      //path to pic
	BigPic        string `json:"big_pic"`
}

//NewEmptyField returns pointer to the default block
func NewEmptyField(x, y int) *EmptyField {
	return &EmptyField{
		ObjectName:    EmptyName,
		Repeatable:    EmptyReusable,
		OccupiedField: EmptyOccupiedField,
		X:             x,
		Y:             y,
		SmallPic:      SmallEmptyPicPath,
		BigPic:        BigEmptyPicPath,
	}
}

//GetLocation returns x and y
func (ef *EmptyField) GetLocation() (int, int) {
	return ef.X, ef.Y
}

//Interact just allows player to step onto this location
func (ef *EmptyField) Interact(player *Player) {
	player.X = ef.X
	player.Y = ef.Y
}

//score update (adds some resources from factories and other bonuses)
func (ef *EmptyField) Update(player *Player) {

}

//return SmallPic Path
func (ef *EmptyField) GetSmallPic() string {
	return ef.SmallPic
}

//return BigPic Path
func (ef *EmptyField) GetBigPic() string {
	return ef.BigPic
}
