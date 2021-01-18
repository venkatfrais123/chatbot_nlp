/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	appointmentprovider "github.com/venkatfrais123/chatbot_nlp/provider/appointment-contract"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	contract := new(appointmentprovider.Contract)
	contract.TransactionContextHandler = new(appointmentprovider.TransactionContext)
	contract.Name = "org.appointmentbook.providers"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "ProviderContractChaincode"
	chaincode.Info.Version = "0.0.1"

	err = chaincode.Start()

	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err.Error()))
	}
}
