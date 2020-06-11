package game

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/game_model"

	"github.com/QuarantineGameTeam/team2_qgame/models"
)

func TestMap_GenerateMap(t *testing.T) {
	user := api.User{ID: 12345, Username: "Player1"}
	game, err := game_model.NewGame(&user)
	if err != nil {
		t.Errorf("got error %v", err)
	}

	columns := Width / ColumnWidth

	testname := fmt.Sprintf("Generate new map")
	t.Run(testname, func(t *testing.T) {
		GenerateMap(game)

		EF := 0
		KF := 0
		NF := 0
		CH := 0
		SI := 0
		MO := 0
		SH := 0
		CP := 0
		BL := 0
		OB := 0

		for i := 0; i < len(game.Locations); i++ {
			switch reflect.TypeOf(game.Locations[i]) {
			case reflect.TypeOf(&models.EmptyField{}):
				EF++
				break
			case reflect.TypeOf(&models.CakeFactory{}):
				KF++
				break
			case reflect.TypeOf(&models.CandyFactory{}):
				NF++
				break
			case reflect.TypeOf(&models.Chest{}):
				CH++
				break
			case reflect.TypeOf(&models.Sign{}):
				SI++
				break
			case reflect.TypeOf(&models.Monster{}):
				MO++
				break
			case reflect.TypeOf(&models.SweetHome{}):
				x, y := game.Locations[i].GetLocation()
				fmt.Printf("Home x=%d, y=%d \n", x, y)
				SH++
				break
			case reflect.TypeOf(&models.CoffeePoint{}):
				CP++
				break
			case reflect.TypeOf(&models.Block{}):
				BL++
				break
			default:
				OB++
				break
			}
		}

		game.Players = []models.Player{
			{Clan: "red"},
			{Clan: "green"},
			{Clan: "blue"},
		}

		LocatePlayers(game)
		for player := 0; player < len(game.Players); player++ {
			fmt.Printf("%s x=%d, y=%d \n", game.Players[player].Clan, game.Players[player].X, game.Players[player].Y)
		}

		if len(game.Locations) != Width*Height || KF != CakeFactories*columns || NF != CandyFactories*columns || CH != Chests*columns ||
			SI != Signs*columns || MO != Monsters*columns || SH != SweetHomes*columns ||
			CP != CoffeePoints*columns || BL != Blocks*columns || OB > 0 ||
			EF != len(game.Locations)-KF-NF-CH-SI-MO-SH-CP-BL {
			t.Errorf("amount of objects is incorrect")
		}
	})
}
