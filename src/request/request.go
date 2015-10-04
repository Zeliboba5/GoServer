package request

import (
	"strings"
	"net/url"
)

const (
	sep := ": "
)

type Header map[string]string

type Request struct {
	URL *url.URL

}

func getHeaders(data []string) Header {
	header := []string{}
	headers := Header{}

	for line := range data {
		header = strings.Split(data, sep)
	}
}