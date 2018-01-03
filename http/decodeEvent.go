package http

import (
	"encoding/json"
	"io"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// DecodeEventInterface decode a response using interface
func DecodeEventInterface(resp io.ReadCloser) interface{} {
	requestData, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Warnf("An error occure while decoding the json Information %s", err)
		return nil
	}
	var jsonResp interface{}
	if err := json.Unmarshal(requestData, &jsonResp); err != nil {
		log.Warnf("An error occure while decoding the json Information %s", err)
		return nil
	}
	return jsonResp
}

// DecodeMapStringInterface decode a response using map[string]interface
func DecodeMapStringInterface(resp io.ReadCloser) map[string]interface{} {
	requestData, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Warnf("An error occure while decoding the json Information %s", err)
		return nil
	}
	jsonResp := make(map[string]interface{})
	if err := json.Unmarshal(requestData, &jsonResp); err != nil {
		log.Warnf("An error occure while decoding the json Information %s", err)
		return nil
	}
	return jsonResp
}

// DecodeEventArrayInterface decode a response using []interface
func DecodeEventArrayInterface(resp io.ReadCloser) []interface{} {
	requestData, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Warnf("An error occure while decoding the json Information %s", err)
		return nil
	}
	var jsonResp []interface{}
	if err := json.Unmarshal(requestData, &jsonResp); err != nil {
		log.Warnf("An error occure while decoding the json Information %s", err)
		return nil
	}
	return jsonResp
}
