package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func Init() (err error) {
	db, err = sql.Open("sqlite", "scamalytics.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			passwordhash TEXT NOT NULL,
			joined TEXT,
			sid TEXT)`)

	return
}

func Close() {
	db.Close()
}
