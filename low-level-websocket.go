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

	
			
		}
	}()
	return
}