package models

import (
	"errors"
)

var ErrNoRecord = errors.New("model: no record found in database")
type Snippet struct{
	ID int
	TITLE string
	CONTENT string
	CREATED time.
}