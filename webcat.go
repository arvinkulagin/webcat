package main

import (
	"github.com/gorilla/websocket"
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"
	"net/http"
)

var origin string

func init() {
	flag.StringVar(&origin, "o", "", "Origin header")
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		fmt.Println("You must specify websocket url")
		return
	}
	url := flag.Args()[0]
	header := http.Header{}
	header.Add("Origin", origin)
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(url, header)
	if err != nil {
		fmt.Printf("Error: Can't connect to %s\n", url)
		return
	}
	go func() {
		// Write loop
		for {
			msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			msg = strings.Trim(msg, "\n")
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				return
			}
		}
	}()

	// Read loop
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		fmt.Println(string(msg))
	}
}