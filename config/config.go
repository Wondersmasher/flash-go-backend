package config

import (
	"os"
)

var (
	MONGO_DB_URL           string
	SALT                   string
	JWT_SECRET_ACCESS_KEY  string
	JWT_SECRET_REFRESH_KEY string
	PORT                   string
	MONGO_DB_DATABASE      string
	MONGO_DB_COLLECTION    string
	REDIS_ADDRESS          string
	REDIS_USERNAME         string
	REDIS_DB               int64
	REDIS_PASSWORD         string
)

func Env() {
	MONGO_DB_URL = os.Getenv("MONGO_DB_URL")
	MONGO_DB_COLLECTION = os.Getenv("MONGO_DB_COLLECTION")
	MONGO_DB_DATABASE = os.Getenv("MONGO_DB_DATABASE")
	SALT = os.Getenv("SALT")
	JWT_SECRET_ACCESS_KEY = os.Getenv("JWT_SECRET_ACCESS_KEY")
	JWT_SECRET_REFRESH_KEY = os.Getenv("JWT_SECRET_REFRESH_KEY")
	PORT = os.Getenv("PORT")
	REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
	REDIS_USERNAME = os.Getenv("REDIS_USERNAME")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}
