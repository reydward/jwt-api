package domain

import (
	"cc_eduardherrera_BackendAPI/entity"
	"errors"
	"github.com/go-pg/pg/v10"
)

func FetchByEmail(db *pg.DB, email string) (entity.User, error) {
	user := &entity.User{Email: email}
	err := db.Model(user).Where("email = ?", user.Email).Select()
	if err != nil {
		return entity.User{}, err
	}
	return *user, nil
}

func FetchAll(db *pg.DB) ([]entity.User, error) {
	var users []entity.User
	err := db.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Create(db *pg.DB, user entity.User) (entity.User, error) {
	_, err := db.Model(&user).Insert()
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func Update(db *pg.DB, user entity.User) (entity.User, error) {
	values := map[string]interface{}{
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
	}

	result, err := db.Model(&values).
		TableExpr("users").
		Where("email = ?", user.Email).
		Update()
	if err != nil {
		return entity.User{}, err
	}

	if result.RowsAffected() == 0 {
		return entity.User{}, errors.New("Nothing updated")
	}

	userUpdated, err := FetchByEmail(db, user.Email)
	if err != nil {
		return entity.User{}, err
	}

	return userUpdated, nil
}