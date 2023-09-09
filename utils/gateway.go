package utils

import (
	"github.com/hiro942/elden-app/global"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"google.golang.org/grpc"
	"time"
)

func NewGateway(clientConnection *grpc.ClientConn) (*client.Gateway, error) {
	ID := NewIDentity(global.CertPath, global.MspID)
	sign := NewSign(global.KeyPath)
	return client.Connect(
		ID,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
}
