package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ArticleBrief struct {
	ID        uuid.UUID
	Title     string
	CreatedAt time.Time
}

func (b ArticleBrief) MarshalJSON() ([]byte, error) {
	j := struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		CreatedAt string `json:"created_at"`
	}{b.ID.String(), b.Title, b.CreatedAt.Format(time.RFC3339)}

	return json.Marshal(j)
}
