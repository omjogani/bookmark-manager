package model

import "time"

type BookmarkItem struct {
	BookmarkItemId int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	URL            string    `json:"url"`
	DateTime       time.Time `json:"date"`
}
