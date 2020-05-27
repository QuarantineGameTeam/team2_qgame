package models

import (
	"math/rand"
	"reflect"
	"time"
)

//Matrix is a playground
type Matrix struct {
	Width, Height, BlockWidth int
	Locations                 [][]Location
}

//GetLocation returns object from location on the map
func (m *Matrix) GetLocation(x, y int) Location {
	return m.Locations[x][y]
}

//SetLocation places object on specified location on the map
func (m *Matrix) SetLocation(x, y int, location *Location) {
	m.Locations[x][y] = *location
}

//GenerateMap fills matrix with empty fields and game objects
func (m *Matrix) GenerateMap() {
	m.Empty()
	//TODO: call m.SetObjects for every object type we need
}

//Empty fills matrix with empty fields
func (m *Matrix) Empty() {
	m.Resize(m.Width, m.Height)
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			m.Locations[x][y] = &EmptyField{x, y}
		}
	}
}

//Resize creates empty array of Locations in matrix with specified width and height
func (m *Matrix) Resize(width, height int) {
	m.Width = width
	m.Height = height
	m.Locations = make([][]Location, m.Height)
	for i := 0; i < m.Height; i++ {
		m.Locations[i] = make([]Location, m.Width)
	}
}

//SetObjects sets specified amount of objects to every block on the map
func (m *Matrix) SetObjects(object string, amount int) {
	seed := rand.NewSource(time.Now().UnixNano()) //setting time as a seed for random
	random := rand.New(seed)                      //setting seed
	var x, y int
	for block := 0; block < m.Width/m.BlockWidth; block++ { //block is a wide column, which should contain certain amount of objects
		for i := 0; i < amount; i++ {
			for fieldIsEmpty := true; fieldIsEmpty; fieldIsEmpty = !m.IsEmpty(x, y) { //do-while loop implementation
				x = random.Intn(m.BlockWidth) + (block * m.BlockWidth) //getting random value within current block
				y = random.Intn(m.Height)                              //getting random value within map height
			}
			switch object {
			case "castle":
				//TODO: set Factory object to x, y location
				//m.SetLocation(x, y, GetCastle(x, y))
			case "factory":
				//TODO: set Factory object to x, y location
				//m.SetLocation(x, y, GetFactory(x, y))
			case "chest":
				//TODO: set Factory object to x, y location
				//m.SetLocation(x, y, GetChest(x, y))
			case "nameplate":
				//TODO: set Factory object to x, y location
				//m.SetLocation(x, y, GetNameplate(x, y))
			case "trap":
				//TODO: set Factory object to x, y location
				//m.SetLocation(x, y, GetTrap(x, y))
			case "enemy":
				//TODO: set Factory object to x, y location
				//m.SetLocation(x, y, GetEnemy(x, y))
			}
		}
	}
}

//IsEmpty returns true if location is EmptyField
func (m *Matrix) IsEmpty(x, y int) bool {
	if reflect.TypeOf(m.GetLocation(x, y)) == reflect.TypeOf(&EmptyField{}) {
		return true
	}
	return false
}
