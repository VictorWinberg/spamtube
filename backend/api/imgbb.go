package api

import (
	"encoding/json"
)

type ImgInputBody struct {
	url    string          `json:"ref"`
}