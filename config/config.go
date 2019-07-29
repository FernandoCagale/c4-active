package config

import (
	"os"

	"github.com/joho/godotenv"
)

type mongoDB struct {
	Addrs    string
	Database string
	Username string
	Password string
}

//Config env connection
type Config struct {
	DatastoreURL string
	MongoDB      *mongoDB
}

func init() {
	godotenv.Load()
}

//LoadEnv loang env
func LoadEnv() *Config {
	return &Config{
		DatastoreURL: os.Getenv("DATASTORE_URL"),
		MongoDB: &mongoDB{
			Addrs:    os.Getenv("MONGO_ADDRS"),
			Database: os.Getenv("MONGO_DATABASE"),
			Username: os.Getenv("MONGO_USERNAME"),
			Password: os.Getenv("MONGO_PASSWORD"),
		},
	}
}
