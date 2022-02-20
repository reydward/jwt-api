package users

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/entity"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	user          entity.User
}

func TestGetUsers(t *testing.T) {
	db := domain.GetDBConnection()
	defer db.Close()
	err := domain.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}},
		{entity.User{"user1@dapper.com", "password1", "firstname1", "lastname1"}},
	}

	expectedResult := []entity.User{
		{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"},
		{"user1@dapper.com", "password1", "firstname1", "lastname1"},
	}

	for _, testCase := range testCases {
		_, err := domain.Create(db, testCase.user)
		if err != nil {
			assert.Error(t, errors.New("Error creating user for test"))
		}
	}

	users, err := GetUsers(db)
	if err != nil {
		assert.Error(t, errors.New("Error getting users for test"))
	}
	assert.Equal(t, expectedResult, users)
}

func TestUpdateUser(t *testing.T) {
	db := domain.GetDBConnection()
	defer db.Close()
	err := domain.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	user := entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}
	userExpected := entity.User{"test@axiomzen.co", "axiomzen", "Alex-modified", "Zimmerman-modified"}
	_, err = domain.Create(db, user)
	if err != nil {
		t.Error("Error creating user for test")
	}

	user.Firstname = "Alex-modified"
	user.Lastname = "Zimmerman-modified"
	userUpdated, err := UpdateUser(db, user)
	if err != nil {
		assert.Error(t, errors.New("Error updating user for test"))
	}

	assert.Equal(t, userExpected, userUpdated)
}