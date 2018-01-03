package common

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("./", "configTest")
	convey.Convey("v should not be nil", t, func() {
		convey.So(v, convey.ShouldNotEqual, nil)
	})
	convey.Convey("v should not be nil", t, func() {
		convey.So(v, convey.ShouldNotEqual, nil)
	})
}

func TestConfig(t *testing.T) {
	test := Config()
	convey.Convey("test should not be nil", t, func() {
		convey.So(test, convey.ShouldNotEqual, nil)
	})
	addr := test.GetString("addr")
	convey.Convey("addr should not be localhost:8081", t, func() {
		convey.So(addr, convey.ShouldEqual, "localhost:8081")
	})
}

func TesGetLogLevel(t *testing.T) {
	lvl := getLogLevel()
	convey.Convey("lvl should not be Debug", t, func() {
		convey.So(lvl, convey.ShouldEqual, "Debug")
	})
}
