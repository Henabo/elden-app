package service

import (
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/pkg/errors"
)

// SatelliteRegister submits a transaction to add public_key of the satellite to the state
func SatelliteRegister(ID string, publicKey string) error {
	fmt.Println("Submit Transaction: SatelliteRegister, registers satellite by adding the public_key")

	_, err := global.Contract.SubmitTransaction("SatelliteRegister", ID, publicKey)

	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}

//// UserRegister submits a transaction to add public_key of the user to the state
//func UserRegister(ID string, macAddr string, publicKey string) error {
//	fmt.Println("Submit Transaction: UserRegister, registers satellite by adding the public_key")
//
//	_, err := global.Contract.SubmitTransaction("UserRegister", ID, macAddr, publicKey)
//
//	if err != nil {
//		return errors.Wrap(err, "failed to submit transaction")
//	}
//
//	fmt.Println("*** Transaction committed successfully.")
//	return nil
//}

// UserRegister submits a transaction to add public_key of the user to the state
func UserRegister(ID string, macAddr string, publicKey string) error {
	fmt.Println("Async Submit Transaction: UserRegister, registers satellite by adding the public_key")

	_, commit, err := global.Contract.SubmitAsync("UserRegister", client.WithArguments(ID, macAddr, publicKey))
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
