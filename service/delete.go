package service

import (
	"fmt"

	"github.com/hiro942/elden-app/global"
	"github.com/pkg/errors"
)

// DeleteNodeById delete the specific node by given id
func DeleteNodeById(id string) error {
	fmt.Println("Submit Transaction: DeleteNodeById, delete the specific node by given id")

	_, err := global.Contract.SubmitTransaction("SatelliteRegister", id)

	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}
