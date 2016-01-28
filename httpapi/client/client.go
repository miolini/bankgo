package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

type GetBalanceResponse struct {
	Error    string          `json:"error"`
	Response []BalanceRecord `json:"response"`
}

type PostTransactionResponse struct {
	Error    string `json:"error"`
	Response struct {
		UserID int64
		Value  int64
	} `json:"response"`
}

type PostTransactionRequest struct {
	UserID int64
	Value  int64
}

type BalanceRecord struct {
	UserID int64
	Value  int64
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
	return c
}

func (c *HttpApiClient) GetBalances() ([]BalanceRecord, error) {
	var response GetBalanceResponse
	err := c.request("GET", "/balances", &response, nil)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf("response err: %s", response.Error)
	}
	return response.Response, err
}

func (c *HttpApiClient) PostTransaction(userID int64, amount int64) (int64, error) {
	request := PostTransactionRequest{UserID: userID, Value: amount}
	var response PostTransactionResponse
	err := c.request("POST", "/transaction", &response, &request)
	if err != nil {
		return 0, err
	}
	if response.Error != "" {
		return 0, fmt.Errorf("response err: %s", response.Error)
	}
	return response.Response.Value, nil
}

func (c *HttpApiClient) request(method, reqUrl string, jsonResponse interface{}, jsonRequest interface{}) error {
	var bodyReader *bytes.Buffer
	if jsonRequest != nil {
		data, err := json.Marshal(jsonRequest)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(data)
	} else {
		bodyReader = &bytes.Buffer{}
	}
	reqUrl = fmt.Sprintf("%s%s", c.baseUrl, reqUrl)
	log.Printf("request %s %s", method, reqUrl)
	req, err := http.NewRequest(method, reqUrl, bodyReader)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	responseData := bytes.Buffer{}
	_, err = io.Copy(&responseData, res.Body)
	if err != nil {
		return err
	}
	log.Printf("http response: %s", responseData.Bytes())
	return json.NewDecoder(&responseData).Decode(jsonResponse)
}
