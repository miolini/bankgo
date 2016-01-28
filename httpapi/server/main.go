package main

import (
	"flag"
	"log"

	"github.com/miolini/bankgo/httpapi/server/common"
)

var (
	flAddr    = flag.String("l", ":14080", "http api listen addr:port")
	flRpcAddr = flag.String("rpc", "rpc:14090", "rpc api listen addr:port")
)

func main() {
	var err error
	flag.Parse()
	app := common.App{}
	if err = app.Init(*flRpcAddr); err != nil {
		log.Fatalf("init error: %s", err)
	}
	if err = app.Run(*flAddr); err != nil {
		log.Fatalf("run error: %s", err)
	}
}
