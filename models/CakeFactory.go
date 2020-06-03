package models

type CakeFactory struct {
	ObjectName    string `json:"object_name"`
	Repeatable    bool   `json:"constantly"`     //player can use repeatedly
	OccupiedField bool   `json:"occupied_field"` //If true -player can occupie this field
	StopMove      bool   `json:"stop_move"`      //Changes the Player speed parameter.
	Message       string `json:"message"`        //
	Active        bool   `json:"active"`         //if active = true, then drawing on the map and use the functional
	X             int
	Y             int
	Owner         int    `json:"owner_id"`  //player Id
	SmallPic      string `json:"small_pic"` //path to pic
	BigPic        string `json:"big_pic"`
	//
	Friendly   bool `json:"friendly"` //if true, the player and ally can visit. If false must fight
	Health     int  `json:"health"`
	Dexterity  int  `json:"dexterity"`
	Mastery    int  `json:"mastery"`
	Damage     int  `json:"speed"`
	Visibility int  `json:"visibility"`
	//bonus
	BonusCake int `json:"bonus_cake"`
	BonusGold int `json:"bonus_gold"`
}

//GetLocation returns x and y
func (cf *CakeFactory) GetLocation() (int, int) {
	return cf.X, cf.Y
}

//Interact just allows player to step onto this location
func (cf *CakeFactory) Interact(player *Player) {
	if cf.OccupiedField == true {
		player.X = cf.X
		player.Y = cf.Y
	}
}

//score update (adds some resources and bonuses from game objects)
func (cf *CakeFactory) Update(player *Player) {
	player.ScoreCake = player.ScoreCake + cf.BonusCake
	player.ScoreGold = player.ScoreGold + cf.BonusGold
}

//return SmallPic Path
func (cf *CakeFactory) GetSmallPic() string {
	return cf.SmallPic
}

//return BigPic Path
func (cf *CakeFactory) GetBigPic() string {
	return cf.BigPic
}