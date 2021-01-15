/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import ledgerapi "booking-app-2/core-files/organization/patientorg/contract-go/ledger-api"

// ListInterface defines functionality needed
// to interact with the world state on behalf
// of a commercial paper
type ListInterface interface {
	AddPatient(*PatientInfo) error
	GetPatient(string, string) (*PatientInfo, error)
	UpdatePatient(*PatientInfo) error
}

type list struct {
	stateList ledgerapi.StateListInterface
}

func (cpl *list) AddPatient(patient *PatientInfo) error {
	return cpl.stateList.AddState(patient)
}

func (cpl *list) GetPatient(patientID string, userID string) (*PatientInfo, error) {
	cp := new(PatientInfo)

	err := cpl.stateList.GetState(CreatePatientKey(patientID, userID), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (cpl *list) UpdatePatient(patient *PatientInfo) error {
	return cpl.stateList.UpdateState(patient)
}

// NewList create a new list from context
func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.appointmentbook.patientlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*PatientInfo))
	}

	list := new(list)
	list.stateList = stateList

	return list
}
