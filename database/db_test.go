package database

import (
	"fmt"
	"team2_qgame/api"
	"testing"
)

func TestDBHandler(t *testing.T) {
	var tests = []api.User{
		api.User{
			ID:       12345,
			Username: "player1",
			State:    0,
		},
		api.User{
			ID:       54321,
			Username: "player2",
			State:    1,
		},
		api.User{
			ID:       56789,
			Username: "player3",
			State:    2,
		},
		api.User{
			ID:       98765,
			Username: "player4",
			State:    3,
		},
	}

	db := NewDBHandler()
	//Insert test
	for _, tt := range tests {
		testname := fmt.Sprintf("Insert user (%d, %s, %d) into database, if there is not one", tt.ID, tt.Username, tt.State)
		t.Run(testname, func(t *testing.T) {
			if !db.ContainsUser(tt) {
				db.InsertUser(tt)
			}
			user := db.GetUserByID(tt.ID)
			if user.Username != tt.Username || user.State != tt.State {
				t.Errorf("got User (%d, %s, %d), want User (%d, %s, %d)", user.ID, user.Username, user.State, tt.ID, tt.Username, tt.State)
			}
		})
	}
	//User is registered test
	for _, tt := range tests {
		testname := fmt.Sprintf("Check, if user (%d, %s) has been already inserted into database", tt.ID, tt.Username)
		t.Run(testname, func(t *testing.T) {
			flag := db.NameExists(tt.Username)
			if !flag {
				t.Errorf("got %v, want %v", flag, true)
			}
		})
	}
	//Update test
	for i, tt := range tests {
		newName := fmt.Sprintf("player%d", i)
		testname := fmt.Sprintf("Update user (%d, %s) to user (%d, %s)", tt.ID, tt.Username, tt.ID, newName)
		t.Run(testname, func(t *testing.T) {
			db.Update("users", "nickname", newName, "telegram_id", tt.ID)
			user := db.GetUserByID(tt.ID)
			if user.Username != newName {
				t.Errorf("got User (%d, %s, %d), want User (%d, %s, %d)", user.ID, user.Username, user.State, tt.ID, newName, tt.State)
			}
		})
	}
}
