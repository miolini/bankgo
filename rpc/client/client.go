package client

import (
	"github.com/miolini/bankgo/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type BalanceRecord struct {
	UserID int64
	Value  int64
}

type BalanceStorageClient struct {
	conn   *grpc.ClientConn
	client proto.BalanceStorageClient
}

func Connect(addr string) (*BalanceStorageClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	bsc := &BalanceStorageClient{}
	bsc.conn = conn
	bsc.client = proto.NewBalanceStorageClient(bsc.conn)
	return bsc, nil
}

func (bsc *BalanceStorageClient) Close() {
	bsc.conn.Close()
}

func (bsc *BalanceStorageClient) GetBalance(userID int64) (int64, error) {
	request := proto.GetRequest{UserId: userID}
	response, err := bsc.client.Get(context.Background(), &request)
	if err != nil {
		return 0, nil
	}
	return response.Value, nil
}

func (bsc *BalanceStorageClient) SetValue(userID int64, value int64) (int64, error) {
	request := proto.SetRequest{UserId: userID, Value: value}
	response, err := bsc.client.Set(context.Background(), &request)
	if err != nil {
		return 0, err
	}
	return response.Value, nil
}

func (bsc *BalanceStorageClient) IncrementValue(userID int64, amount int64) (int64, error) {
	request := proto.IncrementRequest{UserId: userID, Amount: amount}
	response, err := bsc.client.Increment(context.Background(), &request)
	if err != nil {
		return 0, err
	}
	return response.Value, nil
}

func (bsc *BalanceStorageClient) AllBalances() ([]BalanceRecord, error) {
	response, err := bsc.client.All(context.Background(), &proto.AllRequest{})
	if err != nil {
		return nil, err
	}
	result := make([]BalanceRecord, len(response.Records))
	for i := range response.Records {
		result[i] = BalanceRecord{response.Records[i].UserId, response.Records[i].Value}
	}
	return result, nil
}
