package bw

type IdentityTitle string

const (
	IdentityTitleMr  IdentityTitle = "Mr"
	IdentityTitleMrs IdentityTitle = "Mrs"
	IdentityTitleMs  IdentityTitle = "Ms"
	IdentityTitleMx  IdentityTitle = "Mx"
	IdentityTitleDr  IdentityTitle = "Dr"
)

type Identity struct {
	Title          string `json:"title"`
	FirstName      string `json:"firstName"`
	MiddleName     string `json:"middleName"`
	Address1       string `json:"address1"`
	Address2       string `json:"address2"`
	Address3       string `json:"address3"`
	City           string `json:"city"`
	State          string `json:"state"`
	PostalCode     string `json:"postalCode"`
	Country        string `json:"country"`
	Company        string `json:"company"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Ssn            string `json:"ssn"`
	Username       string `json:"username"`
	PassportNumber string `json:"passportNumber"`
	LicenseNumber  string `json:"licenseNumber"`
}
