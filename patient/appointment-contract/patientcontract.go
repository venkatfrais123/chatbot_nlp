/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

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

/*func (c *Contract) CreatePatient(ctx TransactionContextInterface, userID string, patientID string, memberID string, memberOrganization string, providerID string, providerName string, caregiverID string, caregiverName string, person string) <-chan *PatientInfo {
	patientdetails := make(chan *PatientInfo)

	go func() {
		defer close(patientdetails)
		var personresponse Person
		fmt.Println([]byte(person))
		err := json.Unmarshal([]byte(person), &personresponse)
		fmt.Println("Person Resp: ", &personresponse)
		fmt.Println("Person Resp1: ", personresponse.PersonEmail)

		patient := PatientInfo{UserID: userID, PatientID: patientID, MemberID: memberID, MemberOrganization: memberOrganization, ProviderID: providerID, ProviderName: providerName, CaregiverID: caregiverID, CaregiverName: caregiverName, PersonAlias: (*PersonAlias)(&personresponse)}
		//patient.SetIssued()

		fmt.Println("Inside: ", patient)
		err = ctx.GetPatientList().AddPatient(&patient)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		patientdetails <- &patient

	}()

	return patientdetails
}*/

// CreatePatient a new patient
func (c *Contract) CreatePatient(ctx TransactionContextInterface, userID string, patientID string, memberID string, memberOrganization string, providerID string, providerName string, caregiverID string, caregiverName string, person string) (*PatientInfo, error) {

	var personresponse Person
	fmt.Println([]byte(person))
	err := json.Unmarshal([]byte(person), &personresponse)
	fmt.Println("Person Resp: ", &personresponse)
	fmt.Println("Person Resp1: ", personresponse.PersonEmail)

	patient := PatientInfo{UserID: userID, PatientID: patientID, MemberID: memberID, MemberOrganization: memberOrganization, ProviderID: providerID, ProviderName: providerName, CaregiverID: caregiverID, CaregiverName: caregiverName, PersonAlias: (*PersonAlias)(&personresponse)}
	//patient.SetIssued()

	fmt.Println("Inside: ", patient)
	err = ctx.GetPatientList().AddPatient(&patient)

	if err != nil {
		return nil, err
	}

	return &patient, nil
}

// ListPatient Query...
func (c *Contract) ListPatient(ctx TransactionContextInterface, userID string, patientID string) (*PatientInfo, error) {
	fmt.Println("Patient Query...")

	fmt.Println("PatientID: ", patientID)
	patient, err := ctx.GetPatientList().GetPatient(userID)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

// UpdatePatient Patient...
func (c *Contract) UpdatePatient(ctx TransactionContextInterface, userID string, patientID string, providerID string, providerName string, caregiverID string, caregiverName string, person PersonAlias) (*PatientInfo, error) {
	fmt.Println("Patient Update...")

	fmt.Println("Patient ID: ", patientID)

	patient, err := ctx.GetPatientList().GetPatient(userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Patient Details Before Update: ", patient)

	patient.SetProviderID(providerID)
	patient.SetProviderName(providerName)
	patient.SetCaregiverID(caregiverID)
	patient.SetCaregiverName(caregiverName)
	patient.SetPerson(person)

	err = ctx.GetPatientList().UpdatePatient(patient)

	if err != nil {
		return nil, err
	}

	return patient, nil
}

/*
// CreateAppointment creating a new appointment
func (c *Contract) CreateAppointment(ctx TransactionContextInterfaceApp, userID string, appointmentID string, patientID string, providerID string, newProviderID string, caregiverID string, appointmentStart string, appointmentEnd string, appointmentDescription string, appointmentType string, acceptedByPatient string, acceptedByProvider string, AcceptedByCaregiver string, appointmentStatus string, appointmentCreationDate string, appointmentUpdateDate string, cancellationReason string) (*AppointmentInfo, error) {
	appointment := AppointmentInfo{UserID: userID, AppointmentID: appointmentID, PatientID: patientID, ProviderID: providerID, NewProviderID: newProviderID, CaregiverID: caregiverID, AppointmentStart: appointmentStart, AppointmentEnd: appointmentEnd, AppointmentDescription: appointmentDescription, AppointmentType: appointmentType, AcceptedByPatient: acceptedByPatient, AcceptedByProvider: acceptedByProvider, AcceptedByCaregiver: AcceptedByCaregiver, AppointmentStatus: appointmentStatus, AppointmentCreationDate: appointmentCreationDate, AppointmentUpdateDate: appointmentUpdateDate, CancellationReason: cancellationReason}

	err := ctx.GetAppointmentList().AddAppointment(&appointment)

	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

// UpdateAppointmentPat Appointment for Patient...
func (c *Contract) UpdateAppointmentPat(ctx TransactionContextInterfaceApp, patientID string, appointmentID string, appointmentUpdateDate string, acceptedByPatient string, cancellationReason string) (*AppointmentInfo, error) {
	fmt.Println("Appointment Update by patient...")

	appointment, err := ctx.GetAppointmentList().GetAppointment(patientID, appointmentID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Appointment Details Before Update: ", appointment)

	if acceptedByPatient == "true" {
		appointment.SetAcceptedByPatient(acceptedByPatient)
		appointment.SetAppointmentStatus("AppointmentAcceptedByPatient")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
	} else {
		appointment.SetAcceptedByPatient(acceptedByPatient)
		appointment.SetAppointmentStatus("AppointmentRejectedByPatient")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
		appointment.SetCancellationReason(cancellationReason)
	}

	err = ctx.GetAppointmentList().UpdateAppointment(appointment)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}

// ListAppointment Query...
func (c *Contract) ListAppointment(ctx TransactionContextInterfaceApp, patientID string, appointmentID string) (*AppointmentInfo, error) {
	fmt.Println("Appointment Query...")

	appointment, err := ctx.GetAppointmentList().GetAppointment(patientID, appointmentID)

	if err != nil {
		return nil, err
	}

	return appointment, nil
} */