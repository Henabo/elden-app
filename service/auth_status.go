package service

import (
	"fmt"

	"github.com/hiro942/elden-app/global"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/pkg/errors"
)

// ChangeAuthStatus submit a transaction to change user's auth-status
func ChangeAuthStatus(id string) error {
	fmt.Println("Async Submit Transaction: ChangeAuthStatus, change user's auth status.")

	_, commit, err := global.Contract.SubmitAsync("ChangeAuthStatus", client.WithArguments(id))
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
