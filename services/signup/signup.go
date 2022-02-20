package signup

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/entity"
	"cc_eduardherrera_BackendAPI/services/token"
	"errors"
	"github.com/go-pg/pg/v10"
	"strings"
)

func Signup(db *pg.DB, user entity.User) (string, error) {
	token, err := token.GetToken(user)
	if err != nil {
		return "", err
	}

	_, err = domain.Create(db, user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(db *pg.DB, user entity.User) (string, error) {
	userFromDB, err := domain.FetchByEmail(db, user.Email)
	if err != nil {
		return "", errors.New("User not found")
	}

	if strings.Compare(userFromDB.Email, user.Email) != 0 || strings.Compare(userFromDB.Password, user.Password) != 0 {
		return "", errors.New("Username or Password incorrect")
	}

	token, err := token.GetToken(userFromDB)
	if err != nil {
		return "", err
	}

	return token, nil
}
