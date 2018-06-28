package client

import (
	"errors"
	"net/url"
	"strconv"
)

type stopsResponse struct {
	Status   string  `json:"status"`
	Response []*Stop `json:"response"`
	Error    *string `json:"error"`
}

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

	var response stopsResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}

// StopsByLocation gets a list of stops within the radius of a point from the AT GTFS API
func (client *Client) StopsByLocation(latitude, longitude, radius float64) ([]*Stop, error) {
	params := url.Values{}

	params.Add("lat", strconv.FormatFloat(latitude, 'f', 6, 64))
	params.Add("lng", strconv.FormatFloat(longitude, 'f', 6, 64))
	params.Add("distance", strconv.FormatFloat(radius, 'f', 6, 64))

	url := baseURL + "/v2/gtfs/stops/geosearch?" + params.Encode()

	var response stopsResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
