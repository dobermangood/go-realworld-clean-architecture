package user

import "time"

type User struct {
	Name      string
	Email     string
	Password  string
	Bio       *string
	ImageLink *string
	FollowIDs []string
	// Favorites []Article
	CreatedAt time.Time
	UpdatedAt time.Time
}
