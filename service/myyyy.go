package service

import (
	"encoding/json"
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hiro942/elden-app/model"
	"github.com/pkg/errors"
)

func InitLedger2() error {
	fmt.Println("Submit Transaction: InitLedger2, function creates the initial set of nodes on the ledger.")

	_, err := global.Contract.SubmitTransaction("InitLedger2")
	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil

}

func GetAllNodes2() ([]model.Node2, error) {
	fmt.Println("Evaluate Transaction: GetAllNodes2, returns all current nodes on the ledger.")

	nodesBytes, err := global.Contract.EvaluateTransaction("GetAllNodes2")
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate transaction")
	}

	var nodes []model.Node2
	_ = json.Unmarshal(nodesBytes, &nodes)

	return nodes, nil
}