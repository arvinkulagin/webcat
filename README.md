# webcat
Webcat is a tool for reading from and writing to websocket. Currently supports text messages only.

### Dependencies
* github.com/gorilla/websocket
* github.com/arvinkulagin/cli

### Getting webcat
```
go get github.com/arvinkulagin/webcat
cd go/src/github.com/arvinkulagin/webcat
go install
```

### Websocket echo server
`webcat server localhost:8888 /echo`

### Websocket client
`webcat ws://localhost:8888/echo`