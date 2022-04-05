package service

import (
	"encoding/json"
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hiro942/elden-app/model"
	"github.com/pkg/errors"
)

// GetAllNodes evaluates a transaction to read ledger state (read all nodes)
func GetAllNodes() ([]model.Node, error) {
	fmt.Println("Evaluate Transaction: GetAllNodes, returns all current nodes on the ledger.")

	nodesBytes, err := global.Contract.EvaluateTransaction("GetAllNodes")
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate transaction")
	}

	var nodes []model.Node
	_ = json.Unmarshal(nodesBytes, &nodes)

	return nodes, nil
}

// GetNodeById evaluates a transaction to read node by id
func GetNodeById(id string) (model.Node, error) {
	fmt.Println("Evaluate Transaction: GetNodeById, returns specific node by node-id.")

	nodeBytes, err := global.Contract.EvaluateTransaction("GetNodeById", id)
	if err != nil {
		return model.Node{}, errors.Wrap(err, "failed to evaluate transaction")
	}

	var node model.Node
	_ = json.Unmarshal(nodeBytes, &node)

	return node, nil
}

// GetSatellitePublicKey evaluates a transaction to read public key of satellite by id
func GetSatellitePublicKey(id string) (string, error) {
	fmt.Println("Evaluate Transaction: GetSatellitePublicKey, returns the satellite's public key.")

	publicKeyBytes, err := global.Contract.EvaluateTransaction("GetSatellitePublicKey", id)
	if err != nil {
		return "", errors.Wrap(err, "failed to evaluate transaction")
	}

	publicKey := string(publicKeyBytes)
	return publicKey, nil
}

// GetUserPublicKey evaluates a transaction to read public key of user by id and mac address
func GetUserPublicKey(id string, macAddr string) (string, error) {
	fmt.Println("Evaluate Transaction: GetUserPublicKey, returns the satellite's public key.")

	publicKeyBytes, err := global.Contract.EvaluateTransaction("GetUserPublicKey", id, macAddr)
	if err != nil {
		return "", errors.Wrap(err, "failed to evaluate transaction")
	}

	publicKey := string(publicKeyBytes)
	return publicKey, nil
}
