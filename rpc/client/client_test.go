package client

import (
	"github.com/miolini/bankgo/rpc/server/common"
	"log"
	"testing"
	"time"
)

var (
	client *BalanceStorageClient
)

func TestMain(m *testing.M) {
	addr := "127.0.0.1:14090"
	server, err := common.NewServer(addr)
	if err != nil {
		panic(err)
	}
	go func() {
		if err = server.Run(); err != nil {
			log.Fatalf("server run err: %s", err)
		}
	}()
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond * 100)
		client, err = Connect(addr)
		if err != nil {
			continue
		}
	}
	if err != nil {
		log.Fatalf("client connect err: %s", err)
	} else {
		log.Printf("connected")
	}
}

func TestClientGet(t *testing.T) {
	value, err := client.GetBalance(1)
	if err != nil {
		t.Fatal(err)
	}
	if value != 100 {
		t.Fatalf("balance for new user not equals 100: %d", value)
	}
}
