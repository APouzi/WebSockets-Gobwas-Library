package main

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
)

func WebSocketLowLevel(w http.ResponseWriter, r *http.Request) {
	// var conn net.Conn
	
	conn, wr, _, err  := ws.UpgradeHTTP(r, w)
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	bufReader := bufio.NewReaderSize(wr.Reader, 300)
	go func(){
		for {
			header, err := ws.ReadHeader(wr.Reader)
			if err != nil{
				fmt.Println("header error:", err)
				continue
			}
			payload := make([]byte, header.Length)
			
			bufReader.Read(payload)
			if err != nil{
				fmt.Println(err)
				continue
			}
			ws.Cipher(payload,header.Mask,0)
			fmt.Println(string(payload))
			
		}
	}()
}