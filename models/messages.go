package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var MsgData Messages

type Messages struct {
	SweetHomeMsg    []Msg `json:"sweet_home_msg"`
	CofeePointMsg   []Msg `json:"coffee_point_msg"`
	BlockMsg        []Msg `json:"block_msg"`
	PlayerMsg       []Msg `json:"player_msg"`
	CakeFactoryMsg  []Msg `json:"cake_factory_msg"`
	CandyFactoryMsg []Msg `json:"candy_factory_msg"`
	ChestMsg        []Msg `json:"chest_msg"`
	MonsterMsg      []Msg `json:"monster_msg"`
	SignMsg         []Msg `json:"sign_msg"`
}

type Msg struct {
	Ru string `json:"ru"`
	Ua string `json:"ua"`
	En string `json:"en"`
}

func GetMessages() {
	jsonFile, err := os.Open("messages.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &MsgData)
}
