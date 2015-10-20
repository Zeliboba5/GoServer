package http

import (
	// "fmt"
	// "net/url"
	"strings"
)

const (
	defaultfile    = "index.html"
	dateTimeFormat = "Mon, _1 Jan 1970 12:01:23 GMT"
	serverName     = "Tornago"
	headerSep      = ": "
	stringSep      = "\r\n"
	HTTP_PROTO     = "HTTP/1.1"
)

var (
	supportedMethods = []string{"GET", "HEAD"}
	httpMethods      = []string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "PATCH",
		"DELETE", "TRACE", "CONNECT"}
)

var exts = map[string]string{
	"txt":  "application/text",
	"html": "text/html",
	"json": "application/json",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"js":   "text/javascript",
	"css":  "text/css",
	"gif":  "image/gif",
	"swf":  "application/x-shockwave-flash",
}

func getContentTypeByFilename(fileName string) string {
	parts := strings.Split(fileName, ".")
	return exts[parts[len(parts)-1]]
}

func isDirectory(fileName string) bool {
	return fileName[len(fileName)-1:] == "/"
}

func isCorrectMethod(method string) bool {
	for _, v := range httpMethods {
		if method == v {
			return true
		}
	}

	return false
}

func isSupportedMethod(method string) bool {
	for _, v := range supportedMethods {
		if method == v {
			return true
		}
	}

	return false
}
