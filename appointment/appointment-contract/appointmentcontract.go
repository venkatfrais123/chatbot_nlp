/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcontract

import (
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

// CreateAppointment ....
func (c *Contract) CreateAppointment(ctx TransactionContextInterface, userID string, appointmentID string, patientID string, providerID string, newProviderID string, caregiverID string, appointmentStart string, appointmentEnd string, appointmentDescription string, appointmentType string,
	acceptedByPatient string, acceptedByProvider string, acceptedByCaregiver string, appointmentStatus string, appointmentCreationDate string, appointmentUpdateDate string, cancellationReason string) (*AppointmentInfo, error) {

	appointment := AppointmentInfo{UserID: userID, AppointmentID: appointmentID, PatientID: patientID, ProviderID: providerID, NewProviderID: newProviderID, CaregiverID: caregiverID, AppointmentStart: appointmentStart, AppointmentEnd: appointmentEnd, AppointmentDescription: appointmentDescription, AppointmentType: appointmentType, AcceptedByPatient: acceptedByPatient, AcceptedByProvider: acceptedByProvider, AcceptedByCaregiver: acceptedByCaregiver, AppointmentStatus: appointmentStatus, AppointmentCreationDate: appointmentCreationDate, AppointmentUpdateDate: appointmentUpdateDate, CancellationReason: cancellationReason}

	fmt.Println("Inside: ", appointment)
	err := ctx.GetAppointmentList().AddAppointment(&appointment)

	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

// ListAppointment ...
func (c *Contract) ListAppointment(ctx TransactionContextInterface, patientID string, appointmentID string) (*AppointmentInfo, error) {
	fmt.Println("Appointment Query...")

	fmt.Println("AppointmentID: ", appointmentID)
	appointment, err := ctx.GetAppointmentList().GetAppointment(patientID, appointmentID)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}

// UpdateAppointmentPat ...
func (c *Contract) UpdateAppointmentPat(ctx TransactionContextInterface, patientID string, appointmentID string, appointmentUpdateDate string, acceptedByPatient string, cancellationReason string) (*AppointmentInfo, error) {
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

// UpdateAppointmentPro ...
func (c *Contract) UpdateAppointmentPro(ctx TransactionContextInterface, patientID string, appointmentID string, appointmentUpdateDate string, acceptedByProvider string, cancellationReason string) (*AppointmentInfo, error) {
	fmt.Println("Appointment Update by provider...")

	appointment, err := ctx.GetAppointmentList().GetAppointment(patientID, appointmentID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Appointment Details Before Update: ", appointment)

	if acceptedByProvider == "true" {
		appointment.SetAcceptedByProvider(acceptedByProvider)
		appointment.SetAppointmentStatus("AppointmentAcceptedByProvider")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
	} else {
		appointment.SetAcceptedByProvider(acceptedByProvider)
		appointment.SetAppointmentStatus("AppointmentRejectedByProvider")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
		appointment.SetCancellationReason(cancellationReason)
	}

	err = ctx.GetAppointmentList().UpdateAppointment(appointment)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}

// ReferAppointment ...
func (c *Contract) ReferAppointment(ctx TransactionContextInterface, patientID string, appointmentID string, appointmentUpdateDate string, newProviderID string) (*AppointmentInfo, error) {
	fmt.Println("Appointment transferred by provider...")

	appointment, err := ctx.GetAppointmentList().GetAppointment(patientID, appointmentID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Appointment Details Before Update: ", appointment)

	appointment.SetNewProviderID(newProviderID)
	appointment.SetAppointmentStatus("Appointment referred to new Provider")
	appointment.SetAppointmentUpdateDate(appointmentUpdateDate)

	err = ctx.GetAppointmentList().UpdateAppointment(appointment)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}

// UpdateAppointmentCg ...
func (c *Contract) UpdateAppointmentCg(ctx TransactionContextInterface, patientID string, appointmentID string, appointmentUpdateDate string, acceptedByCaregiver string, cancellationReason string) (*AppointmentInfo, error) {
	fmt.Println("Appointment Update by caregiver...")

	appointment, err := ctx.GetAppointmentList().GetAppointment(patientID, appointmentID)
	if err != nil {
		return nil, err
	}
	fmt.Println("Appointment Details Before Update: ", appointment)

	if acceptedByCaregiver == "true" {
		appointment.SetAcceptedByCaregiver(acceptedByCaregiver)
		appointment.SetAppointmentStatus("AppointmentAcceptedByCaregiver")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
	} else {
		appointment.SetAcceptedByCaregiver(acceptedByCaregiver)
		appointment.SetAppointmentStatus("AppointmentRejectedByCaregiver")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
		appointment.SetCancellationReason(cancellationReason)
	}

	err = ctx.GetAppointmentList().UpdateAppointment(appointment)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}
