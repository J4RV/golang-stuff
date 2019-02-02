package sqlite

import (
	"database/sql"
	"fmt"
)

func createTable(table string, columns []string) {
	var createTableStatement string
	createTableStatement = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (`, table)
	for _, column := range columns {
		createTableStatement += column + ","
	}
	// remove last comma
	createTableStatement = createTableStatement[:len(createTableStatement)-1]
	createTableStatement += ");"
	statement, err := db.Prepare(createTableStatement)
	panicIfErr(err)
	_, err = statement.Exec()
	panicIfErr(err)
}

func createIndex(table, column string) {
	indexName := fmt.Sprintf("%s_%s", table, column)
	createIndexStatement := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s);", indexName, table, column)
	statement, err := db.Prepare(createIndexStatement)
	panicIfErr(err)
	_, err = statement.Exec()
	panicIfErr(err)
}

func createTableWhiteCard(db *sql.DB) {
	createTable("white_card", []string{
		"white_card INTEGER PRIMARY KEY AUTOINCREMENT",
		"text TEXT",
		"expansion TEXT",
		"CHECK(text <> '' AND expansion <> '')",
	})
	createIndex("white_card", "expansion")
}

func createTableBlackCard(db *sql.DB) {
	createTable("black_card", []string{
		"black_card INTEGER PRIMARY KEY AUTOINCREMENT",
		"text TEXT",
		"expansion TEXT",
		"blanks INTEGER",
		"CHECK(text <> '' AND expansion <> '' AND blanks > 0)",
	})
	createIndex("black_card", "expansion")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
