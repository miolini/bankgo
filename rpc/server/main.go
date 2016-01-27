package main

import (
	"flag"
	"log"

	"github.com/miolini/bankgo/rpc/server/common"
)

var (
	flAddr = flag.String("l", ":14090", "listen host:port")
)

func main() {
	flag.Parse()
	s, err := common.NewServer(*flAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("run rpc server")
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
