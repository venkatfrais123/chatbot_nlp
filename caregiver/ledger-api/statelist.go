/*
 * SPDX-License-Identifier: Apache-2.0
 */

package ledgerapi

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// PersonAlias ..
type PersonAlias Person

// Person ..
type Person struct {
	PersonFirstName string `json:"personFirstName"`
	PersonLastName  string `json:"personLastName"`
	PersonEmail     string `json:"personEmail"`
	PersonPhone     string `json:"personPhone"`
	PersonAddress1  string `json:"personAddress1"`
	PersonAddress2  string `json:"personAddress2"`
	PersonCity      string `json:"personCity"`
	PersonState     string `json:"personState"`
	PersonZip       string `json:"personZip"`
}

// CaregiverInfo ...
type CaregiverInfo struct {
	UserID      string       `json:"userID"`
	PersonAlias *PersonAlias `json:"person"`
	CaregiverID string       `json:"caregiverID"`
	Class       string       `json:"class"`
	Key         string       `json:"key"`
}

// QueryResult ..
type QueryResult struct {
	Key    string `json:"Key"`
	Record *CaregiverInfo
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

		patient := new(CaregiverInfo)
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
