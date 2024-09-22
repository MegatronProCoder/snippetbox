package mysql

import (
	"database/sql"
	"megatroncodrr/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet to database
func (m *SnippetModel) Insert(title, content, expired string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP, DATE_ADD(UTC_TIMESTAMP , INTERVAL ? DAY))`

	res, err := m.DB.Exec(stmt, title, content, expired)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// give a snippet from database based on its id
func (*SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// return 10 most recently created snippet
func (*SnippetModel) Latest() {}
