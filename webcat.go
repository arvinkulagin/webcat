package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"
	"strings"
	"net/http"
	"github.com/gorilla/websocket"
)

func main() {
	url := flag.String("s", "", "URL")
	origin := flag.String("o", "", "Origin")
	flag.Parse()

	if *url == "" {
		fmt.Println("Error: You must Specify URL with -s")
		return
	}

	header := http.Header{}
	header.Add("Origin", *origin)

	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(*url, header)
	if err != nil {
		fmt.Printf("Error: Can't connect to %s\n", *url)
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
