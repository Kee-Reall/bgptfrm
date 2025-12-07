package core

import "errors"

var (
	UserDoesntExits = errors.New("required user does not exists or banned")
)
