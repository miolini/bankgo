package main

import (
	"flag"
	"log"

	"github.com/miolini/bankgo/httpapi/common"
)

var (
	flAddr = flag.String("l", "127.0.0.1:14080", "http api listen addr:port")
)

func main() {
	var err error
	flag.Parse()
	app := common.App{}
	if err = app.Init(); err != nil {
		log.Fatalf("init error: %s", err)
	}
	app.Run(*flAddr)
}