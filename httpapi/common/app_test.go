package common

import (
	"log"
	"testing"

	"github.com/miolini/bankgo/rpc/server/common"
)

const (
	servAddr = "localhost:14090"
)

func TestMain(m *testing.M) {
	server, err := common.NewServer(servAddr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := server.Run(); err != nil {
			log.Fatal(err)
		}
	}()
}

func TestApiGetBalances(t *testing.T) {
	
}
