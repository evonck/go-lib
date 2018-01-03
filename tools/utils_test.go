package tools

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
)

func TestIsInArray(t *testing.T) {
	log.Print("TestIsInArray")
	var testArray []string
	testArray = append(testArray, "test")
	testArray = append(testArray, "test1")
	testArray = append(testArray, "test2")
	convey.Convey("test1 should be in array", t, func() {
		convey.So(IsInArray("test1", testArray), convey.ShouldEqual, true)
	})
	convey.Convey("test2 should be in array", t, func() {
		convey.So(IsInArray("test2", testArray), convey.ShouldEqual, true)
	})
	convey.Convey("test should be in array", t, func() {
		convey.So(IsInArray("test", testArray), convey.ShouldEqual, true)
	})
	convey.Convey("test4 should not be in array", t, func() {
		convey.So(IsInArray("test4", testArray), convey.ShouldEqual, false)
	})
}

func TestCheckStringParam(t *testing.T) {
	log.Print("TestCheckStringParam")
	test := ""
	test2 := "test2"
	test3 := "test2"
	convey.Convey("test1 should be in array", t, func() {
		convey.So(CheckStringParam(test, test3, test2), convey.ShouldEqual, false)
	})
	convey.Convey("test1 should be in array", t, func() {
		convey.So(CheckStringParam(test3, test2), convey.ShouldEqual, true)
	})
}

func TestCheckSlash(t *testing.T) {
	log.Print("TestCheckSlash")
	url := "http://localhost:8081"
	url2 := CheckSlash(url)
	convey.Convey("url2 should be in http://localhost:8081/", t, func() {
		convey.So(url2, convey.ShouldEqual, "http://localhost:8081/")
	})
	url3 := CheckSlash(url2)
	convey.Convey("url3 should be in http://localhost:8081/", t, func() {
		convey.So(url3, convey.ShouldEqual, "http://localhost:8081/")
	})
}

func TestExist(t *testing.T) {
	log.Print("TestExist")
	exist, err := Exists("./utils.go")
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey("exist should be true", t, func() {
		convey.So(exist, convey.ShouldEqual, true)
	})
	exist2, err := Exists("./utilss.go")
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	convey.Convey("exist2 should be true", t, func() {
		convey.So(exist2, convey.ShouldEqual, false)
	})
}
