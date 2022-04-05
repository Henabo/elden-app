package service

import (
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/pkg/errors"
)

// SatelliteRegister submits a transaction to add public_key of the satellite to the state
func SatelliteRegister(id string, publicKey string) error {
	fmt.Println("Submit Transaction: SatelliteRegister, registers satellite by adding the public_key")

	submitResult, commit, err := global.Contract.SubmitAsync("SatelliteRegister", client.WithArguments(id, publicKey))
	if err != nil {
		return errors.Wrap(err, "failed to submit transaction asynchronously")
	}

	fmt.Println("Successfully submitted transaction.")
	fmt.Println("Waiting for transaction commit.")

	status, err := commit.Status()
	if err != nil {
		return errors.Wrap(err, "failed to get commit status")
	} else if !status.Successful {
		return errors.Wrapf(err, "transaction %s failed to commit with status: %d", status.TransactionID, int32(status.Code))
	}

	fmt.Println("*** Transaction committed successfully.")
	fmt.Printf("submit result: %s", string(submitResult))
	return nil
}

// UserRegister submits a transaction to add public_key of the user to the state
func UserRegister(id string, macAddr string, publicKey string) error {
	fmt.Println("Submit Transaction: UserRegister, registers user by adding the public_key")

	submitResult, commit, err := global.Contract.SubmitAsync("UserRegister", client.WithArguments(id, macAddr, publicKey))
	if err != nil {
		return errors.Wrap(err, "failed to submit transaction asynchronously")
	}

	fmt.Println("Successfully submitted transaction.")
	fmt.Println("Waiting for transaction commit.")

	status, err := commit.Status()
	if err != nil {
		return errors.Wrap(err, "failed to get commit status")

	} else if !status.Successful {
		return errors.Wrapf(err, "transaction %s failed to commit with status: %d", status.TransactionID, int32(status.Code))
	}

	fmt.Println("*** Transaction committed successfully.")
	fmt.Printf("submit result: %s", string(submitResult))
	return nil
}