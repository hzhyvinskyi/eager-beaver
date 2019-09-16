package dal

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Dal ...
type Dal struct {
	config *Config
	db	   *sql.DB
}

// New ...
func New(config *Config) *Dal {
	return &Dal{
		config: config,
	}
}

// Open ...
func (d *Dal) Open() error {
	db, err := sql.Open("postgres", d.config.databaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db

	return nil
}

// Close ...
func (d *Dal) Close() {
	d.db.Close()
}
