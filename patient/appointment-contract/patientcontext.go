/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TransactionContextInterface an interface to
// describe the minimum required functions for
// a transaction context in the commercial
// paper
type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetPatientList() ListInterface
}

// TransactionContext implementation of
// TransactionContextInterface
type TransactionContext struct {
	contractapi.TransactionContext
	patientList *list
}

// GetPatientList function
func (tc *TransactionContext) GetPatientList() ListInterface {
	if tc.patientList == nil {
		tc.patientList = newList(tc)
	}

	return tc.patientList
}
