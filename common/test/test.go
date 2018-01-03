// Package test provide basic test function to test any api endpoint
package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bitbucket.org/onekloud/go-lib/common"
	"github.com/julienschmidt/httprouter"
)

var (
	// ServerTest the testing server
	ServerTest *httptest.Server
	// Router the router
	Router      *httprouter.Router
	HandlerFunc *common.Handlers
)

// Init initialize the test
func Init(handlerFunc *common.Handlers, pathFile, fileName string) {
	HandlerFunc = handlerFunc
	common.LoadConfig(pathFile, fileName)
	if handlerFunc.Handle != nil {
		Router = handlerFunc.Handle()
		ServerTest = httptest.NewServer(Router)
	} else if handlerFunc.HttpHandle != nil {
		ServerTest = httptest.NewServer(handlerFunc.HttpHandle())
		log.Print(ServerTest)
	}
	if ServerTest == nil {
		return
	}
	defer ServerTest.Close()
	if handlerFunc.Init != nil {
		handlerFunc.Init()
	}
}

// File2String change the file to a string
func File2String(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading ", filename, " error: ", err)
	}
	return string(data)
}

// APICall call the API
func APICall(requestType, route, headers, data string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(requestType, ServerTest.URL+route, strings.NewReader(data))
	lines := strings.Split(headers, "\n")
	for i := range lines {
		pos := strings.Index(lines[i], ": ")
		if pos == -1 {
			continue
		}
		hName := lines[i][0:pos]
		hValue := lines[i][pos+2:]
		if strings.Contains(hName, "Request") {
			continue
		}
		request.Header.Set(hName, hValue)
	}
	w := httptest.NewRecorder()
	log.Print(request)
	if Router == nil {
		HandlerFunc.HttpHandle().ServeHTTP(w, request)
	} else {
		Router.ServeHTTP(w, request)
	}
	return w
}

// Events test the event folder
func Events(t *testing.T, requestType, route, folder, fileName string) *httptest.ResponseRecorder {
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {
		if f.Name() == fileName {
			outmsg := File2String(folder + f.Name())
			outHeader := File2String(folder + strings.Replace(f.Name(), ".json", ".header", 1))
			w := APICall(requestType, route, outHeader, outmsg)
			return w
		}
	}
	return nil
}

func Query(t *testing.T, requestType, route string) *httptest.ResponseRecorder {
	w := APICall(requestType, route, "", "")
	return w
}
