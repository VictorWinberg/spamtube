package database

import (
	"database/sql"
	"flag"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	dsn := flag.String("dsn", os.Getenv("DATABASE_URL"), "Postgres data source name")
	flag.Parse()

	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	DB = db

	return nil
}
