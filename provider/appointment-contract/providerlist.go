/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentprovider

import (
	ledgerapi "github.com/venkatfrais123/chatbot_nlp/provider/ledger-api"
)

// ListInterface defines functionality needed
// to interact with the world state on behalf
// of a commercial paper
type ListInterface interface {
	AddProvider(*ProviderInfo) error
	GetProvider(string, string) (*ProviderInfo, error)
	UpdateProvider(*ProviderInfo) error
}

type list struct {
	stateList ledgerapi.StateListInterface
}

// AddProvider ...
func (cpl *list) AddProvider(provider *ProviderInfo) error {
	return cpl.stateList.AddState(provider)
}

// GetProvider ..
func (cpl *list) GetProvider(providerID string, userID string) (*ProviderInfo, error) {
	cp := new(ProviderInfo)

	err := cpl.stateList.GetState(CreateProviderKey(providerID, userID), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

// UpdateProvider ....
func (cpl *list) UpdateProvider(caregiver *ProviderInfo) error {
	return cpl.stateList.UpdateState(caregiver)
}

// NewList create a new list from context
func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.appointmentbook.providerlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*ProviderInfo))
	}

	list := new(list)
	list.stateList = stateList

	return list
}