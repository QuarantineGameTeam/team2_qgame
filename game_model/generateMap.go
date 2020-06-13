package game_model

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/QuarantineGameTeam/team2_qgame/models"
)

const (
	//Map parameters
	Width       = 9
	Height      = 9
	ColumnWidth = 3
	//Objects amount on the one map column
	SweetHomes     = 1
	Blocks         = 1
	CakeFactories  = 1
	CandyFactories = 1
	Chests         = 1
	CoffeePoints   = 1
	Signs          = 1
	Monsters       = 3
)

//Clans contains clan names
var Clans []string = []string{"red", "green", "blue"}

//SweetHomesSet contains amount of SweetHomes, which have been already set
var SweetHomesSet int = 0

//GenerateMap returns array of models.Location with objects needed
func GenerateMap(game *Game) {
	game.Locations = getEmptyMap()

	setObjects(&game.Locations, "CakeFactory", CakeFactories)
	setObjects(&game.Locations, "CandyFactory", CandyFactories)
	setObjects(&game.Locations, "Chest", Chests)
	setObjects(&game.Locations, "CoffeePoint", CoffeePoints)
	setObjects(&game.Locations, "Sign", Signs)
	setObjects(&game.Locations, "Block", Blocks)
	setObjects(&game.Locations, "Monster", Monsters)
	for column := 0; column < Width/ColumnWidth; column++ {
		createSweetHome(game, column)
	}
}

//getEmptyMap returns array of EmptyFields
func getEmptyMap() []models.Location {
	result := make([]models.Location, Width * Height)
	for w := 0; w < Width; w++ {
		for h := 0; h < Height; h++ {
			result[h * Width + w] = models.NewEmptyField(w, h)
		}
	}
	return result
}

//SetObjects sets specified amount of objects to every block on the map
func setObjects(matrix *[]models.Location, object string, amount int) {
	seed := rand.NewSource(time.Now().UnixNano()) //setting time as a seed for random
	random := rand.New(seed)                      //setting seed
	var x, y int
	for col := 0; col < Width/ColumnWidth; col++ { //block is a wide column, which should contain certain amount of objects
		for i := 0; i < amount; i++ {
			for fieldIsEmpty := false; !fieldIsEmpty; fieldIsEmpty = isEmpty(matrix, x, y) { //do-while loop implementation
				x = random.Intn(ColumnWidth) + (col * ColumnWidth) //getting random value within current block
				y = random.Intn(Height)                            //getting random value within map height
			}
			switch object {
			case "CakeFactory":
				(*matrix)[y*Width+x] = models.NewCakeFactory(x, y)
			case "CandyFactory":
				(*matrix)[y*Width+x] = models.NewCandyFactory(x, y)
			case "CoffeePoint":
				(*matrix)[y*Width+x] = models.NewCoffeePoint(x, y)
			case "Chest":
				(*matrix)[y*Width+x] = models.NewChest(x, y)
			case "Sign":
				(*matrix)[y*Width+x] = models.NewSign(x, y)
			case "Block":
				(*matrix)[y*Width+x] = models.NewBlock(x, y)
			case "Monster":
				(*matrix)[y*Width+x] = models.NewMonster(x, y)
			}
		}
	}
}

func isEmpty(matrix *[]models.Location, x, y int) bool {
	if reflect.TypeOf((*matrix)[y*Width+x]) == reflect.TypeOf(&models.EmptyField{}) {
		return true
	}
	return false
}

//Point stores coordinates data
type Point struct {
	X, Y int
}

func (point *Point) isPoint() bool {
	if point.X < 0 || point.X >= Width ||
		point.Y < 0 || point.Y >= Height {
		return false
	}
	return true
}

func addPoints(point1, point2 Point) Point {
	return Point{point1.X + point2.X, point1.Y + point2.Y}
}

func createSweetHome(game *Game, col int) {
	matrix := &game.Locations
	seed := rand.NewSource(time.Now().UnixNano()) //setting time as a seed for random
	random := rand.New(seed)                      //setting seed

	var homeSpawnPoints []Point
	var playerSpawnPoints []Point

	//Fill spawnPoints with coordinates, which has two emptyField in a row to go from
	var delta []Point = []Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} //adding this points returns point rightside, leftside, upside and downside
	for y := 0; y < Height; y++ {                                 //brute force every location in the column
		for x := ColumnWidth * col; x < ColumnWidth*(col+1); x++ {
			curPoint := Point{x, y}
			if isEmpty(matrix, curPoint.X, curPoint.Y) { //check if this location is empty
				for i := 0; i < len(delta); i++ { //brute force every neighbour points from current
					nextPoint := addPoints(curPoint, delta[i])
					if nextPoint.isPoint() && isEmpty(matrix, nextPoint.X, nextPoint.Y) { //validate point and check if it is empty
						for j := 0; j < len(delta); j++ {
							lastPoint := addPoints(nextPoint, delta[j]) //brute force every neighbour points from previous
							if lastPoint.isPoint() && !reflect.DeepEqual(lastPoint, curPoint) &&
								isEmpty(matrix, lastPoint.X, lastPoint.Y) { //validate and check if point is empty
								homeSpawnPoints = append(homeSpawnPoints, curPoint)      //add point to spawn sweet homes
								playerSpawnPoints = append(playerSpawnPoints, nextPoint) //add point to spawn player
							}
						}
					} else {
						continue
					}
				}
			} else {
				continue
			}
		}
	}

	//Select current clan
	curClan := Clans[SweetHomesSet]
	SweetHomesSet++
	//Select one of possible points to spawn SweetHome & Player
	pointIndex := random.Intn(len(homeSpawnPoints))
	//Create sweet home
	sweetHomePoint := homeSpawnPoints[pointIndex]
	homeName := fmt.Sprintf("Sweet Home-%s", curClan)
	picPath := fmt.Sprintf("photos/castle-%s.png", curClan)

	(*matrix)[sweetHomePoint.Y*Width+sweetHomePoint.X] = &models.SweetHome{
		ObjectName: homeName,
		X:          sweetHomePoint.X,
		Y:          sweetHomePoint.Y,
		SmallPic:   picPath,
	}
	//Create playerSpawnPoint
	var playerPoint int = playerSpawnPoints[pointIndex].Y*Width + playerSpawnPoints[pointIndex].X
	switch curClan {
	case Clans[0]:
		game.RedSpawn = playerPoint
		break
	case Clans[1]:
		game.GreenSpawn = playerPoint
		break
	case Clans[2]:
		game.BlueSpawn = playerPoint
		break
	}
}
