package main

import (
	"flag"
	"fmt"
	"http"
	"loger"
	"net"
	"runtime"
)

var (
	root   string
	numCPU int
)

func main() {
	loger.D("STATE", "STARTED")

	flag.IntVar(&numCPU, "c", runtime.NumCPU(), "set max cpu")
	flag.Parse()

	flag.StringVar(&root, "r", "static", "set static root directory")
	flag.Parse()

	runtime.GOMAXPROCS(numCPU)
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
		return
	}
	HttpRequest := string(buf)
	requestObj := http.ParseRequestString(HttpRequest)

	response := http.GenerateResponse(requestObj.Method, root+requestObj.URI)
	conn.Write(response.Bytes())
}
