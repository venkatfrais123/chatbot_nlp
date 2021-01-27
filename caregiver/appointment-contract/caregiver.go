/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcaregiver

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/venkatfrais123/chatbot_nlp/caregiver/ledger-api"
)

// CreateCaregiverKey ...
func CreateCaregiverKey(userID string) string {
	fmt.Println("Ledger Key: ", ledgerapi.MakeKey("CAREGIVERS", userID))
	return ledgerapi.MakeKey("CAREGIVERS", userID)
}

// CaregiverAlias ...
type CaregiverAlias CaregiverInfo

// PersonAlias ..
type PersonAlias Person
type jsonCaregiver struct {
	*CaregiverAlias
	Class string `json:"class"`
	Key   string `json:"key"`
}

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
	class       string       `metadata:"class"`
	key         string       `metadata:"key"`
}

// UnmarshalJSON special handler for managing JSON marshalling
func (cp *CaregiverInfo) UnmarshalJSON(data []byte) error {
	jcp := jsonCaregiver{CaregiverAlias: (*CaregiverAlias)(cp)}

	err := json.Unmarshal(data, &jcp)

	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON special handler for managing JSON marshalling
func (cp CaregiverInfo) MarshalJSON() ([]byte, error) {
	jcp := jsonCaregiver{CaregiverAlias: (*CaregiverAlias)(&cp), Class: "org.appointmentbook.caregivers", Key: ledgerapi.MakeKey("CAREGIVERS", cp.UserID)}

	return json.Marshal(&jcp)
}

// GetPerson ...
func (cp *CaregiverInfo) GetPerson() PersonAlias {
	return *cp.PersonAlias
}

// SetPerson ..
func (cp *CaregiverInfo) SetPerson(newPerson PersonAlias) {
	cp.PersonAlias = &newPerson
}

// GetSplitKey returns values which should be used to form key
func (cp *CaregiverInfo) GetSplitKey() []string {
	return []string{"CAREGIVERS", cp.UserID}
}

// Serialize formats as JSON bytes
func (cp *CaregiverInfo) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats from JSON bytes
func Deserialize(bytes []byte, cp *CaregiverInfo) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing. %s", err.Error())
	}

	return nil
}
