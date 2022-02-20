package token

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/entity"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type testCase struct {
	user          entity.User
	expectedToken string
}

func TestGetToken(t *testing.T) {
	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28ifQ.ou-bfwZzfIQeaO4KQX3xQOcNO247MxluWt-q2jQlISw",},
		{entity.User{"test2@axiomzen.co", "password1", "firstname1", "lastname1"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3QyQGF4aW9temVuLmNvIn0.y4WCAnlxE3fvRNZrhAb-13Y_kbOEf1iHdt5UqN_-wzI",},
	}

	for _, testCase := range testCases {
		token, err := GetToken(testCase.user)
		if err != nil {
			assert.Error(t, errors.New("Error getting the token."))
		}

		splitToken := strings.Split(token, ".")
		splitExpectedToken := strings.Split(testCase.expectedToken, ".")
		if strings.Compare(splitToken[0], splitExpectedToken[0]) != 0 || strings.Compare(splitToken[1], splitExpectedToken[1]) != 0 {
			assert.Error(t, errors.New("Unexpected token."))
		}
	}
}

func TestParseToken(t *testing.T) {
	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28ifQ.IOvmA5zueHCtZKG2gfp6KkEdYc0Uvyn18vhH1MHTkcs"},
		{entity.User{"test2@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3QyQGF4aW9temVuLmNvIn0.qabZC3C0xQ-kgaNAOYJckbPItV6YTGXxpL-_hBw5XqY"},
	}

	for _, testCase := range testCases {
		token, err := ParseToken(testCase.expectedToken)
		if err != nil {
			assert.Error(t, errors.New("Error getting the token."))
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if strings.Compare(fmt.Sprint(claims["Email"]), testCase.user.Email) != 0 {
				assert.Error(t, errors.New("Unexpected token."))
			}
		} else {
			assert.Error(t, errors.New("Unexpected token."))
		}
	}

}

func TestGetEmailFromToken(t *testing.T)  {
	testCases := []testCase{
		{entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28ifQ.IOvmA5zueHCtZKG2gfp6KkEdYc0Uvyn18vhH1MHTkcs"},
		{entity.User{"test2@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3QyQGF4aW9temVuLmNvIn0.qabZC3C0xQ-kgaNAOYJckbPItV6YTGXxpL-_hBw5XqY"},
	}

	expectedEmails := []string{"test@axiomzen.co", "test2@axiomzen.co"}

	for i, testCase := range testCases {
		email, err := GetEmailFromToken(testCase.expectedToken)
		if err != nil {
			assert.Error(t, errors.New("Error updating user for test"))
		}

		assert.Equal(t, expectedEmails[i], email)
	}
}

func TestIsValidToken(t *testing.T)  {
	db := domain.GetDBConnection()
	defer db.Close()
	err := domain.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	user := entity.User{"test@axiomzen.co", "axiomzen", "Alex", "Zimmerman"}
	_, err = domain.Create(db, user)
	if err != nil {
		t.Error("Error creating user for test")
	}

	isValidToken := IsValidToken(db, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlc3RAYXhpb216ZW4uY28ifQ.IOvmA5zueHCtZKG2gfp6KkEdYc0Uvyn18vhH1MHTkcs")

	assert.Equal(t, true, isValidToken)
}
