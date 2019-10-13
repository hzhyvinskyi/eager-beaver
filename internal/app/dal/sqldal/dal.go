package sqldal

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/hzhyvinskyi/eager-beaver/internal/app/dal"
)

// Dal ...
type Dal struct {
	db				*sql.DB
	userRepository	*UserRepository
}

// New ...
func New(db *sql.DB) *Dal {
	return &Dal{
		db: db
	}
}

// User ...
func (d *Dal) User() dal.UserRepository {
	if d.userRepository != nil {
		return d.userRepository
	}

	d.userRepository = &UserRepository{
		dal: d,
	}

	return d.userRepository
}
