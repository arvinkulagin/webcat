# webcat
Webcat is a tool for reading from and writing to websocket. Currently supports text messages only.

### Dependencies
github.com/gorilla/websocket
github.com/arvinkulagin/cli

### Getting webcat
`go get github.com/gorilla/websocket`
`go get github.com/arvinkulagin/cli`
`go get github.com/arvinkulagin/webcat`

### Websocket echo server
`webcat server localhost:8888 /echo`

### Websocket client
`webcat ws://localhost:8888/echo`