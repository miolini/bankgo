package common

import (
	"log"
	"testing"
	"github.com/miolini/bankgo/rpc/server/common"
	"net/http/httptest"
)

const (
	servAddr = "localhost:15090"
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
	httptest.NewServer()
}
