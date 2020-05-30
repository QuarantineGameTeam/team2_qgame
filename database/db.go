package database

import (
	"database/sql"
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
