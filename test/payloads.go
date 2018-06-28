package test

import "io/ioutil"

const (
	partialPath            = "payloads/"
	agencies               = "agencies.json"
	calendars              = "calendars.json"
	calendarByService      = "calendar_by_service.json"
	calendarDates          = "calendar_dates.json"
	calendarDatesByService = "calendar_dates_by_service.json"
	metadata               = "metadata.json"
	routes                 = "routes.json"
	routesByLocation       = "routes_by_location.json"
	stops                  = "stops.json"
	shapesByID             = "shapes_by_id.json"
	shapeGeometryByID      = "shape_geometry_by_id.json"
	trips                  = "trips.json"
	tripsByID              = "trips_by_id.json"
	tripsByRouteID         = "trips_by_route_id.json"
	versions               = "versions.json"
	customerServiceCentres = "customer_service_centres.json"
)

func getJSON(fileName string) string {
	bytes, err := ioutil.ReadFile(partialPath + fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
