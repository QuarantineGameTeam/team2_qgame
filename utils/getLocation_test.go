package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/models"
)

func TestLocation_GetLocations(t *testing.T) {
	var user api.User = api.User{
		ID:       12345,
		Username: "Player1",
	}
	var player *models.Player = models.NewPlayer(user, 4, 4)
	var locations []models.Location = []models.Location{
		models.NewBlock(2, 4),
		models.NewCandyFactory(*player, 8, 1),
		models.NewSign(3, 2),
		models.NewSweetHome(*player, 6, 4),
		models.NewEmptyField(7, 3),
		models.NewCakeFactory(*player, 7, 6),
		models.NewMonster(8, 4),
		models.NewChest(3, 5),
		models.NewCoffeePoint(4, 2),
	}

	serialized, err := json.Marshal(locations)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	deserialized, err := GetLocations(string(serialized))
	if err != nil {
		t.Errorf("got error %v", err)
	}
	for i := 0; i < len(deserialized); i++ {
		wantType := reflect.TypeOf(locations[i])
		wantX, wantY := locations[i].GetLocation()
		gotType := reflect.TypeOf(deserialized[i])
		gotX, gotY := deserialized[i].GetLocation()
		testname := fmt.Sprintf("Compare actual Location (%s, %d, %d) to got Location (%s, %d, %d)", wantType, wantX, wantY,
			gotType, gotX, gotY)
		t.Run(testname, func(t *testing.T) {
			if gotType != wantType || gotX != wantX || gotY != wantY {
				t.Errorf("got Location(%s, %d, %d), want Location(%s, %d, %d)", gotType, gotX, gotY, wantType, wantX, wantY)
			}
		})
	}
}
