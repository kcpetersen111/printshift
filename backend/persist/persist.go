package persist

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
)

func setupTables(db *sql.DB) {
	_, err := db.Exec(`
		create table if not exists test (
			test TEXT
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating test table: %v", err))
	}
	_, err = db.Exec(`
		create table if not exists users (
			userId text not null,
			email text not null,
			name text,
			classes text[],
			printers text[],
			printersCanAssign int
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating test table: %v", err))
	}
}

func NewDB() *sql.DB {
	// Define the connection string (replace with your own database credentials)
	connStr := "user=user dbname=printshift password=secret host=localhost port=5432 sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	if err = db.Ping(); err != nil {
		slog.Error("Error pinging the database: ", err)
		panic("Error pinging db")
	}

	setupTables(db)
	return db
}
