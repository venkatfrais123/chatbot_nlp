/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcontract

import (
	ledgerapi "github.com/venkatfrais123/chatbot_nlp/appointment/ledger-api"
)

// ListInterface defines functionality needed
// to interact with the world state
type ListInterface interface {
	AddAppointment(*AppointmentInfo) error
	GetAppointment(string, string) (*AppointmentInfo, error)
	GetAppointmentByPartialCompositeKey(string, string) ([]ledgerapi.QueryResult, error)
	UpdateAppointment(*AppointmentInfo) error
}

type list struct {
	stateList ledgerapi.StateListInterface
}

func (cpl *list) AddAppointment(appointment *AppointmentInfo) error {
	return cpl.stateList.AddState(appointment)
}

func (cpl *list) GetAppointment(patientID string, appointmentID string) (*AppointmentInfo, error) {
	cp := new(AppointmentInfo)

	err := cpl.stateList.GetState(CreateAppointmentKey(patientID, appointmentID), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (cpl *list) GetAppointmentByPartialCompositeKey(appointmentKey string, searchByPart string) ([]ledgerapi.QueryResult, error) {
	//cp := []QueryResult{}

	cp, err := cpl.stateList.GetStateByPartialCompositeKey(appointmentKey, searchByPart)
	if err != nil {
		return nil, err
	}

	return cp, err
}

func (cpl *list) UpdateAppointment(appointment *AppointmentInfo) error {
	return cpl.stateList.UpdateState(appointment)
}

// NewList create a new list from context
func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.appointmentbook.appointmentlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*AppointmentInfo))
	}

	list := new(list)
	list.stateList = stateList

	return list
}
