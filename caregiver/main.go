/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	appointmentcaregiver "github.com/venkatfrais123/chatbot_nlp/caregiver/appointment-contract"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	contract := new(appointmentcaregiver.Contract)
	contract.TransactionContextHandler = new(appointmentcaregiver.TransactionContext)
	contract.Name = "org.appointmentbook.caregivers"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "CaregiverContractChaincode"
	chaincode.Info.Version = "0.0.1"

	err = chaincode.Start()

	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err.Error()))
	}
}
