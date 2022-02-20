package domain

import (
	"cc_eduardherrera_BackendAPI/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	user          entity.User
}

func TestFetchByEmail(t *testing.T) {
	db := GetDBConnection()
	defer db.Close()
	err := CreateSchema(db)
	if err != nil {
		panic(err)
	}

	user := entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}
	_, err = Create(db, user)
	if err != nil {
		t.Error("Error creating user for test")
	}

	userByEmail, err := FetchByEmail(db, "test@axiomzen.co")
	if err != nil {
		t.Error("Error fetching user by email for test")
	}
	assert.Equal(t, user, userByEmail)
}

func TestFetchByEmailError(t *testing.T) {
	db := GetDBConnection()
	defer db.Close()
	err := CreateSchema(db)
	if err != nil {
		panic(err)
	}

	user := entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}
	_, err = Create(db, user)
	if err != nil {
		t.Error("Error creating user for test")
	}

	userByEmail, err := FetchByEmail(db, "testx@axiomzen.co")
	if err != nil {
		assert.Equal(t, entity.User{}, userByEmail)
	}
}

func TestFetchAll(t *testing.T) {
	db := GetDBConnection()
	defer db.Close()
	err := CreateSchema(db)
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
		userCreated, err := Create(db, testCase.user)
		if err != nil {
			t.Error("Error creating user for test")
		}
		assert.Equal(t, testCase.user, userCreated)
	}

	allUsers, err := FetchAll(db)
	if err != nil {
		t.Error("Error fetching all users for test")
	}
	assert.Equal(t, expectedResult, allUsers)
}

func TestCreate(t *testing.T) {
	db := GetDBConnection()
	defer db.Close()
	err := CreateSchema(db)
	if err != nil {
		panic(err)
	}

	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}},
		{entity.User{"user1@dapper.com", "password1", "firstname1", "lastname1"}},
	}

	for _, testCase := range testCases {
		userCreated, err := Create(db, testCase.user)
		if err != nil {
			t.Error("Error creating user for test")
		}
		assert.Equal(t, testCase.user, userCreated)
	}
}

func TestUpdate(t *testing.T) {
	db := GetDBConnection()
	defer db.Close()
	err := CreateSchema(db)
	if err != nil {
		panic(err)
	}

	user := entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}
	userExpected := entity.User{"test@axiomzen.co", "axiomzen", "Alex-modified", "Zimmerman-modified"}
	_, err = Create(db, user)
	if err != nil {
		t.Error("Error creating user for test")
	}

	user.Firstname = "Alex-modified"
	user.Lastname = "Zimmerman-modified"
	userUpdated, err := Update(db, user)
	if err != nil {
		t.Error("Error updatin user for test")
	}
	assert.Equal(t, userExpected, userUpdated)
}
