package common

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handlers handle fonction of the program
type Handlers struct {
	Handle     func() *httprouter.Router
	HttpHandle func() http.Handler
	Init       func()
}

type MyServer struct {
    r http.Handler
}
