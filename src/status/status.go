package status

const (
	OK          string = "200 OK"
	NOT_FOUND   string = "404 NOT FOUND"
	ERROR       string = "500 INTERNAL SERVER ERROR"
	BAD_REQUEST string = "400 BAD REQUEST"
	FORBIDDEN   string = "403 FORBIDDEN"
	// NOT_IMPLEMENTED string = "501 NOT IMPLEMENTED"
	NOT_ALLOWED  string = "405 METHOD NOT ALLOWED"
	DEFAULT_FILE string = "/index.html"
	FILE_404     string = "/404.html"
)

// func GetStatusLine(status string) string {
// 	return HTTP_PROTO + " " + status + "\r\n"
// }
