package global

import "github.com/hyperledger/fabric-gateway/pkg/client"

// backup
//const (
//	MspID         = "Org1MSP"
//	CryptoPath    = "/root/go/src/github.com/hyperledger/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com"
//	CertPath      = CryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
//	KeyPath       = CryptoPath + "/users/User1@org1.example.com/msp/keystore/"
//	TlsCertPath   = CryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
//	PeerEndpoint  = "localhost:7051"
//	GatewayPeer   = "peer0.org1.example.com"
//	ChannelName   = "mychannel"
//	ChaincodeName = "basic"
//)

const (
	MspID         = "Org1MSP"
	CryptoPath    = "/home/niclab/zyq/distribute/fabric/go/fabric-samples/dev/organizations/peerOrganizations/org1.example.com"
	CertPath      = CryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
	KeyPath       = CryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	TlsCertPath   = CryptoPath + "/peers/peer0.org1.example.com/tls-msp/tlscacerts/tls-0-0-0-0-7052-tls-ca-node1.pem"
	PeerEndpoint  = "localhost:7152"
	GatewayPeer   = "peer1.org1.example.com"
	ChannelName   = "mychannel"
	ChaincodeName = "testFF"
)

var Contract *client.Contract
