/*
 * SPDX-License-Identifier: Apache-2.0
 */

package ledgerapi

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AppointmentInfo ...
type AppointmentInfo struct {
	UserID                  string `json:"userID"`
	AppointmentID           string `json:"appointmentID"`
	PatientID               string `json:"patientID"`
	ProviderID              string `json:"providerID"`
	NewProviderID           string `json:"newProviderID"`
	CaregiverID             string `json:"caregiverID"`
	AppointmentStart        string `json:"appointmentStart"`
	AppointmentEnd          string `json:"appointmentEnd"`
	AppointmentDescription  string `json:"appointmentDescription"`
	AppointmentType         string `json:"appointmentType"`
	AcceptedByPatient       string `json:"acceptedByPatient"`
	AcceptedByProvider      string `json:"acceptedByProvider"`
	AcceptedByCaregiver     string `json:"acceptedByCaregiver"`
	AppointmentStatus       string `json:"appointmentStatus"`
	AppointmentCreationDate string `json:"appointmentCreationDate"`
	AppointmentUpdateDate   string `json:"appointmentUpdateDate"`
	CancellationReason      string `json:"cancellationReason"`
	Class                   string `json:"class"`
	Key                     string `json:"key"`
}

// QueryResult ...
type QueryResult struct {
	Key    string `json:"Key"`
	Record *AppointmentInfo
}

// StateListInterface functions that a state list
// should have
type StateListInterface interface {
	AddState(StateInterface) error
	GetState(string, StateInterface) error
	GetStateByPartialCompositeKey(string, string) ([]QueryResult, error)
	UpdateState(StateInterface) error
}

// StateList useful for managing putting data in and out
// of the ledger. Implementation of StateListInterface
type StateList struct {
	Ctx         contractapi.TransactionContextInterface
	Name        string
	Deserialize func([]byte, StateInterface) error
}

// AddState puts state into world state
func (sl *StateList) AddState(state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return sl.Ctx.GetStub().PutState(key, data)
}

// GetState returns state from world state. Unmarshalls the JSON
// into passed state. Key is the split key value used in Add/Update
// joined using a colon
func (sl *StateList) GetState(key string, state StateInterface) error {
	ledgerKey, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, SplitKey(key))
	fmt.Println("LedgerKey: ", ledgerKey)
	fmt.Println("Key: ", key)
	data, err := sl.Ctx.GetStub().GetState(ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s", key)
	}

	return sl.Deserialize(data, state)
}

// GetStateByPartialCompositeKey ...
func (sl *StateList) GetStateByPartialCompositeKey(key string, searchByPart string) ([]QueryResult, error) {
	fmt.Println("Key: ", key)
	var attributes []string
	attributes = append(attributes, searchByPart)
	fmt.Println("Attributes: ", attributes)

	iterator, err := sl.Ctx.GetStub().GetStateByPartialCompositeKey(sl.Name, attributes)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	results := []QueryResult{}

	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		fmt.Println("queryResponse: ", queryResponse)

		if err != nil {
			return nil, err
		}

		patient := new(AppointmentInfo)
		_ = json.Unmarshal(queryResponse.Value, patient)

		queryResult := QueryResult{Key: queryResponse.Key, Record: patient}
		results = append(results, queryResult)
	}
	fmt.Println("Final: ", results)
	return results, nil
}

// UpdateState puts state into world state. Same as AddState but
// separate as semantically different
func (sl *StateList) UpdateState(state StateInterface) error {
	return sl.AddState(state)
}
