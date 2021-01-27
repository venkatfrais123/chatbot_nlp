/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentprovider

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TransactionContextInterface an interface to
// describe the minimum required functions for
// a transaction context in the commercial
// paper
type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetProviderList() ListInterface
}

// TransactionContext implementation of
// TransactionContextInterface
type TransactionContext struct {
	contractapi.TransactionContext
	providerList *list
}

// GetProviderList return paper list
func (tc *TransactionContext) GetProviderList() ListInterface {
	if tc.providerList == nil {
		tc.providerList = newList(tc)
	}

	return tc.providerList
}
