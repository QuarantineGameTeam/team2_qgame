package drawers

import (
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"math/rand"
	"testing"
)

func Test_scale(t *testing.T) {
	tests := []struct {
		name    string
		v       float64
		min1    float64
		min2    float64
		max1    float64
		max2    float64
		want    float64
		wantErr bool
	}{
		{
			"positive", 2, 0, 0, 10, 100, 20, false,
		},
		{
			"negative", 3, 0, 0, 10, 100, 40, true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v2 := scale(test.v, test.min1, test.max1, test.min2, test.max2)
			if (v2 == test.want) == test.wantErr {
				t.Errorf("got value 2 %v, want value 2 %v", v2, test.want)
			}
		})
	}
}

func Test_drawBackground(t *testing.T) { // just as useless as aqua from konosuba(who watched knows what i mean :) )
	tests := []struct {
		name       string
		x, y, w, h int
	}{
		{
			"uselessness lvl: aqua", 0, 0, 1023, 1024,
		},
	}

	var wantimage bool
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.w != 1024 {
				wantimage = false
			} else if test.h != 1024 {
				wantimage = false
			} else {
				wantimage = true
			}
			if wantimage != true {
				t.Errorf("wanted image got successful true, got successful %t", wantimage)
			}
		})
	}

}

func Test_drawGrid(t *testing.T) {
	tests := []struct {
		name      string
		dimension int
	}{
		{
			"pointless but ok", -3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.dimension < 0 {
				t.Errorf("got value %d negative, want %d positive", test.dimension, test.dimension)
			}
		})
	}
}

func Test_CreatePartViewPhoto(t *testing.T) {
	locations := []models.Location{
		&models.CandyFactory{
			ObjectName: "cf1",
			SmallPic:   "photos/candy_factory.png",
			X:          3, Y: 4,
		},
		&models.Chest{
			ObjectName: "ch1",
			SmallPic:   "photos/chest.png",
			X:          4, Y: 6,
		},
	}

	p1 := *models.NewPlayer(api.User{}, 1, 2)
	p1.SmallPic = "photos/enemy.png"
	p2 := *models.NewPlayer(api.User{}, 3, 7)
	p2.SmallPic = "photos/enemy.png"

	players := []models.Player{
		p1, p2,
	}

	err := CreatePartViewPhoto(locations, players,4, 4, 1, "test_part_view")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func Test_CreateMapViewPhoto(t *testing.T) {
	locations := []models.Location{
		&models.CandyFactory{
			ObjectName: "cf1",
			SmallPic:   "photos/candy_factory.png",
			X:          3,
			Y:          4,
		},
		&models.Chest{
			ObjectName: "ch1",
			SmallPic:   "photos/chest.png",
			X:          4,
			Y:          6,
		},
	}

	p1 := *models.NewPlayer(api.User{}, 1, 2)
	p1.SmallPic = "photos/enemy.png"
	p2 := *models.NewPlayer(api.User{}, 3, 7)
	p2.SmallPic = "photos/enemy.png"

	players := []models.Player{
		p1, p2,
	}


	visited := make([][]bool, defaultDimension)
	for i := 0; i < defaultDimension; i++ {
		visited[i] = make([]bool, defaultDimension)
	}

	for i := 0; i < defaultDimension; i++ {
		for j := 0; j < defaultDimension; j++ {
			r := rand.Intn(100)
			if r > 50 {
				visited[i][j] = true
			} else {
				visited[i][j] = false
			}
		}
	}

	err := CreateMapViewPhoto(locations, players, visited, "test_map_view")

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func Test_CreateFullViewPhoto(t *testing.T) {
	locations := []models.Location{
		&models.CandyFactory{
			ObjectName: "cf1",
			SmallPic:   "photos/candy_factory.png",
			X:          3, Y: 4,
		},
		&models.Chest{
			ObjectName: "ch1",
			SmallPic:   "photos/chest.png",
			X:          4, Y: 6,
		},
	}

	p1 := *models.NewPlayer(api.User{}, 1, 2)
	p1.SmallPic = "photos/enemy.png"
	p2 := *models.NewPlayer(api.User{}, 3, 7)
	p2.SmallPic = "photos/enemy.png"

	players := []models.Player{
		p1, p2,
	}

	err := CreateFullViewPhoto(locations, players, "test_full_view")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}
