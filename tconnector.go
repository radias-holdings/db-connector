package connector

import (
	"database/sql"
	"os"
)

// A convenience constructor that we use in tests
func NewTestConnection() (db *sql.DB, err error) {
	dbaddress := os.Getenv("DB_ADDRESS")
	dbname := os.Getenv("DB_NAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbusername := os.Getenv("DB_USER")

	db, err = NewConnection(dbusername, dbpassword, dbaddress, dbname)
	return db, err
}
