package database

func ResetTables() error {
	dbh, err := NewDBHandler()
	if err != nil {
		return err
	}

	_, err = dbh.Connection.Exec(`drop table games;
delete from users;
delete from players;
CREATE TABLE IF NOT EXISTS games (
                                     game_id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     game_json TEXT,
                                     player_id INTEGER,
                                     startmove_time INTEGER,
                                     players_json TEXT,
                                     red_spawn INTEGER,
                                     green_spawn INTEGER,
                                     blue_spawn INTEGER,
                                     state INTEGER);`)
	return err
}