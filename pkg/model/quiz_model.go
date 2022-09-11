package model

import "github.com/google/uuid"

type Quiz struct {
	ID        uuid.UUID `json:"_id" bson:"_id"`
	Questions []Question
}

type Question struct {
	Title   string   `json:"title"`
	Options []Option `json:"option"`
}

type Option struct {
	IMG      string `json:"img"`
	IsAnswer bool   `json:"is_answer"`
}
