package models

const (
	//Default parameters
	SweetHomeName          = "Sweet Home"
	SweetHomeReusable      = true
	SweetHomeOccupiedField = true
	SweetHomeStopsMoving   = true
	SweetHomeMessage       = "There is a sweet home"
	SweetHomeActive        = true

	SweetHomeFriendly   = true
	SweetHomeHealth     = 120
	SweetHomeDexterity  = 50
	SweetHomeMastery    = 50
	SweetHomeDamage     = 40
	SweetHomeVisibility = 1

	SweetHomeBonusCandy      = 5
	SweetHomeBonusCake       = 2
	SweetHomeBonusGold       = 5
	SweetHomeBonusHealth     = 2
	SweetHomeBonusDexterity  = 5
	SweetHomeBonusMastery    = 2
	SweetHomeBonusVisibility = 5
	SweetHomeBonusSpeed      = 2

	//Paths to pictures
	SmallSweetHomePicPath = "photos/castle-red.png"
	BigSweetHomePicPath   = ""
)

type SweetHome struct {
	ObjectName    string `json:"object_name" mapstructure:"object_name"`
	Repeatable    bool   `json:"constantly" mapstructure:"constantly"`         //player can use repeatedly
	OccupiedField bool   `json:"occupied_field" mapstructure:"occupied_field"` //If true -player can occupie this field
	StopMove      bool   `json:"stop_move" mapstructure:"stop_move"`           //Changes the Player speed parameter.
	Message       string `json:"message"`                                      //
	Active        bool   `json:"active"`                                       //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	Owner         int    `json:"owner_id" mapstructure:"owner_id"`   //player Id
	SmallPic      string `json:"small_pic" mapstructure:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"  mapstructure:"big_pic"`
	//
	Friendly   bool `json:"friendly"` //if true, the player and ally can visit. If false must fight
	Health     int  `json:"health"`
	Dexterity  int  `json:"dexterity"`
	Mastery    int  `json:"mastery"`
	Damage     int  `json:"speed"`
	Visibility int  `json:"visibility"`
	//bonus every day
	BonusCake  int `json:"bonus_cake" mapstructure:"bonus_cake"`
	BonusCandy int `json:"bonus_candy" mapstructure:"bonus_candy"`
	BonusGold  int `json:"bonus_gold" mapstructure:"bonus_gold"`
	//bonus on visit
	BonusHealth int `json:"bonus_health" mapstructure:"bonus_health"`
	//Bonus can be bought
	BonusDexterity  int `json:"bonus_dexterity" mapstructure:"bonus_dexterity"`
	BonusMastery    int `json:"bonus_mastery" mapstructure:"bonus_mastery"`
	BonusVisibility int `json:"bonus_visibility" mapstructure:"bonus_visibility"`
	BonusSpeed      int `json:"bonus_speed" mapstructure:"bonus_speed"`
}

//NewSweetHome returns pointer to the default block
func NewSweetHome(owner Player, x, y int) *SweetHome {
	return &SweetHome{
		ObjectName:      SweetHomeName,
		Repeatable:      SweetHomeReusable,
		OccupiedField:   SweetHomeOccupiedField,
		StopMove:        SweetHomeStopsMoving,
		Message:         SweetHomeMessage,
		Active:          SweetHomeActive,
		Friendly:        SweetHomeFriendly,
		Health:          SweetHomeHealth,
		Dexterity:       SweetHomeDexterity,
		Mastery:         SweetHomeMastery,
		Damage:          SweetHomeDamage,
		Visibility:      SweetHomeVisibility,
		X:               x,
		Y:               y,
		Owner:           owner.PlayerId,
		SmallPic:        SmallSweetHomePicPath,
		BigPic:          BigSweetHomePicPath,
		BonusCake:       SweetHomeBonusCake,
		BonusCandy:      SweetHomeBonusCandy,
		BonusGold:       SweetHomeBonusGold,
		BonusHealth:     SweetHomeBonusHealth,
		BonusDexterity:  SweetHomeDexterity,
		BonusMastery:    SweetHomeBonusMastery,
		BonusVisibility: SweetHomeBonusVisibility,
		BonusSpeed:      SweetHomeBonusSpeed,
	}
}

//GetLocation returns x and y
func (sh *SweetHome) GetLocation() (int, int) {
	return sh.X, sh.Y
}

//Interact just allows player to step onto this location
func (sh *SweetHome) Interact(player *Player) {
	if !sh.OccupiedField {
		player.X = sh.X
		player.Y = sh.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (sh *SweetHome) Update(player *Player) {
	player.ScoreCake = player.ScoreCake + sh.BonusCake
	player.ScoreGold = player.ScoreGold + sh.BonusGold
	player.ScoreCandy = player.ScoreCandy + sh.BonusCandy
	player.Health = sh.BonusHealth //maximum allowed value = Health of Player at start
}

//return SmallPic Path
func (sh *SweetHome) GetSmallPic() string {
	return sh.SmallPic
}

//return BigPic Path
func (sh *SweetHome) GetBigPic() string {
	return sh.BigPic
}
