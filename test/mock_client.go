package test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	at "github.com/Mungrel/at-go/client"
)

const (
	mockServiceID            = "123456ABC"
	mockRouteID              = "98302-20180524131340_v66.89"
	mockTripID               = "98302-20180524131340_v66.89"
	mockShapeID              = "81-20180605110613_v67.1"
	mockShapeGeomID          = "813-20180524131340_v66.89"
	mockShapeTripID          = "2365121452-20180524131340_v66.89"
	mockSearchTerm           = "Papakura"
	mockRoutesLongName       = "Papakura Train Station to Britomart Train Station"
	mockRoutesShortName      = "STH"
	mockNotificationCategory = "MOVED_STOP"
	mockStopID               = "7004"
	mockStopTripID           = "2365121452-20180524131340_v66.89"
	mockStopCode             = "2716"
	mockStopTimeStopID       = "0097-20180524131340_v66.89"
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

	if strings.HasSuffix(url, "/v2/gtfs/stops/search/"+mockSearchTerm) {
		return newOKResponse(getJSON(stopsSearch)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/routeLongName/"+mockRoutesLongName) {
		return newOKResponse(getJSON(routesByLongName)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/routeShortName/"+mockRoutesShortName) {
		return newOKResponse(getJSON(routesByShortName)), nil
	}

	if strings.HasSuffix(url, "/v2/locations/parkinglocations") {
		return newOKResponse(getJSON(parkingLocations)), nil
	}

	if strings.HasSuffix(url, "/v2/locations/scheduledworks") {
		return newOKResponse(getJSON(scheduledWorks)), nil
	}

	if strings.HasSuffix(url, "/v2/notifications/") {
		return newOKResponse(getJSON(notifications)), nil
	}

	if strings.HasSuffix(url, "/v2/notifications/"+mockNotificationCategory) {
		return newOKResponse(getJSON(notificationsByCategory)), nil
	}

	if strings.HasSuffix(url, "/v2/notifications/stop/"+mockStopID) {
		return newOKResponse(getJSON(notificationsByStop)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops/stopinfo/"+mockStopCode) {
		return newOKResponse(getJSON(stopInfoByCode)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stopTimes/stopId/"+mockStopTimeStopID) {
		return newOKResponse(getJSON(stopTimesByID)), nil
	}

	if strings.HasSuffix(url, "/v2/public/realtime/vehiclelocations") {
		return newOKResponse(getJSON(vehiclePositions)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/stopid/"+mockStopID) {
		return newOKResponse(getJSON(routesByStopID)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/routeId/"+mockRouteID) {
		return newOKResponse(getJSON(routesByID)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/routes/search/"+mockSearchTerm) {
		return newOKResponse(getJSON(routesSearch)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/shapes/tripId/"+mockShapeTripID) {
		return newOKResponse(getJSON(shapesByTrip)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops/stopCode/"+mockStopCode) {
		return newOKResponse(getJSON(stopByCode)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops/stopId/"+mockStopTimeStopID) {
		return newOKResponse(getJSON(stopsByID)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stops/tripId/"+mockStopTripID) {
		return newOKResponse(getJSON(stopsByTrip)), nil
	}

	if strings.HasSuffix(url, "/v2/gtfs/stopTimes/tripId/"+mockStopTripID) {
		return newOKResponse(getJSON(stopTimesByTrip)), nil
	}

	return nil, errors.New("Endpoint not mocked")
}

func newOKResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}
