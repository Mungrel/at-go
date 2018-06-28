package client

// Version represents a version in the AT GTFS API. See API docs for more details.
type Version struct {
	Version   string     `json:"version"`
	StartDate *Timestamp `json:"startdate"`
	EndDate   *Timestamp `json:"enddate"`
}

// Versions gets a list of versions from the AT GTFS API
func (client *Client) Versions() ([]*Version, error) {
	url := baseURL + "/v2/gtfs/versions"

	var response []*Version
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
