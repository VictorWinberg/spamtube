 package database

import (
	"log"
	"spamtube/backend/domain"
	"time"

	"github.com/google/uuid"
)

func InsertImgUrl(imgUrl string) (error) {

	err := DB.Exec(`INSERT INTO img (url) VALUES ($1)`, imgUrl)
	if err != nil {
		return err
	}

	return nil
}
