package dal

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Dal ...
type Dal struct {
	config			*Config
	db				*sql.DB
	userRepository	*UserRepository
}

// New ...
func New(config *Config) *Dal {
	return &Dal{
		config: config,
	}
}

// Open ...
func (d *Dal) Open() error {
	db, err := sql.Open("postgres", d.config.DatabaseURL)
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

// User ...
func (d *Dal) User() *UserRepository {
	if d.userRepository != nil {
		return d.userRepository
	}

	d.userRepository = &UserRepository{
		dal: d,
	}

	return d.userRepository
}
