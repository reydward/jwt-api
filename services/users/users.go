package users

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/entity"
	"errors"
	"github.com/go-pg/pg/v10"
)

func GetUsers(db *pg.DB) ([]entity.User, error) {
	users, err := domain.FetchAll(db)
	if err != nil {
		return nil, errors.New("Users not found")
	}
	return users, nil
}

func UpdateUser(db *pg.DB, user entity.User) (entity.User, error) {
	userUpdated, err := domain.Update(db, user)
	if err != nil {
		return entity.User{}, errors.New("Error updating user")
	}
	return userUpdated, nil
}
