package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/QuarantineGameTeam/team2_qgame/api"
	"github.com/QuarantineGameTeam/team2_qgame/models"

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
func (dbh *DBHandler) CreateTables() error {
	err := dbh.CreateUsersTable()
	if err != nil {
		return err
	}
	err = dbh.CreatePlayersTable()
	return err
}

//CreateUsersTable creates a table of Users if one does not already exist
func (dbh *DBHandler) CreateUsersTable() error {
	_, err := dbh.Connection.Exec(
		`CREATE TABLE IF NOT EXISTS users (
    		   		telegram_id INTEGER UNIQUE PRIMARY KEY,
					nickname TEXT,
					state INTEGER);`)
	return err
}

//CreatePlayersTable creates a table of Players if one does not already exist
func (dbh *DBHandler) CreatePlayersTable() error {
	_, err := dbh.Connection.Exec(
		`CREATE TABLE IF NOT EXISTS players (
    		   		player_id INTEGER PRIMARY KEY,
					nickname TEXT,
					message TEXT,
					x INTEGER,
					y INTEGER,
					smallPic TEXT,
					bigPic TEXT,
					active INTEGER,
					health INTEGER,
					dexterity INTEGER,
					mastery INTEGER,
					damage INTEGER,
					speed INTEGER,
					visibility INTEGER,
					candies INTEGER,
					cakes INTEGER,
					gold INTEGER,
					FOREIGN KEY(player_id) REFERENCES users(telegram_id));`)
	return err
}

//InsertUser adds a user to the Users table
func (dbh *DBHandler) InsertUser(user api.User) error {
	//User structure is described in the api package file user.go
	_, err := dbh.Connection.Exec(`INSERT INTO users (telegram_id, nickname, state) VALUES (?, ?, ?);`, user.ID, user.Username, user.State)

	return err
}

//InsertPlayer adds a user to the Users table
func (dbh *DBHandler) InsertPlayer(player models.Player) error {
	var active int = 0
	if player.Active {
		active = 1
	}
	_, err := dbh.Connection.Exec(`INSERT INTO players (player_id, nickname, message, x, y, smallPic, bigPic, active, health, dexterity, mastery, damage, speed, visibility, candies, cakes, gold) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`, player.PlayerId, player.ObjectName, player.Message, player.X, player.Y, player.SmallPic, player.BigPic, active, player.Health,
		player.Dexterity, player.Mastery, player.Damage, player.Speed, player.Visibility, player.ScoreCandy, player.ScoreCake, player.ScoreGold)

	return err
}

//Update updates any field in any table with new value
func (dbh *DBHandler) Update(table, field string, value interface{}, whereField string, whereValue interface{}) error {
	statement := fmt.Sprintf(`UPDATE %s SET %s = ? WHERE %s = ?;`, table, field, whereField)
	_, err := dbh.Connection.Exec(statement, value, whereValue)

	return err
}

// GetField returns value of field in given table in respect to to some parameter
func (dbh *DBHandler) GetField(table, field, whereField string, whereVal interface{}) interface{} {
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

//ContainsUser returns true if a user with the specified id is already registered
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

//ContainsPlayer returns true if a player with the specified id is already registered
func (dbh *DBHandler) ContainsPlayer(player models.Player) (bool, error) {
	result, err := dbh.Connection.Query(`SELECT * FROM players WHERE player_id = ?;`, player.PlayerId)

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
		return user, err
	}
	defer result.Close()
	if result.Next() {
		err := result.Scan(&user.ID, &user.Username, &user.State)
		return user, err
	}
	return user, err
}

//GetPlayerByID returns models.Player object from database with specified id
func (dbh *DBHandler) GetPlayerByID(id int) (*models.Player, error) {
	var player *models.Player = &models.Player{}
	result, err := dbh.Connection.Query(`SELECT * FROM players WHERE player_id = ?;`, id)
	if err != nil {
		return player, err
	}
	defer result.Close()
	if result.Next() {
		var active int = 1
		err := result.Scan(&player.PlayerId, &player.ObjectName, &player.Message, &player.X, &player.Y, &player.SmallPic, &player.BigPic, &active, &player.Health,
			&player.Dexterity, &player.Mastery, &player.Damage, &player.Speed, &player.Visibility, &player.ScoreCandy, &player.ScoreCake, &player.ScoreGold)
		if active == 0 {
			player.Active = false
		} else {
			player.Active = true
		}
		return player, err
	}
	return player, err
}
