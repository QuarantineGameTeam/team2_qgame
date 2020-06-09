package models

//Location is a parent of every game object or matrix field
type Location interface {
	GetLocation() (int, int) //returns x and y locations
	Interact(player *Player) //location event on Player moves on it
	GetSmallPic() string     //return SmallPic Path
	GetBigPic() string       //return BigPic Path
	Update(player *Player)   //score update (adds some resources from factories and other bonuses)
}
