package loger

import (
	"fmt"
)

const (
	debug     = "LOG::DEBUG: "
	loggerSep = " -> "
)

func D(tag string, text string) {
	fmt.Println(debug + tag + loggerSep + text)
}
