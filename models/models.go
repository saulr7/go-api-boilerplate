package models

import "errors"

var (
	ErrPersonCannotBeNil = errors.New("Person cannot be nil")
	ErrIDPersonNoExits   = errors.New("Person ID doesn't exist")
)
