package domain

import (
	"cc_eduardherrera_BackendAPI/entity"
	"cc_eduardherrera_BackendAPI/tools"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetDBConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     tools.GetDotEnvVariable("DB_USER", "postgres"),
		Password: tools.GetDotEnvVariable("DB_PASSWORD", "postgres"),
	})
	return db
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
