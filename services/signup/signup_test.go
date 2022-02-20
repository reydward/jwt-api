package signup

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	user          entity.User
	expectedToken string
}

func TestSignup(t *testing.T) {
	db := domain.GetDBConnection()
	defer db.Close()
	err := domain.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28ifQ.IOvmA5zueHCtZKG2gfp6KkEdYc0Uvyn18vhH1MHTkcs"},
		{entity.User{"test2@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3QyQGF4aW9temVuLmNvIn0.qabZC3C0xQ-kgaNAOYJckbPItV6YTGXxpL-_hBw5XqY"},
	}

	for _, testCase := range testCases {
		token, err := Signup(db, testCase.user)
		if err != nil {
			t.Error("Error getting the token.")
		}
		assert.Equal(t, testCase.expectedToken, token)
	}
}

func TestLogin(t *testing.T) {
	db := domain.GetDBConnection()
	defer db.Close()
	err := domain.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28ifQ.IOvmA5zueHCtZKG2gfp6KkEdYc0Uvyn18vhH1MHTkcs"},
		{entity.User{"test2@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3QyQGF4aW9temVuLmNvIn0.qabZC3C0xQ-kgaNAOYJckbPItV6YTGXxpL-_hBw5XqY"},
	}

	for _, testCase := range testCases {
		_, err := domain.Create(db, testCase.user)
		if err != nil {
			t.Error("Error creating user for test")
		}
		token, err := Login(db, testCase.user)
		if err != nil {
			t.Error("Error getting the token.")
		}
		assert.Equal(t, testCase.expectedToken, token)
	}
}
