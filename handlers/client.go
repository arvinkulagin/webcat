package handlers

import (
	"net/http"
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/arvinkulagin/cli"
	"github.com/gorilla/websocket"
)

func Client(r cli.Request) {
	url := r.Vars()["url"]
	header := http.Header{}
	header.Add("Origin", "")

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