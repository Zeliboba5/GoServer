package http

import (
	// "fmt"
	"net/url"
	"strings"
)

const (
	headerSep = ": "
	stringSep = "\n"
)

type Headers map[string]string

type Request struct {
	uri     string
	Headers Headers
	Method  string
}

func (r *Request) setHeaders(h Headers) {
	r.Headers = h
}

func (r *Request) parseRequest(s string) {
	parts := strings.Split(s, " ")
	uri := strings.Split(parts[1], "?")
	cleanedURI, _ := url.QueryUnescape(uri[0])

	r.uri = cleanedURI
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
