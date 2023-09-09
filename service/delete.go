package service

import (
	"fmt"

	"github.com/hiro942/elden-app/global"
	"github.com/pkg/errors"
)

// DeleteNodeByID delete the specific node by given ID
func DeleteNodeByID(ID string) error {
	fmt.Println("Submit Transaction: DeleteNodeByID, delete the specific node by given ID")

	_, err := global.Contract.SubmitTransaction("DeleteNodeByID", ID)

	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}
