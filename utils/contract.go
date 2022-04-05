package utils

import (
	"github.com/hiro942/elden-app/global"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func GetContract(gateway *client.Gateway) *client.Contract {
	network := gateway.GetNetwork(global.ChannelName)
	contract := network.GetContract(global.ChaincodeName)
	return contract
}
