package entity

type Ticket struct {
	DocNumber    string `json:"docNumber"`
	BoardingTime string `json:"boardingTime"`
	Date         string `json:"date"`
	From         string `json:"from"`
	To           string `json:"to"`
	Seat         string `json:"seat"`
	Gate         string `json:"gate"`
	FlightNumber string `json:"flightNumber"`
}
