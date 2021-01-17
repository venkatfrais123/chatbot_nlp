package patientmodel

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
