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
	err = json.Unmarshal(nodesBytes, &nodes)
	if err != nil {
		return []model.Node{}, errors.Wrap(err, "json unmarshal error")
	}

	return nodes, nil
}

// GetNodeByID evaluates a transaction to read node by ID
func GetNodeByID(ID string) (model.Node, error) {
	fmt.Println("Evaluate Transaction: GetNodeByID, returns specific node by node-ID.")

	nodeBytes, err := global.Contract.EvaluateTransaction("GetNodeByID", ID)
	if err != nil {
		return model.Node{}, errors.Wrap(err, "failed to evaluate transaction")
	}

	var node model.Node
	err = json.Unmarshal(nodeBytes, &node)
	if err != nil {
		return model.Node{}, errors.Wrap(err, "json unmarshal error")
	}

	return node, nil
}

// GetSatellitePublicKey evaluates a transaction to read public key of satellite by ID
func GetSatellitePublicKey(ID string) (string, error) {
	fmt.Println("Evaluate Transaction: GetSatellitePublicKey, returns the satellite's public key.")

	publicKeyBytes, err := global.Contract.EvaluateTransaction("GetSatellitePublicKey", ID)
	if err != nil {
		return "", errors.Wrap(err, "failed to evaluate transaction")
	}

	publicKey := string(publicKeyBytes)
	return publicKey, nil
}

// GetUserPublicKey evaluates a transaction to read public key of user by ID and mac address
func GetUserPublicKey(ID string, macAddr string) (string, error) {
	fmt.Println("Evaluate Transaction: GetUserPublicKey, returns the user's public key.")

	publicKeyBytes, err := global.Contract.EvaluateTransaction("GetUserPublicKey", ID, macAddr)
	if err != nil {
		return "", errors.Wrap(err, "failed to evaluate transaction")
	}

	publicKey := string(publicKeyBytes)
	return publicKey, nil
}
