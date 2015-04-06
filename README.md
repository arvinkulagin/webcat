# webcat
Webcat is a tool for reading from and writing to websocket. Currently supports text messages only.

Dependencies:
https://github.com/gorilla/websocket

`webcat -s ws://localhost:8888 -o http://localhost:8888`

-s Specify websocket URL

-o Specify Origin HTTP header