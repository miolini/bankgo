package client

import (
	"github.com/miolini/bankgo/rpc/server/common"
	"log"
	"testing"
	"time"
)

func TestClientGet(t *testing.T) {
	var client *BalanceStorageClient

	addr := "127.0.0.1:17090"
	server, err := common.NewServer(addr)
	if err != nil {
		log.Fatal(err)
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

	value, err := client.IncrementValue(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	value, err = client.IncrementValue(1, -20)
	if err != nil {
		t.Fatal(err)
	}
	if value != 90 {
		t.Fatalf("balance not equals to 90: %d", value)
	}
	t.Logf("balance tested")
}
