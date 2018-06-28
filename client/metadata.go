package client

// Metadata returns the response field of the metadata endpoint
func (client *Client) Metadata(ivu *bool) (*string, error) {
	url := baseURL + "/v2/gtfs/metadata"

	var response string
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
