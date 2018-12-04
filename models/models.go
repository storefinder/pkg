package models

//StoreRecord represents a record in Stores table
type StoreRecord struct {
	StoreCode       string         `json:"store_code,omitempty"`
	BusinessName    string         `json:"business_name,omitempty"`
	Address1        string         `json:"address_1,omitempty"`
	Address2        string         `json:"address_2,omitempty"`
	City            string         `json:"city,omitempty"`
	State           string         `json:"state,omitempty"`
	PostalCode      string         `json:"postal_code,omitempty"`
	Country         string         `json:"country,omitempty"`
	PrimaryPhone    string         `json:"primary_phone,omitempty"`
	Website         string         `json:"website,omitempty"`
	Description     string         `json:"description,omitempty"`
	PaymentTypes    string         `json:"payment_types,omitempty"`
	PrimaryCategory string         `json:"primary_category,omitempty"`
	Photo           string         `json:"photo,omitempty"`
	Hours           []*StoreHour   `json:"store_hours,omitempty"`
	Location        *StoreLocation `json:"location,omitempty"`
	SapID           string         `json:"sap_id,omitempty"`
}

//StoreHour represents store hours of operation
type StoreHour struct {
	DayOfWeek string `json:"day_of_week,omitempty"`
	OpenTime  string `json:"open_time,omitempty"`
	CloseTime string `json:"close_time,omitempty"`
}

//StoreLocation represents store location
type StoreLocation struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}
