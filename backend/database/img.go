 package database

import (
	"log"
	"spamtube/backend/domain"
	"time"

	"github.com/google/uuid"
)

func InsertImgUrl(url string) (error) {

	err := DB.Exec(`INSERT INTO img (url) VALUES ($1)`, url)
	if err != nil {
		return err
	}

	return nil
}

func SelectAllImages() ([]*domain.ImgBody{}, error) {
	var items []*domain.ImgBody

	rows, err := DB.Query("SELECT * FROM img")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var url string

		err = rows.Scan(&url)
		if err != nil {
			log.Printf("Failed to build item: %v\n", err)
			return nil, err
		}

		item := &domain.ImgBody{
			Url:         url,
		}

		items = append(items, item)
	}

	return items, nil
}
