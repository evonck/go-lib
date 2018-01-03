package tools

//Package tool provide common tool function usefull for go.

import (
	"bytes"
	"context"
	"encoding/gob"
	stdhttp "net/http"
	"os"
	"strconv"

	"github.com/go-kit/kit/transport/http"
	"github.com/satori/go.uuid"
)

// IsInArray check if a value is in a string array
func IsInArray(name string, array []string) bool {
	for _, arrayName := range array {
		if arrayName == name {
			return true
		}
	}
	return false
}

// CheckStringParam check if parameters are nil
func CheckStringParam(params ...string) bool {
	for _, param := range params {
		if param == "" {
			return false
		}
	}
	return true
}

// GetBytes change an interface into []byte
func GetBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CheckSlash check if an url as a slash at the end
func CheckSlash(url string) string {
	if url[len(url)-1:] != "/" {
		url += "/"
	}
	return url
}

// Exists check if a directory or file exist
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ParseString Parse string int64 to int64
func ParseString(timestamp string) int64 {
	dateInt64, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return 0
	}
	return dateInt64
}

// WriteError helper function
func WriteError(w stdhttp.ResponseWriter, errorName string, StatusCode int) {
	w.WriteHeader(StatusCode)
	w.Write([]byte(errorName))
}

// WriteErrorBytes helper function
func WriteErrorBytes(w stdhttp.ResponseWriter, errorName []byte, StatusCode int) {
	w.WriteHeader(StatusCode)
	w.Write(errorName)
}

// AddHeader Add header to the request
func AddHeader(w stdhttp.ResponseWriter) {
	h := w.Header()
	h.Add("Access-Control-Allow-Origin", "*")
	h.Add("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	h.Add("Access-Control-Allow-Headers", "Content-Type")
}

// AddHeaderCorrelationID Add Correlation ID to the header of the request
func AddHeaderCorrelationID() http.RequestFunc {
	return func(ctx context.Context, r *stdhttp.Request) context.Context {
		// type contextKey string
		// k := contextKey("X-Correlation-Id")
		uid := uuid.NewV4()
		r.Header.Add("X-Correlation-Id", uid.String())
		ctx = context.WithValue(context.Background(), "X-Correlation-Id", uid.String())
		return ctx
	}
}
