package domain

import "github.com/ryota1119/gin_webapi/internal/schema"

// Post 投稿
type Post struct {
	ID     uint
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID schema.UserID
}
