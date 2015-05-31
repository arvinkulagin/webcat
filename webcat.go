package main

import (
	"github.com/arvinkulagin/cli"
	"github.com/arvinkulagin/webcat/handlers"
)

func main() {
	cli.Add("{url}", handlers.Client)
	cli.Add("server {addr} {path}", handlers.Server)
	cli.Run()
}
