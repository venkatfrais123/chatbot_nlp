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
type TransactionContextInterfaceApp interface {
	contractapi.TransactionContextInterface
	GetAppointmentList() ListInterfaceApp
}

// TransactionContext implementation of
// TransactionContextInterface for use with
// commercial paper contract
type TransactionContextApp struct {
	contractapi.TransactionContext
	appointmentList *listApp
}

// GetPaperList return paper list
func (tc *TransactionContextApp) GetAppointmentList() ListInterfaceApp {
	if tc.appointmentList == nil {
		tc.appointmentList = newListApp(tc)
	}

	return tc.appointmentList
}
