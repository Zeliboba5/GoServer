package http

import (
	"bytes"
	"io/ioutil"
	// "loger"
	"status"
	"strconv"
	"time"
)

type Response struct {
	Status  string
	Body    string
	Headers Headers
}

func (r *Response) addDefaultHeaders() {
	r.Headers.Add("Server", serverName)
	r.Headers.Add("Date", time.Now().Format(dateTimeFormat))
	r.Headers.Add("Connection", "close")
}

func (r Response) Bytes() []byte {
	var result bytes.Buffer

	result.WriteString(HTTP_PROTO + " " + r.Status + stringSep)
	result.WriteString(r.Headers.ToString() + stringSep)
	result.WriteString(r.Body)

	return result.Bytes()
}

func checkForErrors(response *Response, method string, path string) bool {
	if !isSecurePath(path) {
		response.Status = status.FORBIDDEN
		return true
	} else if !isCorrectMethod(method) {
		response.Status = status.ERROR
		return true
	} else if !isSupportedMethod(method) {
		response.Status = status.NOT_ALLOWED
		return true
	}

	return false
}

func GenerateResponse(method string, path string) Response {
	response := Response{}
	response.Headers = Headers{}
	response.addDefaultHeaders()

	if checkForErrors(&response, method, path) {
		return response
	}

	directoryFlag := isDirectory(path)

	if directoryFlag {
		path += defaultfile
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		if directoryFlag {
			response.Status = status.FORBIDDEN
		} else {
			response.Status = status.NOT_FOUND
		}
		return response
	}

	if method == "GET" {
		response.Body = string(data)
	}

	response.Status = status.OK

	response.Headers.Add("Content-Type", getContentTypeByFilename(path))
	response.Headers.Add("Content-Length", strconv.Itoa(len(data)))

	return response
}
