package service

import (
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/pkg/errors"
)

// SatelliteRegister submits a transaction to add public_key of the satellite to the state
func SatelliteRegister(id string, publicKey string) error {
	fmt.Println("Submit Transaction: SatelliteRegister, registers satellite by adding the public_key")

	_, err := global.Contract.SubmitTransaction("SatelliteRegister", id, publicKey)

	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}

// UserRegister submits a transaction to add public_key of the user to the state
func UserRegister(id string, macAddr string, publicKey string) error {
	fmt.Println("Submit Transaction: UserRegister, registers satellite by adding the public_key")

	_, err := global.Contract.SubmitTransaction("UserRegister", id, macAddr, publicKey)

	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}
