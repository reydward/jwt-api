package main

import (
	"cc_eduardherrera_BackendAPI/domain"
	"cc_eduardherrera_BackendAPI/server"
)

func main() {
	db := domain.GetDBConnection()
	defer db.Close()

	err := domain.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	dbConnection := server.DbConnection{
		DB: db,
	}
	server.Start(dbConnection)
}
