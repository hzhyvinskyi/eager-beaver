package dal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hzhyvinskyi/eager-beaver/internal/app/dal"
	"github.com/hzhyvinskyi/eager-beaver/internal/app/model"
)

func TestUserRepository_Create(t *testing.T) {
	d, teardown := dal.TestDal(t, databaseURL)
	defer teardown("users")

	u, err := d.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	d, teardown := dal.TestDal(t, databaseURL)
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
