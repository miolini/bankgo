package client

import (
	"log"
	"testing"

	"fmt"
	httpserver "github.com/miolini/bankgo/httpapi/server/common"
	rpcserver "github.com/miolini/bankgo/rpc/server/common"
	"time"
)

const (
	HttpAddr = "localhost:16180"
	RpcAddr  = "localhost:16190"
)

func TestMain(m *testing.M) {
	server, err := rpcserver.NewServer(RpcAddr)
	if err != nil {
		log.Fatalf("rpc init err: %s", err)
	}
	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("rpc run err: %s", err)
		}
	}()
	log.Printf("rpc runned")
	app := httpserver.App{}
	if err = app.Init(RpcAddr); err != nil {
		log.Fatalf("http init error: %s", err)
	}
	go func() {
		if err := app.Run(HttpAddr); err != nil {
			log.Fatalf("http run error: %s", err)
		}
	}()
	log.Printf("http runned")
	time.Sleep(time.Second)
	m.Run()
}

func TestHttpApi(t *testing.T) {
	t.Logf("test http client")
	c := NewHttpApiClient(fmt.Sprintf("http://%s", HttpAddr))
	r1, err := c.PostTransaction(1, 10)
	if err != nil {
		t.Fatal(err)
	} else if r1 != 110 {
		t.Fatalf("balance should be 110, received %d", r1)
	}
	r2, err := c.PostTransaction(1, -20)
	if err != nil {
		t.Fatal(err)
	} else if r2 != 100+10-20 {
		t.Fatalf("balance should be 90, received %d", r2)
	}
	r3, err := c.GetBalances()
	if err != nil {
		t.Fatal(err)
	}
	for _, rec := range r3 {
		if rec.UserID == 1 {
			if rec.Value != 100+10-20 {
				log.Fatalf("bad user balance: %d", rec.Value)
			}
			return
		}
	}
	t.Fatal("user not found in all balances")
}
