package postgres

import (
	"fmt"
	"os"
)

const (
	DEFAULT_DB_USER = "root"
	DEFAULT_DB_HOST = "localhost"
	DEFAULT_DB_PORT = "4568"
	DEFAULT_DB_NAME = "thebeyond"
)

func DSN() string {
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = DEFAULT_DB_USER
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		panic("db password is not defined")
	}

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = DEFAULT_DB_HOST
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = DEFAULT_DB_PORT
	}

	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		dbName = DEFAULT_DB_NAME
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbName)
}
