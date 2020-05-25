package models

//Player stores all information related to the ward of the user and its behaviour
type Player struct {
	X, Y int
}

//InteractWith makes player interact with game object on the location
func (p *Player) InteractWith(location *Location) {
	(*location).Interact(p)
}
