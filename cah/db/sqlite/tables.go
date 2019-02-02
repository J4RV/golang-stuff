package sqlite

var createTablesScript = `
CREATE TABLE IF NOT EXISTS white_card (
	white_card INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT,
	expansion TEXT
);
CREATE INDEX IF NOT EXISTS white_card_expansion ON white_card(expansion);

CREATE TABLE IF NOT EXISTS black_card (
	black_card INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT,
	expansion TEXT ,
	blanks INTEGER
);
CREATE INDEX IF NOT EXISTS black_card_expansion ON black_card(expansion);
`
