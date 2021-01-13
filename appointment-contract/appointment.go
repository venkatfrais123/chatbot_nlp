/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/hyperledger/fabric-samples/commercial-paper/organization/magnetocorp/contract-go/ledger-api"
)

// CreateCommercialPaperKey creates a key for commercial papers
func CreateAppKey(patientID string, appointmentID string) string {
	return ledgerapi.MakeKey("APPOINTMENTS", patientID, appointmentID)
}

// CommercialPaper defines a commercial paper
type AppInfo struct {
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

// UnmarshalJSON special handler for managing JSON marshalling
func (cp *AppInfo) UnmarshalAppJSON(data []byte) error {
	//jcp := jsonCommercialPaper{commercialPaperAlias: (*commercialPaperAlias)(cp)}

	err := json.Unmarshal(data, cp)

	if err != nil {
		return err
	}

	//cp.state = jcp.State

	return nil
}

// MarshalJSON special handler for managing JSON marshalling
func (cp AppInfo) MarshalAppJSON() ([]byte, error) {
	//jcp := jsonCommercialPaper{commercialPaperAlias: (*commercialPaperAlias)(&cp), Class: "org.appointment.patient", Key: ledgerapi.MakeKey(cp.Issuer, cp.PaperNumber)}
	cp.Class = "org.appointment.appointments"
	cp.Key = ledgerapi.MakeKey("APPOINTMENTS", cp.PatientID, cp.AppointmentID)

	return json.Marshal(cp)
}

func (cp *AppInfo) GetAppointmentID() string {
	return cp.AppointmentID
}

func (cp *AppInfo) SetAppointmentID(newAppointmentID string) {
	cp.AppointmentID = newAppointmentID
}

func (cp *AppInfo) GetPatientID() string {
	return cp.PatientID
}

func (cp *AppInfo) SetPatientID(newPatientID string) {
	cp.PatientID = newPatientID
}

func (cp *AppInfo) GetProviderID() string {
	return cp.ProviderID
}

func (cp *AppInfo) SetProviderID(newProviderID string) {
	cp.ProviderID = newProviderID
}

func (cp *AppInfo) GetNewProviderID() string {
	return cp.NewProviderID
}

func (cp *AppInfo) SetNewProviderID(newProviderID string) {
	cp.NewProviderID = newProviderID
}

func (cp *AppInfo) GetCaregiverID() string {
	return cp.CaregiverID
}

func (cp *AppInfo) SetCaregiverID(newCaregiverID string) {
	cp.CaregiverID = newCaregiverID
}

func (cp *AppInfo) GetAppointmentStart() string {
	return cp.AppointmentStart
}

func (cp *AppInfo) SetAppointmentStart(newAppointmentStart string) {
	cp.AppointmentStart = newAppointmentStart
}

func (cp *AppInfo) GetAppointmentEnd() string {
	return cp.AppointmentEnd
}

func (cp *AppInfo) SetAppointmentEnd(newAppointmentEnd string) {
	cp.AppointmentEnd = newAppointmentEnd
}

func (cp *AppInfo) GetAppointmentDescription() string {
	return cp.AppointmentDescription
}

func (cp *AppInfo) SetAppointmentDescription(newAppointmentDescription string) {
	cp.AppointmentDescription = newAppointmentDescription
}

func (cp *AppInfo) GetAppointmentType() string {
	return cp.AppointmentType
}

func (cp *AppInfo) SetAppointmentType(newAppointmentType string) {
	cp.AppointmentType = newAppointmentType
}

// GetSplitKey returns values which should be used to form key
func (cp *AppInfo) GetSplitKey() []string {
	return []string{cp.PatientID, cp.AppointmentID}
}

func (cp *AppInfo) GetAcceptedByPatient() string {
	return cp.AcceptedByPatient
}

func (cp *AppInfo) SetAcceptedByPatient(newAcceptedByPatient string) {
	cp.AcceptedByPatient = newAcceptedByPatient
}

func (cp *AppInfo) GetAcceptedByProvider() string {
	return cp.AcceptedByProvider
}

func (cp *AppInfo) SetAcceptedByProvider(newAcceptedByProvider string) {
	cp.AcceptedByProvider = newAcceptedByProvider
}

func (cp *AppInfo) GetAcceptedByCaregiver() string {
	return cp.AcceptedByCaregiver
}

func (cp *AppInfo) SetAcceptedByCaregiver(newAcceptedByCaregiver string) {
	cp.AcceptedByCaregiver = newAcceptedByCaregiver
}

func (cp *AppInfo) GetAppointmentStatus() string {
	return cp.AppointmentStatus
}

func (cp *AppInfo) SetAppointmentStatus(newAppointmentStatus string) {
	cp.AppointmentStatus = newAppointmentStatus
}

func (cp *AppInfo) GetAppointmentCreationDate() string {
	return cp.AppointmentCreationDate
}

func (cp *AppInfo) SetAppointmentCreationDate(newAppointmentCreationDate string) {
	cp.AppointmentCreationDate = newAppointmentCreationDate
}

func (cp *AppInfo) GetAppointmentUpdateDate() string {
	return cp.AppointmentUpdateDate
}

func (cp *AppInfo) SetAppointmentUpdateDate(newAppointmentUpdateDate string) {
	cp.AppointmentUpdateDate = newAppointmentUpdateDate
}

func (cp *AppInfo) GetCancellationReason() string {
	return cp.CancellationReason
}

func (cp *AppInfo) SetCancellationReason(newCancellationReason string) {
	cp.CancellationReason = newCancellationReason
}

// Serialize formats the commercial paper as JSON bytes
func (cp *AppInfo) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats the commercial paper from JSON bytes
func DeserializeApp(bytes []byte, cp *AppInfo) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing Appointment. %s", err.Error())
	}

	return nil
}
