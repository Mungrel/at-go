package test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	at "github.com/Mungrel/at-go/client"
)

const (
	mockServiceID       = "123456ABC"
	mockRouteID         = "98302-20180524131340_v66.89"
	mockTripID          = "98302-20180524131340_v66.89"
	mockShapeID         = "81-20180605110613_v67.1"
	mockShapeGeomID     = "813-20180524131340_v66.89"
	mockStopsSearchTerm = "Papakura"
	mockRoutesLongName  = "Papakura Train Station to Britomart Train Station"
)

type mockClient struct{}

func newMockClient() *at.Client {
	return &at.Client{
		APIKey: "dummy-key",
		Client: &mockClient{},
	}
}

func (mc *mockClient) Do(req *http.Request) (*http.Response, error) {
	url := req.URL.Path

	if strings.HasSuffix(url, "/v2/gtfs/agency") {
		return newOKResponse(getJSON(agencies)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendar") {
		return newOKResponse(getJSON(calendars)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendar/serviceId/"+mockServiceID) {
		return newOKResponse(getJSON(calendarByService)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendarDate") {
		return newOKResponse(getJSON(calendarDates)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/calendarDate/serviceId/"+mockServiceID) {
		return newOKResponse(getJSON(calendarDatesByService)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/metadata") {
		return newOKResponse(getJSON(metadata)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes") {
		return newOKResponse(getJSON(routes)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/geosearch") {
		return newOKResponse(getJSON(routesByLocation)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops") {
		return newOKResponse(getJSON(stops)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/trips") {
		return newOKResponse(getJSON(trips)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/versions") {
		return newOKResponse(getJSON(versions)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/trips/tripId/"+mockTripID) {
		return newOKResponse(getJSON(tripsByID)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/shapes/shapeId/"+mockShapeID) {
		return newOKResponse(getJSON(shapesByID)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/trips/routeid/"+mockRouteID) {
		return newOKResponse(getJSON(tripsByRouteID)), nil
	}

	if strings.HasSuffix(url, "/v2/locations/customerservicecentres") {
		return newOKResponse(getJSON(customerServiceCentres)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/shapes/geometry/"+mockShapeGeomID) {
		return newOKResponse(getJSON(shapeGeometryByID)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops/geosearch") {
		return newOKResponse(getJSON(stopsByLocation)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops/search/"+mockStopsSearchTerm) {
		return newOKResponse(getJSON(stopsSearch)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/routeLongName/"+mockRoutesLongName) {
		return newOKResponse(getJSON(routesByLongName)), nil
	}

	return nil, errors.New("Endpoint not mocked")
}

func newOKResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}
