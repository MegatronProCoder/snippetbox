package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("model: no record found in database")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
