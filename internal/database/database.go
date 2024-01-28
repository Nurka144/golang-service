package database

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "Aa123456"
	dbname   = "bookhood"
)

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	logrus.Info("Успешное подключение к PostgreSQL!")

	return db, nil
}
