package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"

	"github.com/gobwas/ws"
)

func WebSocketLowLevel(w http.ResponseWriter, r *http.Request) {
	
	conn, wr, _, err  := ws.UpgradeHTTP(r, w)
	if err != nil{
		panic(err)
	}
	bufReader := bufio.NewReader(wr.Reader)
	go func(){
		for {
			header, err := ws.ReadHeader(wr.Reader)
			fmt.Println("header",header)
			if err != nil{
				fmt.Println("header error:", err)
				return
			}
			payload := make([]byte, header.Length)
			
			bufReader.Read(payload)
			if err != nil{
				fmt.Println(err)
				return
			}
			ws.Cipher(payload,header.Mask,0)
			fmt.Println(string(payload))
			if header.OpCode == ws.OpClose{
				fmt.Println("hit")
				conn.Close()
				return
			}
		}
	}()

	
	bufwriter := bufio.NewWriter(conn)
	go func(){
		for {
			header := ws.Header{}
			header.Fin = true
			header.Rsv = 0
			header.OpCode = 1
			header.Masked = false
			msg := "writing this message!"
			msgbyte := []byte(msg)
			header.Length = 21
			err := ws.WriteHeader(conn, header)
			if err != nil{
				fmt.Println("header error:", err)
				return
			}
			// ws.Cipher( msgbyte, header.Mask, 0)
			
			_, err = bufwriter.Write(msgbyte)
			bufwriter.Flush()
			if err != nil{
				fmt.Println(err)
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()
	return
}