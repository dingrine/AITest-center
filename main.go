package main

import (
	"examCenter/internal/server"
	"flag"
)

func main() {
	flag.Parse()
	server.New().Run()
}
