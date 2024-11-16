package persist

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessLevel int    `json:"access_level"`
	Password    string `json:"password"`
}

func setupTables(db *sql.DB) {
	_, err := db.Exec(`
		create table if not exists users (
			id serial primary key,
			email varchar(255) not null,
			name varchar(255),
			access_level integer not null default 1,
			password varchar(255) not null,
			unique(email)
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating test table: %v", err))
	}
	//seed a user should remove if ever really hosted
	email := "test@email.com"
	name := "name"
	level := 1
	password := "password"

	row := db.QueryRow(`select id from users where email = ($1);`, email)

	var us User

	switch err := row.Scan(&us.Id); err {
	case sql.ErrNoRows:
		_, err = db.Exec(`Insert into users (email, name, access_level, password) values ($1, $2, $3, $4);`, email, name, level, password)
		if err != nil {
			panic(fmt.Sprintf("error inserting into db: %v", err))
		}
		// return 0, sql.ErrNoRows
		// fmt.Println("No rows were returned!")
	case nil:
		break
	default:
		panic(err)
	}

	_, err = db.Exec(`
		create table if not exists printers (
			id serial primary key,
			name varchar(255) not null,
			description varchar(255) not null default '',
			is_active boolean,
			unique(name)
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating printer table: %v", err))
	}

	_, err = db.Exec(`
		create table if not exists classes (
			id serial primary key,
			name varchar(255),
			description varchar(255),
			is_active boolean,
			unique(name)
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating printer table: %v", err))
	}

	_, err = db.Exec(`
		create table if not exists class_printers (
			id serial primary key,
			class_id integer,
			printer_id integer,
			unique(class_id, printer_id),

			constraint fk_class_printers_class foreign key (class_id) references classes(id),
			constraint fk_class_printers_printer foreign key (printer_id) references printers(id)
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating printer table: %v", err))
	}

	_, err = db.Exec(`
		create table if not exists class_users (
			id serial primary key,
			user_id integer,
			class_id integer,
			unique(class_id, user_id),

			constraint fk_class_users_user foreign key (user_id) references users(id),
			constraint fk_class_users_class foreign key (class_id) references classes(id)
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating printer table: %v", err))
	}

	_, err = db.Exec(`
		create table if not exists class_times(
			id serial primary key,
			class_id integer,
			start_time timestamp,
			end_time timestamp
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating class_times table: %v", err))
	}

	_, err = db.Exec(`
		create table if not exists printer_times (
			id serial primary key,
			printer_id integer,
			start_time timestamp,
			end_time timestamp
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating printers table: %v", err))
	}

	_, err = db.Exec(`
		create table if not exists user_bookings(
			id serial primary key,
			user_id integer,
			printer_id integer,
			start_time timestamp,
			end_time timestamp
		);`,
	)
	if err != nil {
		panic(fmt.Sprintf("error creating printers table: %v", err))
	}
}

func NewDB() *sql.DB {
	godotenv.Load()

	host := os.Getenv("DB_HOST")

	// Define the connection string (replace with your own database credentials)
	connStr := fmt.Sprintf("user=user dbname=printshift password=secret host=%s port=5432 sslmode=disable", host)

	maxTries := 5

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	for i := 0; i <= maxTries; i++ {
		err = db.Ping()
		if err != nil && i == maxTries {
			slog.Error("Error pinging the database: ", err)
			panic("Error pinging db")
		} else if err != nil {
			slog.Error("couldn't connect to db. Retrying...")
			time.Sleep(3 * time.Second)
			continue
		} else {
			break
		}
	}

	setupTables(db)
	return db
}
