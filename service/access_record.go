package service

import (
	"encoding/json"
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hiro942/elden-app/model"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/pkg/errors"
)

// CreateAccessRecord submit a transaction to create access record with the given arguments
func CreateAccessRecord(id string, macAddr string, userAccessRecord model.UserAccessRecord) error {
	fmt.Println("Async Submit Transaction: CreateAccessRecord, add a new access record with given arguments.")

	userAccessRecordBytes, err := json.Marshal(&userAccessRecord)
	if err != nil {
		return errors.Wrap(err, "json marshal error")
	}

	_, commit, err := global.Contract.SubmitAsync("CreateAccessRecord", client.WithArguments(id, macAddr, string(userAccessRecordBytes)))
	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("Waiting for transaction commit.")

	if status, err := commit.Status(); err != nil {
		return errors.Wrap(err, "failed to get commit status")
	} else if !status.Successful {
		return errors.Errorf("transaction %s failed to commit with status: %d", status.TransactionID, int32(status.Code))
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}



