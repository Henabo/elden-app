package service

import (
	"encoding/json"
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hiro942/elden-app/model"
	"github.com/pkg/errors"
)

// CreateAccessRecord submit a transaction to create access record with the given arguments
func CreateAccessRecord(id string, macAddr string, satelliteId string, userAccessRecord model.UserAccessRecord) error {
	fmt.Println("Submit Transaction: CreateAccessRecord, add a new access record with given arguments.")

	userAccessRecordBytes, _ := json.Marshal(&userAccessRecord)
	_, err := global.Contract.SubmitTransaction("CreateAccessRecord", id, macAddr, satelliteId, string(userAccessRecordBytes))
	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}


