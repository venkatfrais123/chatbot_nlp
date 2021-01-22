/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentprovider

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/venkatfrais123/chatbot_nlp/provider/ledger-api"
)

// CreateProviderKey ...
func CreateProviderKey(userID string) string {
	fmt.Println("Ledger Key: ", ledgerapi.MakeKey("PROVIDERS", userID))
	return ledgerapi.MakeKey("PROVIDERS", userID)
}

// ProviderAlias ...
type ProviderAlias ProviderInfo

// PersonAlias ....
type PersonAlias Person

// InstitutionAddressAlias ...
type InstitutionAddressAlias InstitutionAddress

// jsonProvider ...
type jsonProvider struct {
	*ProviderAlias
	Class string `json:"class"`
	Key   string `json:"key"`
}

// Person ...
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

// InstitutionAddress ...
type InstitutionAddress struct {
	InsAddress1 string `json:"insaddress1"`
	InsAddress2 string `json:"insaddress2"`
	InsCity     string `json:"inscity"`
	InsState    string `json:"insstate"`
	InsZip      string `json:"inszip"`
}

// ProviderInfo ...
type ProviderInfo struct {
	UserID                  string                   `json:"userID"`
	PersonAlias             *PersonAlias             `json:"person"`
	ProviderID              string                   `json:"providerID"`
	ProviderType            string                   `json:"providerType"`
	InstitutionName         string                   `json:"institutionName"`
	InstitutionPhone        string                   `json:"institutionPhone"`
	InstitutionEmail        string                   `json:"institutionEmail"`
	InstitutionAddressAlias *InstitutionAddressAlias `json:"institutionAddress"`
	class                   string                   `metadata:"class"`
	key                     string                   `metadata:"key"`
}

// UnmarshalJSON special handler for managing JSON marshalling
func (cp *ProviderInfo) UnmarshalJSON(data []byte) error {
	jcp := jsonProvider{ProviderAlias: (*ProviderAlias)(cp)}

	err := json.Unmarshal(data, &jcp)

	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON special handler for managing JSON marshalling
func (cp ProviderInfo) MarshalJSON() ([]byte, error) {
	jcp := jsonProvider{ProviderAlias: (*ProviderAlias)(&cp), Class: "org.appointmentbook.providers", Key: ledgerapi.MakeKey("PROVIDERS", cp.UserID)}

	return json.Marshal(&jcp)
}

// GetPerson ...
func (cp *ProviderInfo) GetPerson() PersonAlias {
	return *cp.PersonAlias
}

// SetPerson ..
func (cp *ProviderInfo) SetPerson(newPerson PersonAlias) {
	cp.PersonAlias = &newPerson
}

// GetProviderID ...
func (cp *ProviderInfo) GetProviderID() string {
	return cp.ProviderID
}

// SetProviderID ..
func (cp *ProviderInfo) SetProviderID(newProviderID string) {
	cp.ProviderID = newProviderID
}

// GetProviderType ...
func (cp *ProviderInfo) GetProviderType() string {
	return cp.ProviderType
}

// SetProviderType ..
func (cp *ProviderInfo) SetProviderType(newProviderType string) {
	cp.ProviderType = newProviderType
}

// GetInstitutionName ...
func (cp *ProviderInfo) GetInstitutionName() string {
	return cp.InstitutionName
}

// SetInstitutionName ..
func (cp *ProviderInfo) SetInstitutionName(newInstitutionName string) {
	cp.InstitutionName = newInstitutionName
}

// GetInstitutionPhone ...
func (cp *ProviderInfo) GetInstitutionPhone() string {
	return cp.InstitutionPhone
}

// SetInstitutionPhone ..
func (cp *ProviderInfo) SetInstitutionPhone(newInstitutionPhone string) {
	cp.InstitutionPhone = newInstitutionPhone
}

// GetInstitutionEmail ...
func (cp *ProviderInfo) GetInstitutionEmail() string {
	return cp.InstitutionEmail
}

// SetInstitutionEmail ..
func (cp *ProviderInfo) SetInstitutionEmail(newInstitutionEmail string) {
	cp.InstitutionEmail = newInstitutionEmail
}

// GetInsAddress ...
func (cp *ProviderInfo) GetInsAddress() PersonAlias {
	return *cp.PersonAlias
}

// SetInsAddress ..
func (cp *ProviderInfo) SetInsAddress(newAddress InstitutionAddressAlias) {
	cp.InstitutionAddressAlias = &newAddress
}

// GetSplitKey returns values which should be used to form key
func (cp *ProviderInfo) GetSplitKey() []string {
	return []string{"PROVIDERS", cp.UserID}
}

// Serialize formats the commercial paper as JSON bytes
func (cp *ProviderInfo) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats the commercial paper from JSON bytes
func Deserialize(bytes []byte, cp *ProviderInfo) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing. %s", err.Error())
	}

	return nil
}
