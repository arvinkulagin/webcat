package handlers

import (
	"net/http"
	"log"
	"fmt"
	"github.com/arvinkulagin/cli"
	"github.com/gorilla/websocket"
)

func Server(r cli.Request) {
	addr := r.Vars()["addr"]
	path := r.Vars()["path"]

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Connect: %s\n", conn.RemoteAddr().String())

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
				log.Printf("Disconnect: %s\n", conn.RemoteAddr().String())
				return
			}
			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				conn.Close()
				log.Printf("Disconnect: %s\n", conn.RemoteAddr().String())
				return
			}
		}
	})

	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}