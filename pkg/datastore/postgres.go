package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gitlab.com/FernandoCagale/ramicard-usuario/config"
)

//NewPostgres create connection
func NewPostgres(env *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", env.DatastoreURL)
	if err != nil {
		return nil, err
	}

	db.LogMode(false)
	db.SingularTable(true)

	return db, nil
}
