package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/onekloud/go-lib/common"
	"github.com/julienschmidt/httprouter"
	"github.com/smartystreets/goconvey/convey"
)

var (
	arrayHello []interface{}
)

func init() {
	arrayHello = append(arrayHello, hello)
	arrayHello = append(arrayHello, hello)
}

func TestingInterface(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{ "test" : "Hello"}`)
}

func TestingArrayInterface(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `[{ "test" : "Hello"},{ "test" : "Hello"}]`)
}

// Handlers Returns httprouter handlers
func TestingHandlers2() *httprouter.Router {
	r := httprouter.New()
	r.POST("/testInterface", TestingInterface)
	r.POST("/testArray", TestingArrayInterface)
	r.POST("/testMap", TestingInterface)
	return r
}

func Server2() {
	handler := &common.Handlers{Handle: TestingHandlers2}
	router := handler.Handle()
	ServerTest = httptest.NewServer(router)
}

func TestDecodeEventInterface(t *testing.T) {
	Server2()
	defer ServerTest.Close()
	resp := POST(ServerTest.URL+"/testInterface", nil, nil, nil)
	convey.Convey("DecodeEventInterface should be hello", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, hello)
	})
}

func TestDecodeMapStringInterface(t *testing.T) {
	Server2()
	defer ServerTest.Close()
	resp := POST(ServerTest.URL+"/testMap", nil, nil, nil)
	convey.Convey("DecodeEventInterface should be hello", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, hello)
	})
}

func TestDecodeEventArrayInterface(t *testing.T) {
	Server2()
	defer ServerTest.Close()
	resp := POST(ServerTest.URL+"/testArray", nil, nil, nil)
	convey.Convey("DecodeEventInterface should be arrayHello", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, arrayHello)
	})
}
