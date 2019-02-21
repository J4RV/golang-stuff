package sqlite

import (
	"fmt"
	"strings"
)

func createTableWhiteCard() {
	createTable("white_card", []string{
		"text TEXT",
		"expansion TEXT",
		"CHECK(text <> '' AND expansion <> '' AND LENGTH(text) <= 180)",
	})
	createIndex("white_card", "expansion")
}

func createTableBlackCard() {
	createTable("black_card", []string{
		"text TEXT",
		"expansion TEXT",
		"blanks INTEGER",
		"CHECK(text <> '' AND expansion <> '' AND blanks > 0 AND LENGTH(text) <= 180)",
	})
	createIndex("black_card", "expansion")
}

func createTableUser() {
	createTable("user", []string{
		"username TEXT UNIQUE",
		"password TEXT",
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
		"CHECK(username <> '' AND password <> '' AND LENGTH(username) <= 36)",
	})
	createIndex("user", "username")
}

// methods for repetitive stuff

func createTable(table string, columns []string) {
	if len(columns) == 0 {
		panic("createTable method is for tables with at least one column")
	}
	// Using Sprintf since this internal method does not use user inputs
	statement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s INTEGER PRIMARY KEY AUTOINCREMENT,%s);", table, table, strings.Join(columns, ","))
	db.MustExec(statement)
}

func createIndex(table, column string) {
	indexName := fmt.Sprintf("%s_%s", table, column)
	// Using Sprintf since this internal method does not use user inputs
	createIndexStatement := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s);", indexName, table, column)
	db.MustExec(createIndexStatement)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
