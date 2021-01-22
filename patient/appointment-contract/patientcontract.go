/*
 * SPDX-License-Identifier: Apache-2.0
 */

package appointmentpatient

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	ledgerapi "github.com/venkatfrais123/chatbot_nlp/patient/ledger-api"
)

// Contract ...
type Contract struct {
	contractapi.Contract
}

// QueryResult ..
type QueryResult struct {
	Key    string `json:"Key"`
	Record *PatientInfo
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
func (c *Contract) ListPatient(ctx TransactionContextInterface, userID string) (*PatientInfo, error) {
	fmt.Println("Patient Query...")

	fmt.Println("PatientID: ", userID)
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

// ListAllPatient Query...
func (c *Contract) ListAllPatient(ctx TransactionContextInterface) ([]ledgerapi.QueryResult, error) {
	fmt.Println("Patient Query...")

	index := "" + "~" + ""

	patient, err := ctx.GetPatientList().GetPatientByPartialCompositeKey(index, "PATIENTS")

	if err != nil {
		return nil, err
	}

	return patient, err
}

// Index ..
func (c *Contract) Index(ctx TransactionContextInterface) (rets []*PatientInfo, err error) {
	resultsIterator, _, err := ctx.GetStub().GetQueryResultWithPagination(`{"selector": {"_id":{"$ne":"-"}}}`, 0, "")
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

		res := new(PatientInfo)
		if err = json.Unmarshal(queryResponse.Value, res); err != nil {
			return
		}

		rets = append(rets, res)
	}
	fmt.Println("Rets: ", rets)

	return rets, err
}
