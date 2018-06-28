package client

import "errors"

type versionsResponse struct {
	Status   string     `json:"status"`
	Response []*Version `json:"response"`
	Error    *string    `json:"error"`
}

// Version represents a version in the AT GTFS API. See API docs for more details.
type Version struct {
	Version   string     `json:"version"`
	StartDate *Timestamp `json:"startdate"`
	EndDate   *Timestamp `json:"enddate"`
}

// Versions gets a list of versions from the AT GTFS API
func (client *Client) Versions() ([]*Version, error) {
	url := baseURL + "/v2/gtfs/versions"

	var response versionsResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
