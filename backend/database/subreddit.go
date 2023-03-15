package database

import (
	"log"
	"spamtube/backend/domain"
	"time"
)

func GetSubreddits() ([]*domain.Subreddit, error) {
	var items []*domain.Subreddit

	rows, err := DB.Query(`SELECT * FROM "subreddits"`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name string
		var cron *string
		var createdAt time.Time

		err = rows.Scan(&id, &name, &createdAt, &cron)
		if err != nil {
			log.Printf("Failed to build item: %v\n", err)
			return nil, err
		}

		item := &domain.Subreddit{
			Id:         id,
			Name:       name,
			Created_at: createdAt,
			Cron:       cron,
		}

		items = append(items, item)
	}

	return items, nil
}

func DeleteSubreddit(id string) error {
	_, err := DB.Exec(`DELETE FROM "subreddits" WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func UpsertSubreddit(sub domain.Subreddit) (domain.Subreddit, error) {
	query := `
	INSERT INTO subreddits (id, name, created_at, cron)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (id)
	DO
	UPDATE SET
		"cron" = $4
	RETURNING
	id,
	name,
	created_at,
	cron
	`
	subreddit := domain.Subreddit{}
	err := DB.QueryRow(query, sub.Id, sub.Name, time.Now(), sub.Cron).
		Scan(
			&subreddit.Id,
			&subreddit.Name,
			&subreddit.Created_at,
			&subreddit.Cron,
		)

	if err != nil {
		return domain.Subreddit{}, err
	}

	return subreddit, nil
}
