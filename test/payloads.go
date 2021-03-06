package test

import "io/ioutil"

const (
	partialPath                = "payloads/"
	agencies                   = "agencies.json"
	calendars                  = "calendars.json"
	calendarByService          = "calendar_by_service.json"
	calendarDates              = "calendar_dates.json"
	calendarDatesByService     = "calendar_dates_by_service.json"
	metadata                   = "metadata.json"
	routes                     = "routes.json"
	routesByLocation           = "routes_by_location.json"
	routesByLongName           = "routes_by_long_name.json"
	routesByShortName          = "routes_by_short_name.json"
	routesByStopID             = "routes_by_stop_id.json"
	routesByID                 = "routes_by_id.json"
	routesSearch               = "routes_search.json"
	stops                      = "stops.json"
	stopsByID                  = "stops_by_id.json"
	stopsByLocation            = "stops_by_location.json"
	stopsByTrip                = "stops_by_trip.json"
	stopsSearch                = "stops_search.json"
	stopInfoByCode             = "stop_info_by_code.json"
	stopByCode                 = "stops_by_code.json"
	stopTimesByID              = "stop_times_by_id.json"
	stopTimesByTrip            = "stop_times_by_trip.json"
	stopTimesByTripAndSequence = "stop_times_by_trip_and_sequence.json"
	shapesByID                 = "shapes_by_id.json"
	shapeGeometryByID          = "shape_geometry_by_id.json"
	shapesByTrip               = "shapes_by_trip.json"
	trips                      = "trips.json"
	tripsByID                  = "trips_by_id.json"
	tripsByRouteID             = "trips_by_route_id.json"
	versions                   = "versions.json"
	customerServiceCentres     = "customer_service_centres.json"
	parkingLocations           = "parking_locations.json"
	scheduledWorks             = "scheduled_works.json"
	notifications              = "notifications.json"
	notificationsByCategory    = "notifications_by_category.json"
	notificationsByStop        = "notifications_by_stop.json"
	vehiclePositions           = "vehicle_positions.json"
)

func getJSON(fileName string) string {
	bytes, err := ioutil.ReadFile(partialPath + fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
