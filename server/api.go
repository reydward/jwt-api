package server

import (
	"cc_eduardherrera_BackendAPI/entity"
	signupservice "cc_eduardherrera_BackendAPI/services/signup"
	"cc_eduardherrera_BackendAPI/services/token"
	"cc_eduardherrera_BackendAPI/services/users"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Eduard Reyes backend API for Dapper!")
}

func (d DbConnection) signup(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(w, r)

	token, err := signupservice.Signup(d.DB, user)
	if err != nil {
		internalServerError(w, err.Error())
	}

	jsonBytes, err := json.Marshal(entity.Token{Token: token})
	if err != nil {
		unauthorized(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (d DbConnection) login(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(w, r)

	token, err := signupservice.Login(d.DB, user)
	if err != nil {
		jsonBytes, err := json.Marshal(entity.Error{Error: err.Error()})
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	jsonBytes, err := json.Marshal(entity.Token{Token: token})
	if err != nil {
		unauthorized(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (d DbConnection) getUsers(w http.ResponseWriter, r *http.Request) {
	authenticationTokenHeader := r.Header.Get("x-authentication-token")
	if !token.IsValidToken(d.DB, authenticationTokenHeader) {
		unauthorized(w)
		return
	}

	users, err := users.GetUsers(d.DB)
	if err != nil {
		jsonBytes, err := json.Marshal(entity.Error{Error: err.Error()})
		if err != nil {
			internalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		internalServerError(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (d DbConnection) updateUser(w http.ResponseWriter, r *http.Request) {
	authenticationTokenHeader := r.Header.Get("x-authentication-token")
	if !token.IsValidToken(d.DB, authenticationTokenHeader) {
		unauthorized(w)
		return
	}

	user := getUserFromRequest(w, r)
	//Gets email from token instead payload
	email, err := token.GetEmailFromToken(authenticationTokenHeader)
	if err != nil {
		internalServerError(w, err.Error())
		return
	}
	user.Email = email

	userUpdated, err := users.UpdateUser(d.DB, user)
	if err != nil {
		internalServerError(w, err.Error())
		return
	}
	jsonBytes, err := json.Marshal(userUpdated)
	if err != nil {
		internalServerError(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func Start(u ApiUsers) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/signup", u.signup).Methods("POST")
	router.HandleFunc("/login", u.login).Methods("POST")
	router.HandleFunc("/users", u.getUsers).Methods("GET")
	router.HandleFunc("/users", u.updateUser).Methods("PUT")
	log.Println("listen on", "localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
