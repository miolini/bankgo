package common

import (
	"errors"
	"log"
	"net"
	"sync"

	"github.com/miolini/bankgo/rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	StartBalance = 100
)

var (
	ErrNotFound = errors.New("UserID not found")
	ErrUserIDBadValue = errors.New("UserID must be greater than 0")
	ErrDecreaseNegativeBalance = errors.New("Descrease negative balance")
)

type BalanceStorageServer struct {
	proto.BalanceStorageServer
	dataShards      []*dataShard
	dataShardsCount int64
	listener        net.Listener
	grpcServer      *grpc.Server
}

type dataShard struct {
	sync.RWMutex
	data map[int64]int64
}

func newDataShard() *dataShard {
	return &dataShard{
		data: make(map[int64]int64),
	}
}

func NewServer(addr string) (*BalanceStorageServer, error) {
	var i int64
	bss := new(BalanceStorageServer)
	bss.dataShardsCount = 1024
	bss.dataShards = make([]*dataShard, bss.dataShardsCount)
	for i = 0; i < bss.dataShardsCount; i++ {
		bss.dataShards[i] = newDataShard()
	}
	var err error
	log.Printf("run rpc on %s", addr)
	if bss.listener, err = net.Listen("tcp", addr); err != nil {
		return nil, err
	}
	bss.grpcServer = grpc.NewServer()
	proto.RegisterBalanceStorageServer(bss.grpcServer, bss)
	return bss, nil
}

func (bss *BalanceStorageServer) Run() error {
	return bss.grpcServer.Serve(bss.listener)
}

func (bss *BalanceStorageServer) Close() {
	bss.listener.Close()
}

func (bss *BalanceStorageServer) getShard(userID int64) *dataShard {
	return bss.dataShards[int(userID % bss.dataShardsCount)]
}

func (bss *BalanceStorageServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.BalanceResponse, error) {
	if request.UserId < 1 {
		return nil, ErrUserIDBadValue
	}
	shard := bss.getShard(request.UserId)
	shard.RLock()
	value, ok := shard.data[request.UserId]
	shard.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}
	return &proto.BalanceResponse{Value: value}, nil
}

func (bss *BalanceStorageServer) Increment(ctx context.Context, request *proto.IncrementRequest) (*proto.BalanceResponse, error) {
	var err error
	if request.UserId < 1 {
		return nil, ErrUserIDBadValue
	}
	shard := bss.getShard(request.UserId)
	shard.Lock()
	value, ok := shard.data[request.UserId]
	if value < 0 && request.Amount < 0 {
		err = ErrDecreaseNegativeBalance
	} else {
		if !ok {
			value = StartBalance
		}
		value += request.Amount
		shard.data[request.UserId] = value
	}
	shard.Unlock()
	if err != nil {
		return nil, err
	}
	return &proto.BalanceResponse{Value: value}, nil
}

func (bss *BalanceStorageServer) Set(ctx context.Context, request *proto.SetRequest) (*proto.BalanceResponse, error) {
	if request.UserId < 1 {
		return nil, ErrUserIDBadValue
	}
	shard := bss.getShard(request.UserId)
	shard.Lock()
	shard.data[request.UserId] = request.Value
	shard.Unlock()
	return &proto.BalanceResponse{Value: request.Value}, nil
}

func (bss *BalanceStorageServer) All(ctx context.Context, request *proto.AllRequest) (*proto.AllResponse, error) {
	var results []*proto.BalanceRecord
	for i := range bss.dataShards {
		shard := bss.dataShards[i]
		shard.RLock()
		for userID, value := range shard.data {
			results = append(results, &proto.BalanceRecord{userID, value})
		}
		shard.RUnlock()
	}
	return &proto.AllResponse{Records: results}, nil
}
