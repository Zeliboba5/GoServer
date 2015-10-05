package main

import (
	"fmt"
	"http"
	"net"
	"status"
)

func main() {
	fmt.Println("started")
	ln, er := net.Listen("tcp", ":8080")
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

	switch requestObj.Method {
	case "GET":
		fmt.Println("GET")
	case "HEAD":
		fmt.Println("HEAD")
	default:
		{
			fmt.Println(requestObj.Method)
		}
	}
	responseString := status.NOT_FOUND
	conn.Write([]byte(responseString))
}
