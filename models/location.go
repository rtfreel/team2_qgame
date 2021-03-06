package models

//Location is a parent of every game object or matrix field
type Location interface {
	GetLocation() (int, int) //returns x and y locations
	Interact(player *Player) //location event on Player moves on it
}
