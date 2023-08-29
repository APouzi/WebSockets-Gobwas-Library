package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func WebSocketStreamToClient(writer http.ResponseWriter, request *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(request, writer)
	if err != nil {
		fmt.Println(err)
	}
	go func(conn net.Conn) {
		defer conn.Close()
		for {
			msg, op, err := wsutil.ReadClientData(conn)
			fmt.Println("op", op)
			if err != nil {
				fmt.Println(err)
				fmt.Println("ReadClient error")
				return
			}
			if string(msg) == "end" {
				return
			}
		}
	}(conn)

	go func(conn net.Conn) {
		defer conn.Close()
		for {
			msg := []byte("hello over and over")
			err = wsutil.WriteServerMessage(conn, ws.OpCode(1), msg)
			if err != nil {
				fmt.Println("WriteServer error")
				return
			}
			time.Sleep(time.Second * 1)
		}
	}(conn)

	// writer.Write([]byte("hello"))
}