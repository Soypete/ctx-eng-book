package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type SessionState struct {
	ID             int
	CurrentChapter int
	CurrentModule  int
	LastPosition   int
	LastModified   time.Time
}

func New(connStr string) (*DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	d := &DB{db}
	d.ensureSessionTable()
	return d, nil
}

func (d *DB) ensureSessionTable() {
	_, _ = d.Exec(`
		CREATE TABLE IF NOT EXISTS authorpedro_session (
			id SERIAL PRIMARY KEY,
			current_chapter INTEGER DEFAULT 0,
			current_module INTEGER DEFAULT 0,
			last_position INTEGER DEFAULT 0,
			last_modified TIMESTAMP DEFAULT NOW()
		)
	`)
}

func (d *DB) GetSession() (*SessionState, error) {
	row := d.QueryRow("SELECT id, current_chapter, current_module, last_position, last_modified FROM authorpedro_session LIMIT 1")

	var s SessionState
	err := row.Scan(&s.ID, &s.CurrentChapter, &s.CurrentModule, &s.LastPosition, &s.LastModified)
	if err == sql.ErrNoRows {
		_, err = d.Exec("INSERT INTO authorpedro_session (current_chapter, current_module) VALUES (0, 0)")
		if err != nil {
			return nil, err
		}
		return &SessionState{CurrentChapter: 0, CurrentModule: 0}, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (d *DB) SaveSession(chapter, module, position int) error {
	_, err := d.Exec(`
		UPDATE authorpedro_session
		SET current_chapter = $1, current_module = $2, last_position = $3, last_modified = NOW()
		WHERE id = (SELECT id FROM authorpedro_session LIMIT 1)
	`, chapter, module, position)
	return err
}

func (d *DB) Close() error {
	return d.DB.Close()
}
