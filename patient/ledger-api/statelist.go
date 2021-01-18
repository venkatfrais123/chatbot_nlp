/*
 * SPDX-License-Identifier: Apache-2.0
 */

package ledgerapi

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// StateListInterface functions that a state list
// should have
type StateListInterface interface {
	AddState(StateInterface) error
	GetState(string, StateInterface) error
	//GetStateByPartialCompositeKey(string) ([]*patientmodel.PatientInfo, error)
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

/*
func (sl *StateList) GetStateByPartialCompositeKey(searchByPart string) ([]*patientmodel.PatientInfo, error) {

	var attributes []string
	attributes = append(attributes, searchByPart)
	fmt.Println("Attributes: ", attributes)

	iterator, err := sl.Ctx.GetStub().GetStateByPartialCompositeKey(sl.Name, attributes)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var records []*patientmodel.PatientInfo
	for iterator.HasNext() {
		queryresponse, err := iterator.Next()
		if err != nil {
			return nil, err
		}

		var record patientmodel.PatientInfo
		err = json.Unmarshal(queryresponse.Value, &record)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)
	}
	return records, nil
}
*/

// UpdateState puts state into world state. Same as AddState but
// separate as semantically different
func (sl *StateList) UpdateState(state StateInterface) error {
	return sl.AddState(state)
}