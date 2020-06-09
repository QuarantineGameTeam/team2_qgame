package models

const (
	//Default parameters
	MonsterName          = "Monster"
	MonsterReusable      = false
	MonsterOccupiedField = false
	MonsterStopsMoving   = true
	MonsterMessage       = "There is a monster"
	MonsterActive        = true

	MonsterHealth     = 50
	MonsterDexterity  = 15
	MonsterMastery    = 15
	MonsterDamage     = 20
	MonsterVisibility = 1

	MonsterBonusDexterity = 5
	MonsterBonusMastery   = 5

	//Paths to pictures
	SmallMonsterPicPath = ""
	BigMonsterPicPath   = ""
)

type Monster struct {
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
	//
	Health     int `json:"health"`
	Dexterity  int `json:"dexterity"`
	Mastery    int `json:"mastery"`
	Damage     int `json:"speed"`
	Visibility int `json:"visibility"`
	//bonus
	BonusDexterity int `json:"bonus_dexterity" mapstructure:"bonus_dexterity"`
	BonusMastery   int `json:"bonus_mastery" mapstructure:"bonus_mastery"`
}

//NewMonster returns pointer to the default block
func NewMonster(x, y int) *Monster {
	return &Monster{
		ObjectName:     MonsterName,
		Repeatable:     MonsterReusable,
		OccupiedField:  MonsterOccupiedField,
		StopMove:       MonsterStopsMoving,
		Message:        MonsterMessage,
		Active:         MonsterActive,
		Health:         MonsterHealth,
		Dexterity:      MonsterDexterity,
		Mastery:        MonsterMastery,
		Damage:         MonsterDamage,
		Visibility:     MonsterVisibility,
		BonusDexterity: MonsterBonusDexterity,
		BonusMastery:   MonsterBonusMastery,
		X:              x,
		Y:              y,
		SmallPic:       SmallMonsterPicPath,
		BigPic:         BigMonsterPicPath,
	}
}

//GetLocation returns x and y
func (mn *Monster) GetLocation() (int, int) {
	return mn.X, mn.Y
}

//Interact just allows player to step onto this location
func (mn *Monster) Interact(player *Player) {
	if !mn.OccupiedField {
		player.X = mn.X
		player.Y = mn.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (mn *Monster) Update(player *Player) {
	player.Dexterity = player.Dexterity + mn.BonusDexterity
	player.Mastery = player.Mastery + mn.BonusMastery
}

//return SmallPic Path
func (mn *Monster) GetSmallPic() string {
	return mn.SmallPic
}

//return BigPic Path
func (mn *Monster) GetBigPic() string {
	return mn.BigPic
}
