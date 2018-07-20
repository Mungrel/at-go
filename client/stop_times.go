package client

// StopTime represents a stop time in the AT GTFS API
type StopTime struct {
	TripID                 string  `json:"trip_id"`
	ArrivalTime            string  `json:"arrival_time"`
	DepartureTime          string  `json:"departure_time"`
	StopID                 string  `json:"stop_id"`
	StopSequence           int     `json:"stop_sequence"`
	StopHeadsign           string  `json:"stop_headsign"`
	PickupType             int     `json:"pickup_type"`
	DropoffType            int     `json:"drop_off_type"`
	ShapeDistanceTravelled float64 `json:"shape_dist_traveled"`
	Timepoint              string  `json:"timepoint"`
	ArrivalTimeSeconds     int64   `json:"arrival_time_seconds"`
	DepartureTimeSeconds   int64   `json:"departure_time_seconds"`
}

// StopTimesByID gets a list of stop times, filtered by stop ID, from the AT GTFS API
func (client *Client) StopTimesByID(stopID string) ([]*StopTime, error) {
	url := baseURL + "/v2/gtfs/stopTimes/stopId/" + stopID

	var response []*StopTime
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// StopTimesByTripID gets a list of stop times, filtered by trip ID, from the AT GTFS API
func (client *Client) StopTimesByTripID(tripID string) ([]*StopTime, error) {
	url := baseURL + "/v2/gtfs/stopTimes/tripId/" + tripID

	var response []*StopTime
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
