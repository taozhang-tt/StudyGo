package main

import (
	"StudyGo/socket/impl"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//握手的过程中允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsHandler(w http.ResponseWriter, r *http.Request)  {
	vars := r.URL.Query()
	a, ok := vars["key"]
	if !ok {
		fmt.Printf("param a does not exist\n")
	} else {
		fmt.Printf(a[0])
	}
	var (
		wsConn *websocket.Conn
		err error
		conn *impl.Connection
		data []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

	ERR:
		conn.Close()
}

func ConnHandler(w http.ResponseWriter, r *http.Request)  {
	vars := r.URL.Query()
	a, ok := vars["key"]
	if !ok {
		fmt.Printf("param a does not exist\n")
	} else {
		fmt.Printf(a[0])
	}
}
func main()  {
	//http://localhost:7777/ws
	http.HandleFunc("/ws", WsHandler)
	http.HandleFunc("/ping", ConnHandler)

	http.ListenAndServe("127.0.0.1:7777", nil)
}
