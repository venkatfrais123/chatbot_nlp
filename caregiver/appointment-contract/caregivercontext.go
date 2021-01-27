/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcaregiver

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TransactionContextInterface an interface to
// describe the minimum required functions for
// a transaction context in the commercial
// paper
type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetCaregiverList() ListInterface
}

// TransactionContext implementation of
// TransactionContextInterface
type TransactionContext struct {
	contractapi.TransactionContext
	caregiverList *list
}

// GetCaregiverList ..
func (tc *TransactionContext) GetCaregiverList() ListInterface {
	if tc.caregiverList == nil {
		tc.caregiverList = newList(tc)
	}

	return tc.caregiverList
}
