package dal

import "github.com/hzhyvinskyi/eager-beaver/internal/app/model"

// UserRepository ...
type UserRepository struct {
	dal	*Dal
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.dal.db.QueryRow(
		"INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		u.FirstName, u.LastName, u.Email, u.Password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
