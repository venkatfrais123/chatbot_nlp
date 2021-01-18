/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcontract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TransactionContextInterface an interface to
// describe the minimum required functions for
// a transaction context in the commercial
// paper
type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetAppointmentList() ListInterface
}

// TransactionContext implementation of
// TransactionContextInterface for use with
// commercial paper contract
type TransactionContext struct {
	contractapi.TransactionContext
	appointmentList *list
}

// GetAppointmentList return paper list
func (tc *TransactionContext) GetAppointmentList() ListInterface {
	if tc.appointmentList == nil {
		tc.appointmentList = newList(tc)
	}

	return tc.appointmentList
}
