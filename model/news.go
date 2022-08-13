package model

import "time"

type News struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"body"`
	Updated time.Time `json:"updated_at"`
	Created time.Time `json:"created_at"`
}
