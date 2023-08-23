package connector

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	c "github.com/smartystreets/goconvey/convey"
)

func TestNewConnection(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Error("No .env.test file found", "rsp", err)
	}

	db, err := NewConnection()
	if err != nil {
		log.Error("No DB connection", "rsp", err)
	}
	defer db.Close()

	c.Convey(".env.test environment variables", t, func() {
		c.Convey("should load without error", func() {
			c.So(err, c.ShouldBeNil)
		})
		c.Convey("should be set correctly", func() {
			// Requires a DB username of "local", change this if required
			c.So(os.Getenv("DB_USER"), c.ShouldEqual, "local")
			c.So(os.Getenv("DB_PASSWORD"), c.ShouldNotBeEmpty)
		})
	})
	c.Convey("Calling NewConnection", t, func() {
		c.Convey("should return valid connection without error", func() {
			c.So(db, c.ShouldNotBeNil)
			c.So(err, c.ShouldBeNil)
		})
	})
}
