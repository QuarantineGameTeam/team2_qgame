package models

//EmptyField is a location, which does not contain any game objects
type EmptyField struct {
	X, Y int
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
