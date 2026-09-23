package server

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/moeinht/http-protocol/request"
	"github.com/moeinht/http-protocol/response"
	"github.com/moeinht/http-protocol/router"
)

func Listener(port int) (net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, err
	}

	return l, nil
}

func Server(l net.Listener, r *router.Router) error {
	defer l.Close()

	for {
		con, err := l.Accept()

		if err != nil {
			return err
		}

		go handleConnetion(con, r)
	}
}

func handleConnetion(con net.Conn, r *router.Router) {
	defer con.Close()

	parser := request.NewParser(con)

	for {
		req, err := parser.Parse()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}

			fmt.Println(err)

			res := response.NewResponse(400)
			if err := writer(con, *res); err != nil {
				fmt.Println(err)
			}

			return
		}

		res, err := r.Handler(req)
		if err != nil {
			errorResponse := response.NewResponse(500)
			if err := writer(con, *errorResponse); err != nil {
				fmt.Println(err)
			}
			return
		}

		if err := writer(con, res); err != nil {
			fmt.Println(err)
			return
		}

		if connection, exist := req.GetHeader("Connection"); connection == "close" && exist {
			return
		}
	}
}

func writer(con net.Conn, res response.Response) error {
	serializedResponse, err := res.Serialize()
	if err != nil {
		return err
	}

	_, err = con.Write(serializedResponse)
	return err
}
