/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import ledgerapi "github.com/hyperledger/fabric-samples/commercial-paper/organization/magnetocorp/contract-go/ledger-api"

// ListInterface defines functionality needed
// to interact with the world state on behalf
// of a commercial paper
type ListInterfaceApp interface {
	AddAppointment(*AppInfo) error
	GetAppointment(string, string) (*AppInfo, error)
	UpdateAppointment(*AppInfo) error
}

type listApp struct {
	stateList ledgerapi.StateListInterface
}

func (cpl *listApp) AddAppointment(appointment *AppInfo) error {
	return cpl.stateList.AddState(appointment)
}

func (cpl *listApp) GetAppointment(patientID string, appointmentID string) (*AppInfo, error) {
	cp := new(AppInfo)

	err := cpl.stateList.GetState(CreateAppKey(patientID, appointmentID), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (cpl *listApp) UpdateAppointment(appointment *AppInfo) error {
	return cpl.stateList.UpdateState(appointment)
}

// NewList create a new list from context
func newListApp(ctx TransactionContextInterfaceApp) *listApp {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.appointment.appointmentlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return DeserializeApp(bytes, state.(*AppInfo))
	}

	list := new(listApp)
	list.stateList = stateList

	return list
}
