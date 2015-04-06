package main

import (
	"bufio"
	"log"
	"os"
	"flag"
	"strings"
	"net/http"
	"github.com/gorilla/websocket"
)

func main() {
	url := flag.String("s", "ws://localhost:8888", "URL")
	origin := flag.String("o", "", "Origin")
	flag.Parse()

	header := http.Header{}
	header.Add("Origin", *origin)

	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(*url, header)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// Write
		for {
			msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			msg = strings.Trim(msg, "\n")
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println(err)
			}
		}
	}()

	for {
		// Read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		println(string(msg))
	}
}