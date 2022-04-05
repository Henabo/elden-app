package service

import (
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/pkg/errors"
)

// InitLedger would only be run once by an application at the first time the network started
func InitLedger() error {
	fmt.Println("Submit Transaction: InitLedger, function creates the initial set of nodes on the ledger.")

	_, err := global.Contract.SubmitTransaction("InitLedger")
	if err != nil {
		return errors.Wrap(err, "failed to submit transaction")
	}

	fmt.Println("*** Transaction committed successfully.")
	return nil
}