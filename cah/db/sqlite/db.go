package sqlite

import (
	"database/sql"

	// driver for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB(dbFileName string) *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	return db
}

func createTables(db *sql.DB) {
	createTableWhiteCard(db)
	createTableBlackCard(db)
}
