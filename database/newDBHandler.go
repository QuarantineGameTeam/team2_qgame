package database

const (
	//Driver is a name of db driver (requires go get github.com/mattn/go-sqlite3)
	Driver = "sqlite3"
	//Path is a path to file with database
	Path = "database/CandyWarGoDatabase.sqlite"
)

//NewDBHandler returns pointer to the default ready to use DBHandler
func NewDBHandler() (*DBHandler, error) {
	dbh := &DBHandler{
		DriverName: Driver,
		DBPath:     Path,
	}
	err := dbh.Connect()
	if err != nil {
		return dbh, err
	}

	err = dbh.CreateTables()
	return dbh, err
}

//NewDBHandler returns pointer to the default ready to use DBHandler
func NewDBHandlerWithPath(path string) (*DBHandler, error) {
	dbh := &DBHandler{
		DriverName: Driver,
		DBPath:     path,
	}
	err := dbh.Connect()
	if err != nil {
		return dbh, err
	}

	err = dbh.CreateTables()
	return dbh, err
}
