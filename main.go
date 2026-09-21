package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal(err)
	}

	con, err := l.Accept()

	if err != nil {
		log.Fatal(err)
	}

	defer con.Close()

	fmt.Println("server has been started on port :8000")
}
