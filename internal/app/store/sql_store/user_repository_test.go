package sql_store_test

import (
	"RestAPI/internal/app/model"
	"RestAPI/internal/app/store"
	"RestAPI/internal/app/store/sql_store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql_store.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sql_store.TestDB(t, databaseURL)
	defer teardown("users")

	s := sql_store.New(db)

	email := "durov@mail.ru"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
