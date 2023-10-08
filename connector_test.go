package connector

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestNewConnection(t *testing.T) {
	db, err := NewConnection("local", "password", "localhost:3306", "test")
	if err != nil {
		log.Error("No DB connection", "rsp", err)
	}
	defer db.Close()

	c.Convey(".env.test environment variables", t, func() {
		c.Convey("should load without error", func() {
			c.So(err, c.ShouldBeNil)
		})
	})
	c.Convey("Calling NewConnection", t, func() {
		c.Convey("should return valid connection without error", func() {
			c.So(db, c.ShouldNotBeNil)
			c.So(err, c.ShouldBeNil)
		})
	})
}
