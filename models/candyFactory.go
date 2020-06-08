package models

const (
	//Default parameters
	CandyFactoryName          = "Cake Factory"
	CandyFactoryReusable      = true
	CandyFactoryOccupiedField = true
	CandyFactoryStopsMoving   = true
	CandyFactoryMessage       = "There is a candy factory"
	CandyFactoryActive        = true

	CandyFactoryFriendly   = false
	CandyFactoryHealth     = 50
	CandyFactoryDexterity  = 5
	CandyFactoryMastery    = 5
	CandyFactoryDamage     = 10
	CandyFactoryVisibility = 1

	CandyFactoryBonusCandy = 5
	CandyFactoryBonusGold  = 2

	//Paths to pictures
	SmallCandyFactoryPicPath = ""
	BigCandyFactoryPicPath   = ""
)

type CandyFactory struct {
	ObjectName    string `json:"object_name" mapstructure:"object_name"`
	Repeatable    bool   `json:"constantly" mapstructure:"constantly"`     //player can use repeatedly
	OccupiedField bool   `json:"occupied_field" mapstructure:"occupied_field"` //If true -player can occupie this field
	StopMove      bool   `json:"stop_move" mapstructure:"stop_move"`      //Changes the Player speed parameter.
	Message       string `json:"message"`        //
	Active        bool   `json:"active"`         //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	Owner         int    `json:"owner_id" mapstructure:"owner_id"`  //player Id
	SmallPic      string `json:"small_pic" mapstructure:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"  mapstructure:"big_pic"`
	//
	Friendly   bool `json:"friendly"` //if true, the player and ally can visit. If false must fight
	Health     int  `json:"health"`
	Dexterity  int  `json:"dexterity"`
	Mastery    int  `json:"mastery"`
	Damage     int  `json:"speed"`
	Visibility int  `json:"visibility"`
	//bonus
	BonusCandy int `json:"bonus_candy" mapstructure:"bonus_candy"`
	BonusGold  int `json:"bonus_gold" mapstructure:"bonus_gold"`
}

//NewCakeFactory returns pointer to the default block
func NewCandyFactory(owner Player, x, y int) *CakeFactory {
	return &CakeFactory{
		ObjectName:    CandyFactoryName,
		Repeatable:    CandyFactoryReusable,
		OccupiedField: CandyFactoryOccupiedField,
		StopMove:      CandyFactoryStopsMoving,
		Message:       CandyFactoryMessage,
		Active:        CandyFactoryActive,
		Friendly:      CandyFactoryFriendly,
		Health:        CandyFactoryHealth,
		Dexterity:     CandyFactoryDexterity,
		Mastery:       CandyFactoryMastery,
		Damage:        CandyFactoryDamage,
		Visibility:    CandyFactoryVisibility,
		BonusCake:     CandyFactoryBonusCandy,
		BonusGold:     CandyFactoryBonusGold,
		X:             x,
		Y:             y,
		Owner:         owner.PlayerId,
		SmallPic:      SmallCandyFactoryPicPath,
		BigPic:        BigCandyFactoryPicPath,
	}
}

//GetLocation returns x and y
func (cf *CandyFactory) GetLocation() (int, int) {
	return cf.X, cf.Y
}

//Interact just allows player to step onto this location
func (cf *CandyFactory) Interact(player *Player) {
	if !cf.OccupiedField {
		player.X = cf.X
		player.Y = cf.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (cf *CandyFactory) Update(player *Player) {
	player.ScoreCandy = player.ScoreCandy + cf.BonusCandy
	player.ScoreGold = player.ScoreGold + cf.BonusGold
}

//return SmallPic Path
func (cf *CandyFactory) GetSmallPic() string {
	return cf.SmallPic
}

//return BigPic Path
func (cf *CandyFactory) GetBigPic() string {
	return cf.BigPic
}
