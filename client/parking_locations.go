package client

// ParkingLocation represents a parking location in the AT Locations API
type ParkingLocation struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	MobilitySpaces int     `json:"mobilitySpaces"`
	Type           string  `json:"type"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

// ParkingLocations gets a list of parking locations from the AT Locations API
func (client *Client) ParkingLocations() ([]*ParkingLocation, error) {
	url := baseURL + "/v2/locations/parkinglocations"

	var response []*ParkingLocation
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
