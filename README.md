# webcat
Webcat is a networking service for reading from and writing to network connections using websocket.
It works like netcat, but over websocket.

App uses https://github.com/gorilla/websocket

Flags:
-s Specify websocket URL
-o Specify Origin HTTP handler

webcat -s ws://localhost:8888 -o http://localhost:8888