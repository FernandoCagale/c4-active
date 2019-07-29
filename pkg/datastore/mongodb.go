package datastore

import (
	"gitlab.com/FernandoCagale/ramicard-usuario/config"
	mgo "gopkg.in/mgo.v2"
)

//NewMongoDB create connection
func NewMongoDB(env *config.Config) (*mgo.Session, error) {
	Host := []string{
		env.MongoDB.Addrs,
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Database: env.MongoDB.Database,
		Username: env.MongoDB.Username,
		Password: env.MongoDB.Password,
		FailFast: true,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}
