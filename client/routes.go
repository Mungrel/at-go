package client

import (
	"net/url"
	"strconv"
)

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

	var response []*Route
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// RoutesByLocation gets a list of routes within the radius of a geo point from the AT GTFS API
func (client *Client) RoutesByLocation(latitude, longitude, radius float64) ([]*Route, error) {
	params := url.Values{}

	params.Add("lat", strconv.FormatFloat(latitude, 'f', 6, 64))
	params.Add("lng", strconv.FormatFloat(longitude, 'f', 6, 64))
	params.Add("distance", strconv.FormatFloat(radius, 'f', 6, 64))

	url := baseURL + "/v2/gtfs/routes/geosearch?" + params.Encode()

	var response []*Route
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// RoutesByLongName gets a list of routes by long name from the AT GTFS API
func (client *Client) RoutesByLongName(longName string) ([]*Route, error) {
	url := baseURL + "/v2/gtfs/routes/routeLongName/" + longName

	var response []*Route
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// RoutesByShortName gets a list of routes by short name from the AT GTFS API
func (client *Client) RoutesByShortName(shortName string) ([]*Route, error) {
	url := baseURL + "/v2/gtfs/routes/routeShortName/" + shortName

	var response []*Route
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// RoutesByStopID gets a list of routes by stop ID from the AT GTFS API
func (client *Client) RoutesByStopID(stopID string) ([]*Route, error) {
	url := baseURL + "/v2/gtfs/routes/stopid/" + stopID

	var response []*Route
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// RoutesByID gets a list of routes by route ID from the AT GTFS API
func (client *Client) RoutesByID(routeID string) ([]*Route, error) {
	url := baseURL + "/v2/gtfs/routes/routeId/" + routeID

	var response []*Route
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
