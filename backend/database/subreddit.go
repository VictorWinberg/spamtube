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

func DeleteSubreddit(id string) error {
	_, err := DB.Exec(`DELETE FROM "subreddits" WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}
