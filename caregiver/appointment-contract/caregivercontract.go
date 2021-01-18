/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcaregiver

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Contract chaincode that defines
// the business logic for managing commercial
// paper
type Contract struct {
	contractapi.Contract
}

// Instantiate does nothing
func (c *Contract) Instantiate() {
	fmt.Println("Instantiated")
}

// Creates a new caregiver
func (c *Contract) CreateCaregiver(ctx TransactionContextInterface, userID string, person string, caregiverID string) (*CaregiverInfo, error) {

	var personresponse Person
	fmt.Println([]byte(person))
	err := json.Unmarshal([]byte(person), &personresponse)
	fmt.Println("Person Resp: ", &personresponse)

	caregiver := CaregiverInfo{UserID: userID, PersonAlias: (*PersonAlias)(&personresponse), CaregiverID: caregiverID}

	fmt.Println("Inside: ", caregiver)
	err = ctx.GetCaregiverList().AddCaregiver(&caregiver)

	if err != nil {
		return nil, err
	}

	return &caregiver, nil
}

// Caregiver Query...
func (c *Contract) ListCaregiver(ctx TransactionContextInterface, userID string, caregiverID string) (*CaregiverInfo, error) {
	fmt.Println("Caregiver Query...")

	fmt.Println("CaregiverID: ", caregiverID)
	caregiver, err := ctx.GetCaregiverList().GetCaregiver(caregiverID, userID)

	if err != nil {
		return nil, err
	}

	return caregiver, nil
}

// Update Caregiver...
func (c *Contract) UpdateCaregiver(ctx TransactionContextInterface, userID string, person PersonAlias, caregiverID string) (*CaregiverInfo, error) {
	fmt.Println("Caregiver Update...")

	fmt.Println("Caregiver ID: ", caregiverID)

	caregiver, err := ctx.GetCaregiverList().GetCaregiver(caregiverID, userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Caregiver Details Before Update: ", caregiver)

	caregiver.SetPerson(person)

	err = ctx.GetCaregiverList().UpdateCaregiver(caregiver)

	if err != nil {
		return nil, err
	}

	return caregiver, nil
}
