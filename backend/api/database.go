package api

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"spamtube/backend/domain"

	_ "github.com/lib/pq"
)

func QueryMySubReddits() ([]*domain.Subreddit, error) {

	var items []*domain.Subreddit
	var db *sql.DB
	dsn := flag.String("dsn", os.Getenv("DATABASE_URL"), "Postgres data source name")
	flag.Parse()

	db, err := openDB(*dsn)

	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM "subreddits"`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for rows.Next() {
		var id, name string
		var createdAt time.Time

		err = rows.Scan(&id, &name, &createdAt)
		if err != nil {
			log.Printf("Failed to build item: %v\n", err)
			return nil, err
		}

		item := &domain.Subreddit{
			Id:         id,
			Name:       name,
			Created_at: createdAt,
		}

		items = append(items, item)
	}

	return items, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
