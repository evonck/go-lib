package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"bitbucket.org/onekloud/go-lib/common"
	"github.com/julienschmidt/httprouter"
	"github.com/smartystreets/goconvey/convey"
)

var (
	// ServerTest the testing server
	ServerTest *httptest.Server
	// Router the router
	Router *httprouter.Router
	u      *url.URL
	hello  map[string]interface{}
)

func init() {
	hello = make(map[string]interface{})
	hello["test"] = "Hello"
}

func Testing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"test":"Hello"}`)
}

// Handlers Returns httprouter handlers
func TestingHandlers() *httprouter.Router {
	r := httprouter.New()
	r.POST("/test", Testing)
	r.PUT("/test", Testing)
	r.GET("/test", Testing)
	r.DELETE("/test", Testing)
	return r
}

func TestAddQuery(t *testing.T) {
	queries := make(map[string]string)
	queries["testParam"] = "testValue"
	urlExpected := "http://localhost?testParam=testValue"
	url := addQuery("http://localhost", queries)
	convey.Convey("url should be urlExpected", t, func() {
		convey.So(url, convey.ShouldEqual, urlExpected)
	})
}

func TestAddheader(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost", nil)
	header := make(map[string]string)
	header["Testheader"] = "testValue"
	req = addHeader(req, header)
	convey.Convey("req.Header.Testheader should be testValue", t, func() {
		convey.So(req.Header.Get("Testheader"), convey.ShouldEqual, "testValue")
	})
}

func Server() {
	handler := &common.Handlers{Handle: TestingHandlers}
	router := handler.Handle()
	ServerTest = httptest.NewServer(router)
}

func TestPOST(t *testing.T) {
	Server()
	defer ServerTest.Close()
	resp := POST(ServerTest.URL+"/test", nil, nil, nil)
	convey.Convey("resp should not be nil", t, func() {
		convey.So(resp, convey.ShouldNotEqual, nil)
	})
	if resp == nil {
		return
	}
	convey.Convey("resp.StatusCode should be 200", t, func() {
		convey.So(resp.StatusCode, convey.ShouldEqual, 200)
	})
	convey.Convey("DecodeEventInterface should be { test : Hello", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, hello)
	})
}

func TestPUT(t *testing.T) {
	Server()
	defer ServerTest.Close()
	resp := PUT(ServerTest.URL+"/test", nil, nil, nil)
	convey.Convey("resp should not be nil", t, func() {
		convey.So(resp, convey.ShouldNotEqual, nil)
	})
	if resp == nil {
		return
	}
	convey.Convey("resp.StatusCode should be 200", t, func() {
		convey.So(resp.StatusCode, convey.ShouldEqual, 200)
	})
	convey.Convey("DecodeEventInterface should be { test : Hello", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, hello)
	})
}

func TestGET(t *testing.T) {
	Server()
	defer ServerTest.Close()
	resp := GET(ServerTest.URL+"/test", nil, nil)
	convey.Convey("resp should not be nil", t, func() {
		convey.So(resp, convey.ShouldNotEqual, nil)
	})
	if resp == nil {
		return
	}
	convey.Convey("resp.StatusCode should be 200", t, func() {
		convey.So(resp.StatusCode, convey.ShouldEqual, 200)
	})
	convey.Convey("DecodeEventInterface should be { test : Hello }", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, hello)
	})
}

func TestDELETE(t *testing.T) {
	Server()
	defer ServerTest.Close()
	resp := DELETE(ServerTest.URL+"/test", nil, nil)
	convey.Convey("resp should not be nil", t, func() {
		convey.So(resp, convey.ShouldNotEqual, nil)
	})
	if resp == nil {
		return
	}
	convey.Convey("resp.StatusCode should be 200", t, func() {
		convey.So(resp.StatusCode, convey.ShouldEqual, 200)
	})
	convey.Convey("DecodeEventInterface should be { test : Hello", t, func() {
		convey.So(DecodeEventInterface(resp.Body), convey.ShouldResemble, hello)
	})
}
