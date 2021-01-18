/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcaregiver

import (
	ledgerapi "github.com/venkatfrais123/chatbot_nlp/caregiver/ledger-api"
)

// ListInterface defines functionality needed
// to interact with the world state on behalf
// of a commercial paper
type ListInterface interface {
	AddCaregiver(*CaregiverInfo) error
	GetCaregiver(string, string) (*CaregiverInfo, error)
	UpdateCaregiver(*CaregiverInfo) error
}

var querystring = "CAREGIVERS"

type list struct {
	stateList ledgerapi.StateListInterface
}

func (cpl *list) AddCaregiver(caregiver *CaregiverInfo) error {
	return cpl.stateList.AddState(caregiver)
}

func (cpl *list) GetCaregiver(caregiverID string, userID string) (*CaregiverInfo, error) {
	cp := new(CaregiverInfo)

	err := cpl.stateList.GetState(CreateCaregiverKey(caregiverID, userID), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (cpl *list) UpdateCaregiver(caregiver *CaregiverInfo) error {
	return cpl.stateList.UpdateState(caregiver)
}

// NewList create a new list from context
func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.appointmentbook.caregiverlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*CaregiverInfo))
	}

	list := new(list)
	list.stateList = stateList

	return list
}
