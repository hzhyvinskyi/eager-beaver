package sqldal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hzhyvinskyi/eager-beaver/internal/app/dal/sqldal"
	"github.com/hzhyvinskyi/eager-beaver/internal/app/model"
)

func TestUserRepository_Create(t *testing.T) {
	d, teardown := sqldal.TestDB(t, databaseURL)
	defer teardown("users")

    u := model.TestUser
	err := d.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	d, teardown := sqldal.TestDB(t, databaseURL)
	defer teardown("users")

	email := "user@user.com"

	_, err := d.User().FindByEmail(email)

	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email

	_, err = d.User().Create(u)

	u, err = d.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
