package connector

import (
	"database/sql"
	"os"

	"github.com/fzer0zer0/logger"
	"github.com/go-sql-driver/mysql"
)

var log = logger.NewLogger(nil)

// New Connection opens a database connection to the MySQL DB specified in the environment variables
func NewConnection() (db *sql.DB, err error) {
	log.Debug("DB_USER: ", os.Getenv("DB_USER"))
	log.Debug("DB_PASSWORD: ", os.Getenv("DB_PASSWORD"))
	log.Debug("DB_ADDRESS: ", os.Getenv("DB_ADDRESS"))
	log.Debug("DB_NAME: ", os.Getenv("DB_NAME"))

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDRESS"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error("Failed to open DB connection", "rsp", err)
		return nil, err
	}

	return db, err
}
