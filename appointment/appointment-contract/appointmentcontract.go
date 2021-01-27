/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentcontract

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/venkatfrais123/chatbot_nlp/appointment/ledger-api"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Contract chaincode that defines
// the business logic
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
	// 01/26 updated
	appointment.SetAcceptedByCaregiver("NA")

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

	if acceptedByCaregiver == "true" && appointment.NewProviderID == "NA" {
		fmt.Println("New Provider ID: ", appointment.NewProviderID)
		appointment.SetAcceptedByCaregiver(acceptedByCaregiver)
		appointment.SetAppointmentStatus("AppointmentAcceptedByCaregiver")
		appointment.SetAppointmentUpdateDate(appointmentUpdateDate)
		// 01/26 - Case 5
	} else if appointment.NewProviderID != "NA" && acceptedByCaregiver == "true" {
		fmt.Println("New Provider ID: ", appointment.NewProviderID)
		appointment.SetAcceptedByCaregiver(acceptedByCaregiver)
		appointment.SetAcceptedByProvider("NA")
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

// ListAllApps Query...
func (c *Contract) ListAllApps(ctx TransactionContextInterface) ([]ledgerapi.QueryResult, error) {
	fmt.Println("Appointment Query...")

	index := "" + "~" + "" + "~" + ""

	appointment, err := ctx.GetAppointmentList().GetAppointmentByPartialCompositeKey(index, "APPOINTMENTS")

	if err != nil {
		return nil, err
	}

	return appointment, err
}

// UserApps ..
func (c *Contract) UserApps(ctx TransactionContextInterface, userID string, orgname string) (rets []*AppointmentInfo, err error) {

	var queryString string

	if orgname == "patientorg" {
		queryString = `{"selector": {"patientID":{"$eq":"` + userID + `"}}}`
	} else if orgname == "providerorg" {
		queryString = `{"selector": {"$or": [{"providerID": {"$eq": "` + userID + `"}},{"newProviderID": {"$eq": "` + userID + `"}}]}}`
	} else {
		queryString = `{"selector": {"caregiverID":{"$eq":"` + userID + `"}}}`
	}

	fmt.Println("QueryStrng: ", queryString)

	resultsIterator, _, err := ctx.GetStub().GetQueryResultWithPagination(queryString, 0, "")
	fmt.Println("ResultIterator: ", resultsIterator)
	if err != nil {
		return
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, err2 := resultsIterator.Next()
		if err2 != nil {
			return nil, err2
		}

		fmt.Println("queryresp: ", queryResponse.Value)

		res := new(AppointmentInfo)
		if err = json.Unmarshal(queryResponse.Value, res); err != nil {
			return
		}

		rets = append(rets, res)
	}
	fmt.Println("Rets: ", rets)

	return rets, err
}
