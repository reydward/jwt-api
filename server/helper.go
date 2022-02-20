package server

import (
	"cc_eduardherrera_BackendAPI/entity"
	"encoding/json"
	"github.com/go-pg/pg/v10"
	"net/http"
)

type ApiUsers interface {
	signup(w http.ResponseWriter, r *http.Request)
	login(w http.ResponseWriter, r *http.Request)
	getUsers(w http.ResponseWriter, r *http.Request)
	updateUser(w http.ResponseWriter, r *http.Request)
}

type DbConnection struct {
	DB *pg.DB
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}

func internalServerError(w http.ResponseWriter, error string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(error))
}

func getUserFromRequest(w http.ResponseWriter, r *http.Request) entity.User {
	user := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		internalServerError(w, err.Error())
	}
	return user
}
