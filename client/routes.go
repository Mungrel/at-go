package client

import (
	"errors"
	"net/url"
	"strconv"
)

type routesResponse struct {
	Status   string   `json:"status"`
	Response []*Route `json:"response"`
	Error    *string  `json:"error"`
}

// Route represents a route from the AT GTFS API
type Route struct {
	ID               string  `json:"route_id"`
	AgencyID         string  `json:"agency_id"`
	ShortName        string  `json:"route_short_name"`
	LongName         string  `json:"route_long_name"`
	Description      string  `json:"route_desc"`
	Type             int     `json:"route_type"`
	URL              string  `json:"route_url"`
	Color            string  `json:"route_color"`
	TextColor        string  `json:"route_text_color"`
	StDistanceSphere float64 `json:"st_distance_sphere"`
}

// Routes gets a list of routes from the AT GTFS API
func (client *Client) Routes() ([]*Route, error) {
	url := baseURL + "/v2/gtfs/routes"

	var response routesResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}

// RoutesByLocation gets a list of routes within the radius of a geo point from the AT GTFS API
func (client *Client) RoutesByLocation(latitude, longitude, radius float64) ([]*Route, error) {
	lat := strconv.FormatFloat(latitude, 'f', 6, 64)
	lng := strconv.FormatFloat(longitude, 'f', 6, 64)

	params := url.Values{}
	params.Add("lat", lat)
	params.Add("lng", lng)

	url := baseURL + "/v2/gtfs/routes/geosearch?" + params.Encode()

	var response routesResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
