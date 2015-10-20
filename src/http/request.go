package http

import (
	// "fmt"
	"bytes"
	"net/url"
	"strings"
)

type Headers map[string]string

type Request struct {
	URI     string
	Headers Headers
	Method  string
}

func (h Headers) ToString() string {
	var result bytes.Buffer

	for k, v := range h {
		result.WriteString(k + headerSep + v + stringSep)
	}

	return result.String()
}

func (h Headers) Add(key string, value string) {
	h[key] = value
}

func (r *Request) setHeaders(h Headers) {
	r.Headers = h
}

func (r *Request) parseRequest(s string) {
	parts := strings.Split(s, " ")
	URI := strings.Split(parts[1], "?")
	cleanedURI, _ := url.QueryUnescape(URI[0])

	r.URI = cleanedURI
	r.Method = parts[0]
}

func getHeaders(data []string) Headers {
	header := make([]string, 2)
	headers := Headers{}

	for line := range data {
		header = strings.Split(data[line], headerSep)
		if len(header) > 1 {
			headers[header[0]] = header[1]
		}
	}

	return headers
}

func ParseRequestString(s string) *Request {
	splittedRequestString := strings.Split(s, stringSep)
	request := new(Request)

	request.parseRequest(splittedRequestString[0])
	request.setHeaders(getHeaders(splittedRequestString[1:]))
	return request
}
