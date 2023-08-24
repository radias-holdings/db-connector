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
	log.Debug("USER: ", os.Getenv("DB_USER"))
	log.Debug("PASSWORD: ", os.Getenv("DB_PASSWORD"))
	log.Debug("DB ADDRESS: ", os.Getenv("DB_ADDRESS"))

	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_ADDRESS"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error("Failed to open DB connection", "rsp", err)
		return nil, err
	}

	return db, err
}
