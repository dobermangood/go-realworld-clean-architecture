package article

import "time"

type Article struct {
	Slug        string
	Title       string
	Description string
	Body        string
	TagList     []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// FavoritedBy []User
	// Author      User
	// Comments    []Comment
}
