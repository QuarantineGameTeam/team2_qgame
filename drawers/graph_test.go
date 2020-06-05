package drawing

import (
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"testing"
)

func Test_scale(t *testing.T) {
	tests := []struct {
		name   string
		v     float64
		min1   float64
		min2   float64
		max1   float64
		max2   float64
		want float64
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

func Test_CreatePartViewPhoto(t *testing.T) {
	locations := []models.Location {
		&models.CandyFactory{
			ObjectName: "cf1",
			SmallPic: "photos/candy_factory.png",
			X: 3, Y: 4,
		},
		&models.Chest {
			ObjectName: "ch1",
			SmallPic: "photos/chest.png",
			X: 4, Y: 6,
		},
	}

	err := CreatePartViewPhoto(locations, 4, 4, 1, "test_part_view")
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

