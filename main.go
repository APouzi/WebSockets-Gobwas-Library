package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	fmt.Println("Starting GobWas Server")
	http.HandleFunc("/ws", WebSocket)
	http.HandleFunc("/ws-stream", WebSocketStreamToClient)
	// http.HandleFunc("/ws-low-level",WebSocket2)
	err :=  http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println("error is yeah", err)
		return
	}
	

}

func WebSocket (writer http.ResponseWriter, request *http.Request){
	conn, _, _, err := ws.UpgradeHTTP(request, writer)
	if err != nil{
		fmt.Println(err)
	}
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil{
			fmt.Println("ReadClient error")
			return
		}
		if string(msg) == "end"{
			conn.Close()
			return
		}

		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil{
			fmt.Println("WriteServer error")
			return
		}
	}
	// writer.Write([]byte("hello"))
}
