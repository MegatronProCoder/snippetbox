package mysql

import (
	"database/sql"
	"megatroncodrr/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet to database
func (*SnippetModel) Insert() {}

// give a snippet from database based on its id
func (*SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// return 10 most recently created snippet
func (*SnippetModel) Latest() {}
