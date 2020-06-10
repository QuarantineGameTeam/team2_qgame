package game

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/QuarantineGameTeam/team2_qgame/models"
)

var MO = 0
var KF = 0
var NF = 0
var CH = 0
var SI = 0
var SH = 0
var BL = 0
var CP = 0

func TestMap_GenerateMap(t *testing.T) {
	//this test prints generated map by text
	test := GenerateMap()
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			fmt.Print(ObjectString(reflect.TypeOf(test[j*9+i])))
			if i != 8 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func ObjectString(typ reflect.Type) string {
	switch typ {
	case reflect.TypeOf(&models.EmptyField{}):
		return "EF"
	case reflect.TypeOf(&models.CakeFactory{}):
		return "KF"
	case reflect.TypeOf(&models.CandyFactory{}):
		return "NF"
	case reflect.TypeOf(&models.Chest{}):
		return "CH"
	case reflect.TypeOf(&models.Sign{}):
		return "SI"
	case reflect.TypeOf(&models.Monster{}):
		return "MO"
	case reflect.TypeOf(&models.SweetHome{}):
		return "SH"
	case reflect.TypeOf(&models.CoffeePoint{}):
		return "CP"
	case reflect.TypeOf(&models.Block{}):
		return "BL"
	default:
		return "OB"
	}
}
