package connector

import (
	"database/sql"

	"github.com/fzer0zer0/logger"
	"github.com/go-sql-driver/mysql"
)

var log = logger.NewLogger(nil)

// New Connection opens a database connection to the MySQL DB specified in the environment variables
func NewConnection(user, password, address, dbname string) (db *sql.DB, err error) {
	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 address,
		DBName:               dbname,
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error("Failed to open DB connection", "rsp", err)
		return nil, err
	}

	return db, err
}
