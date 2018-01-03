//Package http provides common function use to communicate threw http
package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"bitbucket.org/onekloud/go-lib/models"
	log "github.com/Sirupsen/logrus"
)

// callerWithData make an http request with payload data
func callerWithData(requestType, url string, queries map[string]string, headers map[string]string, payload interface{}) *http.Response {
	urlQuery := addQuery(url, queries)
	payloadJSON := GetPayload(payload)
	req, _ := http.NewRequest(requestType, urlQuery, bytes.NewBuffer(payloadJSON))
	req = addHeader(req, headers)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warn("An error occure while sending the json Information ", err)
		return nil
	}
	return resp
}

func callerWithDataText(requestType, url string, queries map[string]string, headers map[string]string, payload []byte) *http.Response {
	urlQuery := addQuery(url, queries)
	req, _ := http.NewRequest(requestType, urlQuery, bytes.NewBuffer(payload))
	req = addHeader(req, headers)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warn("An error occure while sending the json Information ", err)
		return nil
	}
	return resp
}

// GetPayload return the Json payload already passed Marshal function
func GetPayload(payload interface{}) []byte {

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil
	}
	return payloadJSON
}

// POSTText make a POST request
func POSTText(url string, queries, headers map[string]string, payload []byte) *http.Response {
	return callerWithDataText("POST", url, queries, headers, payload)
}

// POST make a POST request
func POST(url string, queries, headers map[string]string, payload interface{}) *http.Response {
	return callerWithData("POST", url, queries, headers, payload)
}

// PUT make a PUT request
func PUT(url string, queries map[string]string, headers map[string]string, payload interface{}) *http.Response {
	return callerWithData("PUT", url, queries, headers, payload)
}

// GET make a GET request
func GET(url string, queries map[string]string, headers map[string]string) *http.Response {
	return callerWithData("GET", url, queries, headers, nil)
}

// DELETE make a GET request
func DELETE(url string, queries map[string]string, headers map[string]string) *http.Response {
	return callerWithData("DELETE", url, queries, headers, nil)
}

// addQuery add query to the url
func addQuery(urlValue string, queries map[string]string) string {
	queryNumber := 0
	for query, value := range queries {
		queryNumber++
		if queryNumber == 1 {
			urlValue += "?"
		} else {
			urlValue += "&"
		}
		urlValue += query + "=" + url.QueryEscape(value)
	}
	return urlValue
}

// addHeader add header to the request
func addHeader(req *http.Request, headers map[string]string) *http.Request {
	for header, value := range headers {
		req.Header.Set(header, value)
	}
	return req
}

// GetBasicAuth return an encoded username:password string
func GetBasicAuth(username, password string) string {
	authorizationString := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(authorizationString))

}

// AsyncHTTPGets allow to do multiple synch http get call
func AsyncHTTPGets(urls []string, headers map[string]string) []*models.SyncHTTPResponse {
	ch := make(chan *models.SyncHTTPResponse)
	responses := []*models.SyncHTTPResponse{}
	for _, url := range urls {
		go func(url string) {
			log.Printf("Fetching %s", url)
			resp := GET(url, nil, headers)
			ch <- &models.SyncHTTPResponse{URL: url, Response: resp}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Println() // empty on purpose
			log.Printf("%s was fetched", r.URL)
			if r.Response == nil {
				log.Error("with an error", r.Response.Status)
			}
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(25 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}
