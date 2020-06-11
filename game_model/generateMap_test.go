package game_model

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/QuarantineGameTeam/team2_qgame/models"
)

func TestMap_GenerateMap(t *testing.T) {

	columns := Width / ColumnWidth

	testname := fmt.Sprintf("Generate new map")
	t.Run(testname, func(t *testing.T) {
		test := GenerateMap()

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

		for i := 0; i < len(test); i++ {
			switch reflect.TypeOf(test[i]) {
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
		if len(test) != Width*Height || KF != CakeFactories*columns || NF != CandyFactories*columns || CH != Chests*columns ||
			SI != Signs*columns || MO != Monsters*columns || SH != SweetHomes*columns ||
			CP != CoffeePoints*columns || BL != Blocks*columns || OB > 0 ||
			EF != len(test)-KF-NF-CH-SI-MO-SH-CP-BL {
			t.Errorf("amount of objects is incorrect")
		}
	})
}
