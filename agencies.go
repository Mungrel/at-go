package at

import "errors"

type agencyResponse struct {
	Status   string    `json:"status"`
	Response []*Agency `json:"response"`
	Error    *string   `json:"error"`
}

// Agency represents an agency from the AT GTFS API
type Agency struct {
	ID       string `json:"agency_id"`
	Name     string `json:"agency_name"`
	URL      string `json:"agency_url"`
	Timezone string `json:"agency_timezone"`
	Language string `json:"agency_lang"`
	Phone    string `json:"agency_phone"`
	FareURL  string `json:"agency_fare_url"`
}

// Agencies returns a list of agencies from the AT GTFS API
func (client *Client) Agencies() ([]*Agency, error) {
	url := baseURL + "/v2/gtfs/agency"

	var response agencyResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
