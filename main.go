package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hiro942/elden-app/global"
	"github.com/hiro942/elden-app/router"
	"github.com/hiro942/elden-app/utils"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

func main() {
	// create gRPC connection to the PEER
	clientConnection := utils.NewGrpcConnection(global.TlsCertPath, global.GatewayPeer, global.PeerEndpoint)
	defer clientConnection.Close()

	// get gateway
	gateway, err := utils.NewGateway(clientConnection)
	if err != nil {
		log.Panicln(err)
	}
	defer gateway.Close()

	// get contract from the FABRIC NET
	global.Contract = utils.GetContract(gateway)

	fmt.Printf("%%%%%%%%%%%%%%%%")
	initLedger(global.Contract)
	getAllAssets(global.Contract)
	createAsset(global.Contract)
	getAllAssets(global.Contract)
	fmt.Printf("%%%%%%%%%%%%%%%%")


	// router start
	sysRouter := router.Routers()
	sysRouter.Run(":8080")
}

func initLedger(contract *client.Contract) {
	fmt.Printf("Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	contract.SubmitTransaction("InitLedger")


	fmt.Printf("*** Transaction committed successfully\n")
}

// Evaluate a transaction to query ledger state.
func getAllAssets(contract *client.Contract) {
	fmt.Println("Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, _ := contract.EvaluateTransaction("GetAllAssets")
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

// Submit a transaction synchronously, blocking until it has been committed to the ledger.
func createAsset(contract *client.Contract) {
	fmt.Printf("Submit Transaction: CreateAsset, creates new asset with ID, Color, Size, Owner and AppraisedValue arguments \n")

	contract.SubmitTransaction("CreateAsset", "asset-999", "yellow", "5", "Tom", "1300")

	fmt.Printf("*** Transaction committed successfully\n")
}

//Format JSON data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, data, " ", "")
	return prettyJSON.String()
}
