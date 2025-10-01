package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		status TEXT NOT NULL DEFAULT 'active'

	);`

	createQuestionsTable := `
	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		up_votes INTEGER DEFAULT 0,
		down_votes INTEGER DEFAULT 0,
		status TEXT NOT NULL DEFAULT 'active',
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	createAnswersTable := `
	CREATE TABLE IF NOT EXISTS answers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		body TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		question_id INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		up_votes INTEGER DEFAULT 0,
		down_votes INTEGER DEFAULT 0,
		status TEXT NOT NULL DEFAULT 'active',
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (question_id) REFERENCES questions(id)
	);`

	createVotesTable := `
	CREATE TABLE IF NOT EXISTS votes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		answer_id INTEGER,
		question_id INTEGER,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		type INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (answer_id) REFERENCES answers(id),
		FOREIGN KEY (question_id) REFERENCES questions(id),
		CHECK (type IN (-1, 1)),
		CHECK ((answer_id IS NULL AND question_id IS NOT NULL) OR 
			   (answer_id IS NOT NULL AND question_id IS NULL))
	);`

	// Execute all table creation queries
	queries := []string{createUsersTable, createQuestionsTable, createAnswersTable, createVotesTable}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			panic("Could not create table: " + err.Error())
		}
	}
}
