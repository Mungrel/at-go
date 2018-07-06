package client

import (
	"net/url"
	"strconv"
)

// Stop represents a stop in the AT GTFS API
type Stop struct {
	ID                 string  `json:"stop_id"`
	Name               string  `json:"stop_name"`
	Description        string  `json:"stop_desc"`
	Latitude           float64 `json:"stop_lat"`
	Longitude          float64 `json:"stop_lon"`
	ZoneID             string  `json:"zone_id"`
	URL                string  `json:"stop_url"`
	Code               string  `json:"stop_code"`
	Street             string  `json:"stop_street"`
	City               string  `json:"stop_city"`
	Region             string  `json:"stop_region"`
	PostCode           string  `json:"stop_postcode"`
	Country            string  `json:"stop_country"`
	LocationType       int     `json:"location_type"`
	ParentStation      string  `json:"parent_station"`
	Timezone           string  `json:"stop_timezone"`
	WheelchairBoarding string  `json:"wheelchair_boarding"`
	Direction          string  `json:"direction"`
	Position           string  `json:"position"`
	Geom               string  `json:"the_geom"`
	StDistanceSphere   float64 `json:"st_distance_sphere"`
}

// Stops gets a list of stops from the AT GTFS API
func (client *Client) Stops() ([]*Stop, error) {
	url := baseURL + "/v2/gtfs/stops"

	var response []*Stop
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// StopsByLocation gets a list of stops within the radius of a point from the AT GTFS API
func (client *Client) StopsByLocation(latitude, longitude, radius float64) ([]*Stop, error) {
	params := url.Values{}

	params.Add("lat", strconv.FormatFloat(latitude, 'f', 6, 64))
	params.Add("lng", strconv.FormatFloat(longitude, 'f', 6, 64))
	params.Add("distance", strconv.FormatFloat(radius, 'f', 6, 64))

	url := baseURL + "/v2/gtfs/stops/geosearch?" + params.Encode()

	var response []*Stop
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// StopsSearch searches the stops list by stop name
func (client *Client) StopsSearch(searchTerm string) ([]*Stop, error) {
	url := baseURL + "/v2/gtfs/stops/search/" + searchTerm

	var response []*Stop
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// StopInfo represents a stop info object in the AT GTFS API
type StopInfo struct {
	TripID         string `json:"trip_id"`
	DepartureTime  string `json:"departure_time"`
	TripShortName  string `json:"trip_short_name"`
	TripHeadsign   string `json:"trip_headsign"`
	RouteLongName  string `json:"route_long_name"`
	RouteShortName string `json:"route_short_name"`
	StopSequence   int    `json:"stop_sequence"`
	PickupType     int    `json:"pickup_type"`
	DropOffType    int    `json:"drop_off_type"`
}

// StopInfoByCode gets a list of stop info objects, filtered by stop code
func (client *Client) StopInfoByCode(code string) ([]*StopInfo, error) {
	url := baseURL + "/v2/gtfs/stops/stopinfo/" + code

	var response []*StopInfo
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
