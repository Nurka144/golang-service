package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func InitConnectDBPg() (*sqlx.DB, error) {
	host := viper.Get("database.pg.host")
	port := viper.Get("database.pg.port")
	user := viper.Get("database.pg.user")
	password := viper.Get("database.pg.password")
	dbname := viper.Get("database.pg.db")

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	logrus.Info("Успешное подключение к PostgreSQL!")

	return db, nil
}
