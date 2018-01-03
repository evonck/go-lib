package tools

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/smartystreets/goconvey/convey"
)

var (
	mapInterface     map[interface{}]interface{}
	arrayOfInterface []interface{}
)

func TestCheckMapInterfaceInterface(t *testing.T) {
	log.Print("TesCheckMapInterfaceInterface")
	mapInterface = make(map[interface{}]interface{})
	convey.Convey("CheckMapInterfaceInterface(mapInterface) should be true", t, func() {
		convey.So(CheckMapInterfaceInterface(mapInterface), convey.ShouldEqual, true)
	})
	arrayOfInterface = make([]interface{}, 0)
	convey.Convey("CheckMapInterfaceInterface(arrayOfInterface) should be false", t, func() {
		convey.So(CheckMapInterfaceInterface(arrayOfInterface), convey.ShouldEqual, false)
	})
}

func TestCheckArrayInterface(t *testing.T) {
	log.Print("TestCheckArrayInterface")
	convey.Convey("CheckMapInterfaceInterface(mapInterface) should be true", t, func() {
		convey.So(CheckArrayInterface(mapInterface), convey.ShouldEqual, false)
	})
	convey.Convey("CheckMapInterfaceInterface(arrayOfInterface) should be false", t, func() {
		convey.So(CheckArrayInterface(arrayOfInterface), convey.ShouldEqual, true)
	})
}

func TestCheckString(t *testing.T) {
	log.Print("TestCheckString")
	convey.Convey("CheckMapInterfaceInterface(mapInterface) should be true", t, func() {
		convey.So(CheckString(mapInterface), convey.ShouldEqual, false)
	})
	convey.Convey("CheckMapInterfaceInterface(arrayOfInterface) should be false", t, func() {
		convey.So(CheckString(arrayOfInterface), convey.ShouldEqual, false)
	})
	testString := "test"
	convey.Convey("CheckMapInterfaceInterface(arrayOfInterface) should be false", t, func() {
		convey.So(CheckString(testString), convey.ShouldEqual, true)
	})
}
