package at

import "errors"

type metadataResponse struct {
	Status   string  `json:"status"`
	Response string  `json:"response"`
	Error    *string `json:"error"`
}

// Metadata returns the response field of the metadata endpoint
func (client *Client) Metadata(ivu *bool) (*string, error) {
	url := baseURL + "/v2/gtfs/metadata"

	var response metadataResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return &response.Response, nil
}
