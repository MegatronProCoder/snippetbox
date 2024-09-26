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
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title , content , created, expires
	FROM snippets 
	WHERE expires > UTC_TIMESTAMP && id = ?`
	row := *m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// return 10 most recently created snippet
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content , created, expires
	FROM snippets
	WHERE expires > UTC_TIMESTAMP ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
