package client

// Trip represents a Trip from the AT GTFS API
type Trip struct {
	RouteID       string `json:"route_id"`
	ServiceID     string `json:"service_id"`
	TripID        string `json:"trip_id"`
	TripHeadsign  string `json:"trip_headsign"`
	DirectionID   int    `json:"direction_id"`
	BlockID       string `json:"block_id"`
	ShapeID       string `json:"shape_id"`
	TripShortName string `json:"trip_short_name"`
	TripType      string `json:"trip_type"`
}

// Trips returns a list of Trips from the AT GTFS API
func (client *Client) Trips() ([]*Trip, error) {
	url := baseURL + "/v2/gtfs/trips"

	var response []*Trip
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// TripsByID returns a list of Trips, filtered by tripID from the AT GTFS API
func (client *Client) TripsByID(tripID string) ([]*Trip, error) {
	url := baseURL + "/v2/gtfs/trips/tripId/" + tripID

	var response []*Trip
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// TripsByRouteID returns a list of Trips, filtered by routeID from the AT GTFS API
func (client *Client) TripsByRouteID(routeID string) ([]*Trip, error) {
	url := baseURL + "/v2/gtfs/trips/routeid/" + routeID

	var response []*Trip
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
