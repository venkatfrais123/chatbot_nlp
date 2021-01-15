/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import (
	"encoding/json"
	"fmt"

	ledgerapi "booking-app-2/core-files/organization/patientorg/contract-go/ledger-api"
)

// CreateCommercialPaperKey creates a key for commercial papers
func CreatePatientKey(patientID string, userID string) string {
	return ledgerapi.MakeKey("PATIENTS", patientID, userID)
}

// Used for managing the fact status is private but want it in world state
type PatientAlias PatientInfo
type PersonAlias Person
type jsonPatient struct {
	*PatientAlias
	Class string `json:"class"`
	Key   string `json:"key"`
}

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

// CommercialPaper defines a commercial paper
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
	jcp := jsonPatient{PatientAlias: (*PatientAlias)(&cp), Class: "org.appointmentbook.appointments", Key: ledgerapi.MakeKey(cp.PatientID, cp.UserID)}

	return json.Marshal(&jcp)
}

func (cp *PatientInfo) GetMemberID() string {
	return cp.MemberID
}

func (cp *PatientInfo) SetMemberID(newMemberID string) {
	cp.MemberID = newMemberID
}

func (cp *PatientInfo) GetMemberOrg() string {
	return cp.MemberOrganization
}

func (cp *PatientInfo) SetMemberOrg(newMemberOrg string) {
	cp.MemberOrganization = newMemberOrg
}

func (cp *PatientInfo) GetProviderID() string {
	return cp.ProviderID
}

func (cp *PatientInfo) SetProviderID(newProviderID string) {
	cp.ProviderID = newProviderID
}

func (cp *PatientInfo) GetProviderName() string {
	return cp.ProviderName
}

func (cp *PatientInfo) SetProviderName(newProviderName string) {
	cp.ProviderName = newProviderName
}

func (cp *PatientInfo) GetCaregiverID() string {
	return cp.CaregiverID
}

func (cp *PatientInfo) SetCaregiverID(newCaregiverID string) {
	cp.CaregiverID = newCaregiverID
}

func (cp *PatientInfo) GetCaregiverName() string {
	return cp.CaregiverName
}

func (cp *PatientInfo) SetCaregiverName(newCaregiverName string) {
	cp.CaregiverName = newCaregiverName
}

func (cp *PatientInfo) GetPerson() PersonAlias {
	return *cp.PersonAlias
}

func (cp *PatientInfo) SetPerson(newPerson PersonAlias) {
	cp.PersonAlias = &newPerson
}

// GetSplitKey returns values which should be used to form key
func (cp *PatientInfo) GetSplitKey() []string {
	return []string{cp.PatientID, cp.UserID}
}

// Serialize formats the commercial paper as JSON bytes
func (cp *PatientInfo) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats the commercial paper from JSON bytes
func Deserialize(bytes []byte, cp *PatientInfo) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing. %s", err.Error())
	}

	return nil
}
