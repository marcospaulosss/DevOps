package rpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"backend/apps/severino/src/interfaces"
	log "backend/libs/logger"
)

type RpcClient struct {
	address       string
	connection    *grpc.ClientConn
	ctx           context.Context
	userClient    interfaces.Client
	accountClient interfaces.Client
}

func NewRpcClient(address string) RpcClient {
	return RpcClient{
		address: address,
		ctx:     context.Background(),
	}
}

func (this RpcClient) Connect() *grpc.ClientConn {
	k := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}
	opts := grpc.WaitForReady(false)
	conn, err := grpc.Dial(
		this.address,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(k),
		grpc.WithDefaultCallOptions(opts),
	)
	if err != nil {
		log.Error("Falhei. Nao consegui conectar em "+this.address, err.Error())
	} else {
		log.Info("Conectando em", this.address)
	}
	this.connection = conn
	return conn
}

func (this RpcClient) Disconnect() {
	if this.connection != nil {
		log.Info("Desconectei de " + this.address)
		this.connection.Close()
	}
}
