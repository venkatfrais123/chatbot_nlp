/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/venkatfrais123/chatbot_nlp/patient/ledger-api"
)

// CreatePatientKey creates a key
func CreatePatientKey(userID string) string {
	//fmt.Println("Ledger Key: ", ledgerapi.MakeKey("PATIENTS", patientID, userID))
	fmt.Println("Ledger Key: ", ledgerapi.MakeKey("PATIENTS", userID))
	//return ledgerapi.MakeKey("PATIENTS", patientID, userID)
	return ledgerapi.MakeKey("PATIENTS", userID)
}

// PatientAlias ...
type PatientAlias PatientInfo

// PersonAlias ...
type PersonAlias Person
type jsonPatient struct {
	*PatientAlias
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

// PatientInfo ..
type PatientInfo struct {
	UserID             string       `json:"userID"`
	PatientID          string       `json:"patientID"`
	MemberID           string       `json:"memberID"`
	MemberOrganization string       `json:"memberOrganization"`
	ProviderID         string       `json:"providerID"`
	ProviderName       string       `json:"providerName"`
	CaregiverID        string       `json:"caregiverID"`
	CaregiverName      string       `json:"caregiverName"`
	PersonAlias        *PersonAlias `json:"person"`
	class              string       `metadata:"class"`
	key                string       `metadata:"key"`
}

// UnmarshalJSON special handler for managing JSON marshalling
func (cp *PatientInfo) UnmarshalJSON(data []byte) error {
	jcp := jsonPatient{PatientAlias: (*PatientAlias)(cp)}

	err := json.Unmarshal(data, &jcp)

	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON special handler for managing JSON marshalling
func (cp PatientInfo) MarshalJSON() ([]byte, error) {
	//jcp := jsonPatient{PatientAlias: (*PatientAlias)(&cp), Class: "org.appointmentbook.appointments", Key: ledgerapi.MakeKey("PATIENTS", cp.PatientID, cp.UserID)}
	jcp := jsonPatient{PatientAlias: (*PatientAlias)(&cp), Class: "org.appointmentbook.patients", Key: ledgerapi.MakeKey("PATIENTS", cp.UserID)}

	return json.Marshal(&jcp)
}

// GetMemberID ..
func (cp *PatientInfo) GetMemberID() string {
	return cp.MemberID
}

// SetMemberID ..
func (cp *PatientInfo) SetMemberID(newMemberID string) {
	cp.MemberID = newMemberID
}

// GetMemberOrg ..
func (cp *PatientInfo) GetMemberOrg() string {
	return cp.MemberOrganization
}

// SetMemberOrg ..
func (cp *PatientInfo) SetMemberOrg(newMemberOrg string) {
	cp.MemberOrganization = newMemberOrg
}

// GetProviderID ..
func (cp *PatientInfo) GetProviderID() string {
	return cp.ProviderID
}

// SetProviderID ..
func (cp *PatientInfo) SetProviderID(newProviderID string) {
	cp.ProviderID = newProviderID
}

// GetProviderName ..
func (cp *PatientInfo) GetProviderName() string {
	return cp.ProviderName
}

// SetProviderName ..
func (cp *PatientInfo) SetProviderName(newProviderName string) {
	cp.ProviderName = newProviderName
}

// GetCaregiverID ..
func (cp *PatientInfo) GetCaregiverID() string {
	return cp.CaregiverID
}

// SetCaregiverID ..
func (cp *PatientInfo) SetCaregiverID(newCaregiverID string) {
	cp.CaregiverID = newCaregiverID
}

// GetCaregiverName ..
func (cp *PatientInfo) GetCaregiverName() string {
	return cp.CaregiverName
}

// SetCaregiverName ..
func (cp *PatientInfo) SetCaregiverName(newCaregiverName string) {
	cp.CaregiverName = newCaregiverName
}

// GetPerson ..
func (cp *PatientInfo) GetPerson() PersonAlias {
	return *cp.PersonAlias
}

// SetPerson ..
func (cp *PatientInfo) SetPerson(newPerson PersonAlias) {
	cp.PersonAlias = &newPerson
}

// GetSplitKey returns values which should be used to form key
func (cp *PatientInfo) GetSplitKey() []string {
	//return []string{cp.PatientID, cp.UserID}
	return []string{"PATIENTS", cp.UserID}
}

// Serialize formats as JSON bytes
func (cp *PatientInfo) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats from JSON bytes
func Deserialize(bytes []byte, cp *PatientInfo) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing. %s", err.Error())
	}

	return nil
}
