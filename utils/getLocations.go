package utils

import (
	"encoding/json"
	"strings"

	"github.com/QuarantineGameTeam/team2_qgame/models"
	"github.com/mitchellh/mapstructure"
)

//GetLocations returns array of pointers to models.Location from json string
func GetLocations(GotJSON string) ([]*models.Location, error) {
	result := []*models.Location{}
	var got []map[string]interface{}
	err := json.Unmarshal([]byte(GotJSON), &got)
	if err != nil {
		return result, err
	}
	for i := 0; i < len(got); i++ {
		var location models.Location
		switch {
		case strings.HasPrefix(got[i]["object_name"].(string), "Block"):
			location = &models.Block{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Cake Factory"):
			location = &models.CakeFactory{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Candy Factory"):
			location = &models.CandyFactory{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Chest"):
			location = &models.Chest{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Coffee Point"):
			location = &models.CoffeePoint{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Empty Field"):
			location = &models.EmptyField{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Monster"):
			location = &models.Monster{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Sign"):
			location = &models.Sign{}
			break
		case strings.HasPrefix(got[i]["object_name"].(string), "Sweet Home"):
			location = &models.SweetHome{}
			break
		}
		err = mapstructure.Decode(got[i], &location)
		if err != nil {
			return result, err
		}
		result = append(result, &location)
	}
	return result, err
}
