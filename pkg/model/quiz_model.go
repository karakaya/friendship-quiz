package model

import "github.com/google/uuid"

type Quiz struct {
	ID       uuid.UUID  `json:"_id"`
	Question []Question `json:"question"`
}

type Question struct {
	Options []Option `json:"option"`
}

type Option struct {
	IMG      string `json:"img"`
	IsAnswer bool   `json:"is_answer"`
}
