package database

import (
	"database/sql"
	"team2_qgame/api"

	//import sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

//DBHandler holds interfaces to interact with database
type DBHandler struct {
	DriverName string
	DBPath     string
	Connection *sql.DB
}

//Connect creates connection with database file by driver
func (dbh *DBHandler) Connect() {
	var err error
	dbh.Connection, err = sql.Open(dbh.DriverName, dbh.DBPath)
	if err != nil {
		panic(err)
	}
}

/*
type User struct {
	TelegramID sql.NullInt //it is Telegram user ID
	NickName   sql.NullString //nickname in the game
}
*/

//CreateUsersTable creates a table of Users with two fields, if one does not already exist
func (dbh *DBHandler) CreateUsersTable() {
	_, err := dbh.Connection.Exec(`CREATE TABLE IF NOT EXISTS Users (
		TelegramID INTEGER PRIMARY KEY, 
		NickName TEXT);`)
	if err != nil {
		panic(err)
	}
}

//InsertUser adds a user to the Users table
func (dbh *DBHandler) InsertUser(user api.User) {
	//User structure is described in the api package file user.go
	_, err := dbh.Connection.Exec("INSERT INTO Users (TelegramID, NickName) VALUES (?, ?);", user.ID, user.Username) //name is NickName
	if err != nil {
		panic(err)
	}
}

//NameExists returns true if a user with the same name is already registered
func (dbh *DBHandler) NameExists(name string) bool {
	result, err := dbh.Connection.Query("SELECT * FROM Users WHERE NickName = ?;", name)
	if err != nil {
		panic(err)
	}
	if result != nil {
		return true
	}
	return false
}
