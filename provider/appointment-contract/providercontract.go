/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentprovider

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
	fmt.Println("Instantiate")
}

// CreateProvider ...
func (c *Contract) CreateProvider(ctx TransactionContextInterface, userID string, person string, providerID string, providerType string, institutionName string, institutionPhone string, institutionEmail string, institutionAddress string) (*ProviderInfo, error) {

	var personresponse Person
	fmt.Println([]byte(person))
	err := json.Unmarshal([]byte(person), &personresponse)
	fmt.Println("Person Resp: ", &personresponse)

	var addressresponse InstitutionAddress
	err = json.Unmarshal([]byte(institutionAddress), &addressresponse)
	fmt.Println("Address Response: ", &addressresponse)

	provider := ProviderInfo{UserID: userID, PersonAlias: (*PersonAlias)(&personresponse), ProviderID: providerID, ProviderType: providerType, InstitutionName: institutionName, InstitutionPhone: institutionPhone, InstitutionEmail: institutionEmail, InstitutionAddressAlias: (*InstitutionAddressAlias)(&addressresponse)}

	fmt.Println("Inside: ", provider)
	err = ctx.GetProviderList().AddProvider(&provider)

	if err != nil {
		return nil, err
	}

	return &provider, nil
}

// ListProvider ...
func (c *Contract) ListProvider(ctx TransactionContextInterface, userID string, providerID string) (*ProviderInfo, error) {
	fmt.Println("Provider Query...")

	fmt.Println("ProviderID: ", providerID)
	provider, err := ctx.GetProviderList().GetProvider(providerID, userID)

	if err != nil {
		return nil, err
	}

	return provider, nil
}

// UpdateProvider ...
func (c *Contract) UpdateProvider(ctx TransactionContextInterface, userID string, person PersonAlias, providerID string, providerType string, institutionName string, institutionPhone string, institutionEmail string, institutionAddress InstitutionAddressAlias) (*ProviderInfo, error) {
	fmt.Println("Provider Update...")

	fmt.Println("Provider ID: ", providerID)

	provider, err := ctx.GetProviderList().GetProvider(providerID, userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Caregiver Details Before Update: ", provider)

	provider.SetInstitutionName(institutionName)
	provider.SetInstitutionPhone(institutionPhone)
	provider.SetInstitutionEmail(institutionEmail)
	provider.SetInsAddress(institutionAddress)
	provider.SetPerson(person)

	err = ctx.GetProviderList().UpdateProvider(provider)

	if err != nil {
		return nil, err
	}

	return provider, nil
}
