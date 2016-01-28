package client

import (
	"testing"
)

const (
	HttpAddr = "http://localhost:14180"
	RpcAddr = "http://localhost:14190"
)

func TestMain(m *testing.M) {

}

func TestHttpApi(t *testing.T) {
	c := NewHttpApiClient("http://localhost:15080")
}