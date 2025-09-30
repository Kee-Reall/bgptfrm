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

type UserInputDto struct {
	Login    string `json:"login" validate:"required,min=3,max=10"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
}
