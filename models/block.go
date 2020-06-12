package models

const (
	//Default parameters
	BlockName          = "Block"
	BlockReusable      = true
	BlockOccupiedField = true
	BlockStopsMoving   = true
	BlockMessage       = "There is a block"
	BlockActive        = true

	//Paths to pictures
	SmallBlockPicPath = "photos/block.png"
	BigBlockPicPath   = ""
)

type Block struct {
	ObjectName    string `json:"object_name" mapstructure:"object_name"`
	Repeatable    bool   `json:"constantly" mapstructure:"constantly"` //player can use repeatedly
	OccupiedField bool   `json:"occupied_field" mapstructure:"occupied_field"`
	StopMove      bool   `json:"stop_move" mapstructure:"stop_move"` //Changes the Player speed parameter.
	Message       string `json:"message"`   //You tried to overcome this obstacle, but the forces left you and you remained in the same place
	Active        bool   `json:"active"`    //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	SmallPic      string `json:"small_pic" mapstructure:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"  mapstructure:"big_pic"`
}

//NewBlock returns pointer to the default block
func NewBlock(x, y int) *Block {
	return &Block{
		ObjectName:    BlockName,
		Repeatable:    BlockReusable,
		OccupiedField: BlockOccupiedField,
		StopMove:      BlockStopsMoving,
		Message:       BlockMessage,
		Active:        BlockActive,
		X:             x,
		Y:             y,
		SmallPic:      SmallBlockPicPath,
		BigPic:        BigBlockPicPath,
	}
}

//GetLocation returns x and y
func (bl *Block) GetLocation() (int, int) {
	return bl.X, bl.Y
}

//Interact just allows player to step onto this location
func (bl *Block) Interact(player *Player) {
	if !bl.OccupiedField {
		player.X = bl.X
		player.Y = bl.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (bl *Block) Update(player *Player) {
}

//return SmallPic Path
func (bl *Block) GetSmallPic() string {
	return bl.SmallPic
}

//return BigPic Path
func (bl *Block) GetBigPic() string {
	return bl.BigPic
}
