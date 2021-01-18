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
// TransactionContextInterface for use with
// commercial paper contract
type TransactionContext struct {
	contractapi.TransactionContext
	caregiverList *list
}

// GetPaperList return paper list
func (tc *TransactionContext) GetCaregiverList() ListInterface {
	if tc.caregiverList == nil {
		tc.caregiverList = newList(tc)
	}

	return tc.caregiverList
}
