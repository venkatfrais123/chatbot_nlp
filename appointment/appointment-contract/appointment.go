/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcontract

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/venkatfrais123/chatbot_nlp/appointment/ledger-api"
)

// CreateAppointmentKey creates a key for commercial papers
func CreateAppointmentKey(patientID string, appointmentID string) string {
	fmt.Println("Ledger Key: ", ledgerapi.MakeKey("APPOINTMENTS", patientID, appointmentID))
	return ledgerapi.MakeKey("APPOINTMENTS", patientID, appointmentID)
}

// AppointmentAlias ..
type AppointmentAlias AppointmentInfo
type jsonAppointment struct {
	*AppointmentAlias
	Class string `json:"class"`
	Key   string `json:"key"`
}

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
	class                   string `metadata:"class"`
	key                     string `metadata:"key"`
}

// UnmarshalJSON special handler for managing JSON marshalling
func (cp *AppointmentInfo) UnmarshalJSON(data []byte) error {
	jcp := jsonAppointment{AppointmentAlias: (*AppointmentAlias)(cp)}

	err := json.Unmarshal(data, &jcp)

	if err != nil {
		return err
	}

	return nil
}

// MarshalJSON special handler for managing JSON marshalling
func (cp AppointmentInfo) MarshalJSON() ([]byte, error) {
	jcp := jsonAppointment{AppointmentAlias: (*AppointmentAlias)(&cp), Class: "org.appointmentbook.appointments", Key: ledgerapi.MakeKey("APPOINTMENTS", cp.PatientID, cp.AppointmentID)}

	return json.Marshal(&jcp)
}

// GetAppointmentID ..
func (cp *AppointmentInfo) GetAppointmentID() string {
	return cp.AppointmentID
}

// SetAppointmentID ..
func (cp *AppointmentInfo) SetAppointmentID(newAppointmentID string) {
	cp.AppointmentID = newAppointmentID
}

// GetPatientID ..
func (cp *AppointmentInfo) GetPatientID() string {
	return cp.PatientID
}

// SetPatientID ..
func (cp *AppointmentInfo) SetPatientID(newPatientID string) {
	cp.PatientID = newPatientID
}

// GetProviderID ..
func (cp *AppointmentInfo) GetProviderID() string {
	return cp.ProviderID
}

// SetProviderID ..
func (cp *AppointmentInfo) SetProviderID(newProviderID string) {
	cp.ProviderID = newProviderID
}

// GetNewProviderID ..
func (cp *AppointmentInfo) GetNewProviderID() string {
	return cp.NewProviderID
}

// SetNewProviderID ..
func (cp *AppointmentInfo) SetNewProviderID(newProviderID string) {
	cp.NewProviderID = newProviderID
}

// GetCaregiverID ..
func (cp *AppointmentInfo) GetCaregiverID() string {
	return cp.CaregiverID
}

// SetCaregiverID ..
func (cp *AppointmentInfo) SetCaregiverID(newCaregiverID string) {
	cp.CaregiverID = newCaregiverID
}

// GetAppointmentStart ..
func (cp *AppointmentInfo) GetAppointmentStart() string {
	return cp.AppointmentStart
}

// SetAppointmentStart ..
func (cp *AppointmentInfo) SetAppointmentStart(newAppointmentStart string) {
	cp.AppointmentStart = newAppointmentStart
}

// GetAppointmentEnd ..
func (cp *AppointmentInfo) GetAppointmentEnd() string {
	return cp.AppointmentEnd
}

// SetAppointmentEnd ..
func (cp *AppointmentInfo) SetAppointmentEnd(newAppointmentEnd string) {
	cp.AppointmentEnd = newAppointmentEnd
}

// GetAppointmentDescription ..
func (cp *AppointmentInfo) GetAppointmentDescription() string {
	return cp.AppointmentDescription
}

// SetAppointmentDescription ..
func (cp *AppointmentInfo) SetAppointmentDescription(newAppointmentDescription string) {
	cp.AppointmentDescription = newAppointmentDescription
}

// GetAppointmentType ..
func (cp *AppointmentInfo) GetAppointmentType() string {
	return cp.AppointmentType
}

// SetAppointmentType ..
func (cp *AppointmentInfo) SetAppointmentType(newAppointmentType string) {
	cp.AppointmentType = newAppointmentType
}

// GetAcceptedByPatient ..
func (cp *AppointmentInfo) GetAcceptedByPatient() string {
	return cp.AcceptedByPatient
}

// SetAcceptedByPatient ..
func (cp *AppointmentInfo) SetAcceptedByPatient(newAcceptedByPatient string) {
	cp.AcceptedByPatient = newAcceptedByPatient
}

// GetAcceptedByProvider ..
func (cp *AppointmentInfo) GetAcceptedByProvider() string {
	return cp.AcceptedByProvider
}

// SetAcceptedByProvider ..
func (cp *AppointmentInfo) SetAcceptedByProvider(newAcceptedByProvider string) {
	cp.AcceptedByProvider = newAcceptedByProvider
}

// GetAcceptedByCaregiver ..
func (cp *AppointmentInfo) GetAcceptedByCaregiver() string {
	return cp.AcceptedByCaregiver
}

// SetAcceptedByCaregiver ..
func (cp *AppointmentInfo) SetAcceptedByCaregiver(newAcceptedByCaregiver string) {
	cp.AcceptedByCaregiver = newAcceptedByCaregiver
}

// GetAppointmentStatus ..
func (cp *AppointmentInfo) GetAppointmentStatus() string {
	return cp.AppointmentStatus
}

// SetAppointmentStatus ..
func (cp *AppointmentInfo) SetAppointmentStatus(newAppointmentStatus string) {
	cp.AppointmentStatus = newAppointmentStatus
}

// GetAppointmentCreationDate ..
func (cp *AppointmentInfo) GetAppointmentCreationDate() string {
	return cp.AppointmentCreationDate
}

// SetAppointmentCreationDate ..
func (cp *AppointmentInfo) SetAppointmentCreationDate(newAppointmentCreationDate string) {
	cp.AppointmentCreationDate = newAppointmentCreationDate
}

// GetAppointmentUpdateDate ..
func (cp *AppointmentInfo) GetAppointmentUpdateDate() string {
	return cp.AppointmentUpdateDate
}

// SetAppointmentUpdateDate ..
func (cp *AppointmentInfo) SetAppointmentUpdateDate(newAppointmentUpdateDate string) {
	cp.AppointmentUpdateDate = newAppointmentUpdateDate
}

// GetCancellationReason ..
func (cp *AppointmentInfo) GetCancellationReason() string {
	return cp.CancellationReason
}

// SetCancellationReason ...
func (cp *AppointmentInfo) SetCancellationReason(newCancellationReason string) {
	cp.CancellationReason = newCancellationReason
}

// GetSplitKey returns values which should be used to form key
func (cp *AppointmentInfo) GetSplitKey() []string {
	return []string{"APPOINTMENTS", cp.PatientID, cp.AppointmentID}
}

// Serialize formats the commercial paper as JSON bytes
func (cp *AppointmentInfo) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats the commercial paper from JSON bytes
func Deserialize(bytes []byte, cp *AppointmentInfo) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("Error deserializing. %s", err.Error())
	}

	return nil
}
