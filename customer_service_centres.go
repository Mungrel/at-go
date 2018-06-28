package at

import "errors"

type customerServiceCentreResponse struct {
	Status   string                   `json:"status"`
	Response []*CustomerServiceCentre `json:"response"`
	Error    *string                  `json:"error"`
}

// CustomerServiceCentre represents a customer service location in the AT Locations API
type CustomerServiceCentre struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Suburb      string  `json:"suburb"`
	Phone       string  `json:"phone"`
	OpenHours   string  `json:"openHours"`
	HopServices bool    `json:"hopServices"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

// CustomerServiceCentres gets the list of customer service centres from the AT Locations API
func (client *Client) CustomerServiceCentres() ([]*CustomerServiceCentre, error) {
	url := baseURL + "/v2/locations/customerservicecentres"

	var response customerServiceCentreResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
