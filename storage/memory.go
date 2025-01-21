package storage

import (
	"database/sql"
	"time"
)

type SQLiteDB struct {
	*sql.DB
}

func InitDB(dataSourceName string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_code TEXT UNIQUE NOT NULL,
		long_url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &SQLiteDB{db}, nil
}

func (db *SQLiteDB) SaveLink(shortCode, longURL string, createdAt time.Time) error {
	_, err := db.Exec("INSERT INTO links (short_code, long_url, created_at) VALUES (?, ?, ?)",
		shortCode, longURL, createdAt)
	return err
}

func (db *SQLiteDB) GetLink(shortCode string) (string, error) {
	var longURL string
	err := db.QueryRow("SELECT long_url FROM links WHERE short_code = ?", shortCode).Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}
