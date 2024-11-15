package persist

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupTables(db *sql.DB) {
	_, err := db.Exec(`
		create table if not exists users (
			userId text not null,
			email text not null,
			name text,
			accessLevel int
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating test table: %v", err))
	}

	// _, err = db.Exec(`
	// 	create table if not exists classes (
	// 		classId string,
	// 		professor string,
	// 		name string,

	// 	);`,
	// )
	// if err != nil {
	// 	panic(fmt.Sprintf("error creating test table: %v", err))
	// }

}

func NewDB() *sql.DB {
	godotenv.Load()

	host := os.Getenv("DB_HOST")

	// Define the connection string (replace with your own database credentials)
	connStr := fmt.Sprintf("user=user dbname=printshift password=secret host=%s port=5432 sslmode=disable", host)

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
