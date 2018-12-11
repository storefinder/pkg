package models

//StoreRecord represents a record in Stores table
type StoreRecord struct {
	StoreCode       string        `json:"store_code"`
	BusinessName    string        `json:"business_name"`
	Address1        string        `json:"address_1"`
	Address2        string        `json:"address_2"`
	City            string        `json:"city"`
	State           string        `json:"state"`
	PostalCode      string        `json:"postal_code"`
	Country         string        `json:"country"`
	PrimaryPhone    string        `json:"primary_phone"`
	Website         string        `json:"website"`
	Description     string        `json:"description"`
	PaymentTypes    string        `json:"payment_types"`
	PrimaryCategory string        `json:"primary_category"`
	Photo           string        `json:"photo"`
	Hours           []StoreHour   `json:"store_hours"`
	Location        StoreLocation `json:"location"`
	SapID           string        `json:"sap_id"`
}

//StoreHour represents store hours of operation
type StoreHour struct {
	DayOfWeek string `json:"day_of_week"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

//StoreLocation represents store location
type StoreLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

//Error represents error
type Error struct {
	Message string `json:"message,omitempty"`
}

//StoreQueryRequest represents store locator query request
type StoreQueryRequest struct {
	Lat    float64 `json:"latitude,omitempty"`
	Lon    float64 `json:"longitude,omitempty"`
	Radius string  `json:"radius,omitempty"`
	SapID  string  `json:"sap_id,omitempty"`
}

//StoreQueryResponse represents response from store query
type StoreQueryResponse struct {
	Hits         int64         `json:"count,omitempty"`
	TookInMillis int64         `json:"took_in_millis,omitempty"`
	Stores       []StoreRecord `json:"stores,omitempty"`
	Errors       []Error       `json:"errors,omitempty"`
}

//IndexerResponse represents response from indexing request
type IndexerResponse struct {
	IndexName           string        `json:"index_name,omitempty"`
	StoresIndexed       []StoreRecord `json:"stores_indexed,omitempty"`
	StoresFailedToIndex []StoreRecord `json:"stores_failed_to_index,omitempty"`
}

//Index represents elastic search index for stores
type Index struct {
	Name             string  `json:"name,omitempty"`
	NumberOfShards   int     `json:"number_of_shards,omitempty"`
	NumberOfReplicas int     `json:"number_of_replicas,omitempty"`
	Errors           []Error `json:"errors,omitempty"`
}
