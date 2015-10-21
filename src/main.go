package main

import (
	"fmt"
	"http"
	"loger"
	"net"
)

var (
	root = "static"
)

func main() {
	loger.D("STATE", "STARTED")
	ln, er := net.Listen("tcp", ":80")
	if er != nil {
		fmt.Printf("error in listen")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("error in go")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("error in handle")
	}
	HttpRequest := string(buf)
	requestObj := http.ParseRequestString(HttpRequest)

	response := http.GenerateResponse(requestObj.Method, root+requestObj.URI)
	conn.Write(response.Bytes())
}
