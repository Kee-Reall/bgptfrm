package shared

import "time"

type User struct {
	id        string
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"createdAt"`
	updateAt  time.Time
	deletedAt time.Time
	hash      string
	salt      string
}
