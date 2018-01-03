# Http

Library http provides function to help write http communication protocol in GO.

## How To Use
Make a simple http request
```go
 	queries := make(map[string]string)
	queries["url"] = "test"
	data := make(map[string]string)
	data["name"] = "Maestro"
	header := make(map[string]string)
	header["test"] = "testHEADER"
	http.POST("http://localhost", queries, header, data)
```

Decode http Request
```go
resp := NexHttp.POST(url, query, header, payload)
jsonResp := http.DecodeMapStringInterface(resp.Body)
```