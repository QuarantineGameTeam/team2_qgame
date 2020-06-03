package database

import (
	"database/sql"
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"

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

//CreateTables is a shortcut to create all necessary tables
func (dbh *DBHandler) CreateTables() {
	dbh.CreateUsersTable()
}

//CreateUsersTable creates a table of Users with two fields, if one does not already exist
func (dbh *DBHandler) CreateUsersTable() {
	_, err := dbh.Connection.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    		   		telegram_id INTEGER UNIQUE PRIMARY KEY,
					nickname TEXT,
					state INTEGER);`)
	if err != nil {
		panic(err)
	}
}

//InsertUser adds a user to the Users table
func (dbh *DBHandler) InsertUser(user api.User) {
	//User structure is described in the api package file user.go
	_, err := dbh.Connection.Exec(`INSERT INTO users (telegram_id, nickname, state) VALUES (?, ?, ?);`, user.ID, user.Username, user.State)
	if err != nil {
		panic(err)
	}
}

//Update updates any field in any table with new value
func (dbh *DBHandler) Update(table, field string, value interface{}, whereField string, whereValue interface{}) {
	statement := fmt.Sprintf(`UPDATE %s SET %s = ? WHERE %s = ?;`, table, field, whereField)
	_, err := dbh.Connection.Exec(statement, value, whereValue)
	if err != nil {
		panic(err)
	}
}

//NameExists returns true if a user with the same name is already registered
func (dbh *DBHandler) NameExists(name string) bool {
	result, err := dbh.Connection.Query(`SELECT * FROM users WHERE nickname = ?;`, name)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	if result.Next() {
		return true
	}
	return false
}

//ContainsUser returns true if a user with the same name is already registered
func (dbh *DBHandler) ContainsUser(user api.User) bool {
	result, err := dbh.Connection.Query(`SELECT * FROM users WHERE telegram_id = ?;`, user.ID)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	if result.Next() {
		return true
	}
	return false
}

//GetUserByID returns api.User object from database with specified id
func (dbh *DBHandler) GetUserByID(id int) *api.User {
	var user *api.User = &api.User{}
	result, err := dbh.Connection.Query(`SELECT * FROM users WHERE telegram_id = ?;`, id)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	if result.Next() {
		err := result.Scan(&user.ID, &user.Username, &user.State)
		if err != nil {
			fmt.Println(err)
		}
	}
	return user
}
