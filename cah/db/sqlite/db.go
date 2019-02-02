package sqlite

import (
	"database/sql"

	// driver for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB(dbFileName string) {
	var err error
	db, err = sql.Open("sqlite3", dbFileName)
	if err != nil {
		panic(err)
	}
	createTables()
}

func createTables() {
	statement, err := db.Prepare(createTablesScript)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}
