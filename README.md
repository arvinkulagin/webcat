# webcat
Webcat is a tool for reading from and writing to websocket. Currently supports text messages only.

Dependencies:
github.com/gorilla/websocket
github.com/arvinkulagin/cli

Websocket echo server:
`webcat server localhost:8888 /chat`

Websocket client:
`webcat ws://localhost:8888/chat`