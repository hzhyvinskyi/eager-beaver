package dal

import "github.com/hzhyvinskyi/eager-beaver/internal/app/model"

// UserRepository ...
type UserRepository interface {
    Create(*model.User) error
    FindByEmail(string) (*model.User, error)
}
