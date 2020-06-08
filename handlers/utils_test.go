package handlers

import (
	"github.com/QuarantineGameTeam/team2_qgame/database"
	"github.com/QuarantineGameTeam/team2_qgame/game"
	"github.com/QuarantineGameTeam/team2_qgame/models"
	"reflect"
	"testing"
)

func Test_updateGameAfterMove(t *testing.T) {
	type have struct {
		game   *game.Game
		player *models.Player
	}
	type want struct {
		game *game.Game
	}
	tests := []struct {
		name string
		have have
		want want
	}{
		{
			"Ok test",
			have{
				&game.Game{
					GameID: 0,
					Locations: []models.Location{
						&models.CandyFactory{
							X: 7, Y: 4,
						},
						&models.Chest{
							X: 4, Y: 6,
						},
						&models.Player{
							PlayerId: 1234,
							X:        3, Y: 4,
						},
					},
					Players: []models.Player{
						models.Player{
							PlayerId: 1234,
							X:        3, Y: 4,
						},
					},
				},
				&models.Player{
					PlayerId: 1234,
					X:        3, Y: 4,
				},
			},
			want{
				&game.Game{
					GameID: 0,
					Locations: []models.Location{
						&models.CandyFactory{
							X: 7, Y: 4,
						},
						&models.Chest{
							X: 4, Y: 6,
						},
						&models.Player{
							PlayerId: 1234,
							X:        3, Y: 5,
						},
					},
					Players: []models.Player{
						models.Player{
							PlayerId: 1234,
							X:        3, Y: 5,
						},
					},
				},
			},
		},
	}

	dbh, err := database.NewDBHandlerWithPath("CandyWarGoDatabaseTest.sqlite")
	if err != nil {
		t.Errorf("Got err %v", err)
	}

	for _, tt := range tests {
		err = dbh.InsertGame(*tt.have.game)
		if err != nil {
			t.Errorf("Got err %v", err)
		}

		tt.have.player.Y++
		updateGameAfterMove(tt.have.game, tt.have.player)

		got, err := dbh.GetGameByID(tt.have.game.GameID)
		if err != nil {
			t.Errorf("Got err %v", err)
		}

		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("Test failed. Got %v. wanted %v", got, tt.want.game)
		}
	}
}
