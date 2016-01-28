package client

import (
	"net/http"
	"time"
	"fmt"
	"encoding/json"
	"bytes"
)

const (
	MaxIdleConnections int = 20
	RequestTimeout int = 5
)

type GetBalanceResponse []BalanceRecord

type BalanceRecord struct {
	UserID int64
	Value int64
}

type HttpApiClient struct {
	baseUrl    string
	httpClient *http.Client
}

func NewHttpApiClient(baseUrl string) *HttpApiClient {
	c := &HttpApiClient{}
	c.baseUrl = baseUrl
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
}

func (c *HttpApiClient) GetBalances(userID int64) ([]BalanceRecord, error) {
	var response GetBalanceResponse
	err := c.request("GET", "/balances", &response, nil)
	return response, err
}

func (c *HttpApiClient) request(method, reqUrl, jsonResponse interface{}, jsonRequest interface{}) error {
	var bodyReader *bytes.Buffer
	if jsonRequest != nil {
		data, err := json.Marshal(jsonRequest)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(data)
	}
	res, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseUrl, reqUrl), bodyReader)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(jsonResponse)
}