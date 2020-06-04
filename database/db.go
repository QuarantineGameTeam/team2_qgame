package database

import (
	"database/sql"
	"fmt"
	"github.com/QuarantineGameTeam/team2_qgame/api"
	"log"

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
func (dbh *DBHandler) Connect() error {
	var err error
	dbh.Connection, err = sql.Open(dbh.DriverName, dbh.DBPath)

	return err
}

//CreateTables is a shortcut to create all necessary tables
func (dbh *DBHandler) CreateTables() {
	dbh.CreateUsersTable()
}

//CreateUsersTable creates a table of Users with two fields, if one does not already exist
func (dbh *DBHandler) CreateUsersTable() error {
	_, err := dbh.Connection.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    		   		telegram_id INTEGER UNIQUE PRIMARY KEY,
					nickname TEXT,
					state INTEGER);`)
	return err
}

//InsertUser adds a user to the Users table
func (dbh *DBHandler) InsertUser(user api.User) error {
	//User structure is described in the api package file user.go
	_, err := dbh.Connection.Exec(`INSERT INTO users (telegram_id, nickname, state) VALUES (?, ?, ?);`, user.ID, user.Username, user.State)

	return err
}

//Update updates any field in any table with new value
func (dbh *DBHandler) Update(table, field string, value interface{}, whereField string, whereValue interface{}) error {
	statement := fmt.Sprintf(`UPDATE %s SET %s = ? WHERE %s = ?;`, table, field, whereField)
	_, err := dbh.Connection.Exec(statement, value, whereValue)

	return err
}

// GetField returns value of field in given table in respect to to some parameter
func (dbh *DBHandler) GetField(table, field, whereField string, whereVal interface {}) interface{} {
	result, err := dbh.Connection.Query(fmt.Sprintf(`SELECT %s FROM %s WHERE %s = ?;`, field, table, whereField), whereVal)

	if err != nil {
		log.Println(err)
	}
	defer result.Close()

	var state int
	if result.Next() {
		err = result.Scan(&state)
		if err != nil {
			log.Println(err)
			return 0
		}
	}

	return state
}


//NameExists returns true if a user with the same name is already registered
func (dbh *DBHandler) NameExists(name string) (bool, error) {
	result, err := dbh.Connection.Query(`SELECT * FROM users WHERE nickname = ?;`, name)

	if result != nil {
		defer result.Close()
		if result.Next() {
			return true, err
		}
	}
	return false, err
}

//ContainsUser returns true if a user with the same name is already registered
func (dbh *DBHandler) ContainsUser(user api.User) (bool, error) {
	result, err := dbh.Connection.Query(`SELECT * FROM users WHERE telegram_id = ?;`, user.ID)

	if result != nil {
		defer result.Close()
		if result.Next() {
			return true, err
		}
	}
	return false, err
}

//GetUserByID returns api.User object from database with specified id
func (dbh *DBHandler) GetUserByID(id int) (*api.User, error) {
	var user *api.User = &api.User{}
	result, err := dbh.Connection.Query(`SELECT * FROM users WHERE telegram_id = ?;`, id)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	if result.Next() {
		err := result.Scan(&user.ID, &user.Username, &user.State)
		return user, err
	}
	return user, err
}
