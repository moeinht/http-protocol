package main

import (
	"fmt"
	"http-protocol/request"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	fmt.Println("server has been started on port :8000")

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	con, err := l.Accept()

	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()

	parser := request.NewParser(con)

	req, err := parser.Parse()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Method:", req.RequestLine.Method)
	fmt.Println("Path:", req.RequestLine.Path)
	fmt.Println("body:", string(req.Body))

	for _, h := range req.Headers {
		fmt.Println(h.Name, "=", h.Value)
	}
}
