package common

// Identifier is identity of resource
type Identifier struct {
	Href string `json:"href,omitempty" binding:"omitempty,url"`
	ID   string `json:"id" binding:"omitempty,uuid3"`
}

// Address information
type Address struct {
	City       string `json:"city"`
	Country    string `json:"country,omitempty" binding:"omitempty,alpha"`
	PostalCode string `json:"postal_code,omitempty"`
	Street     string `json:"street"`
	Type       string `json:"type" binding:"omitempty,alpha,oneof=home work other"`
	Zip        string `json:"zip,omitempty" binding:"omitempty,numeric"`
}

// Phone information
type Phone struct {
	Number string `json:"number"`
	Type   string `json:"type" binding:"omitempty,alpha,oneof=home mobile other"`
}
