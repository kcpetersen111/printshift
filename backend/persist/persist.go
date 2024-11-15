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
			id integer autoincrement primary key,
			email text not null,
			name text,
			accessLevel integer not null default 1,
			password text not null,
			unique(email)
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating test table: %v", err))
	}
	//seed a user should remove if ever really hosted
	email := "email"
	name := "name"
	level := 1
	password := "password"
	_, err = db.Exec(`Insert into users values ($1, $2, $3, $4);`, email, name, level, password)
	if err != nil {
		panic(fmt.Sprintf("error inserting into db: %v", err))
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
