package client

import "net/url"

type vehiclePositionsResponse struct {
	Entity []*vehiclePositionEntity `json:"entity"`
}

type vehiclePositionEntity struct {
	ID        string   `json:"id"`
	IsDeleted bool     `json:"is_deleted"`
	Vehicle   *vehicle `json:"vehicle"`
}

type vehicle struct {
	Trip            *trip        `json:"trip"`
	Vehicle         *tripVehicle `json:"vehicle"`
	Position        *Position    `json:"position"`
	Timestamp       float64      `json:"timestamp"`
	OccupancyStatus int64        `json:"occupancy_status"`
}

type trip struct {
	TripID               string `json:"trip_id"`
	RouteID              string `json:"route_id"`
	StartTime            string `json:"start_time"`
	ScheduleRelationship int64  `json:"schedule_relationship"`
}

type tripVehicle struct {
	ID string `json:"id"`
}

// Position holds lat, long and a bearing
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Bearing   string  `json:"bearing"`
}

// VehiclePosition is a convenient representation of the vehicle position update that comes back from the API
type VehiclePosition struct {
	ID                   string
	IsDeleted            bool
	TripID               string
	RouteID              string
	StartTime            string
	ScheduleRelationship int64
	VehicleID            string
	Position             *Position
	Timestamp            float64
	OccupancyStatus      int64
}

func (vpe *vehiclePositionEntity) toVehiclePosition() *VehiclePosition {
	return &VehiclePosition{
		ID:                   vpe.ID,
		IsDeleted:            vpe.IsDeleted,
		TripID:               vpe.Vehicle.Trip.TripID,
		RouteID:              vpe.Vehicle.Trip.RouteID,
		StartTime:            vpe.Vehicle.Trip.StartTime,
		ScheduleRelationship: vpe.Vehicle.Trip.ScheduleRelationship,
		VehicleID:            vpe.Vehicle.Vehicle.ID,
		Position:             vpe.Vehicle.Position,
		Timestamp:            vpe.Vehicle.Timestamp,
		OccupancyStatus:      vpe.Vehicle.OccupancyStatus,
	}
}

// VehiclePositions gets a list of vehicle positions from the AT Realtime transit feed
func (client *Client) VehiclePositions(tripID, vehicleID *string) ([]*VehiclePosition, error) {
	params := url.Values{}

	if tripID != nil {
		params.Add("tripid", *tripID)
	}

	if vehicleID != nil {
		params.Add("vehicleid", *vehicleID)
	}

	url := baseURL + "/v2/public/realtime/vehiclelocations?" + params.Encode()

	var response vehiclePositionsResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	vehiclePositions := make([]*VehiclePosition, 0, len(response.Entity))

	for _, entity := range response.Entity {
		vehiclePositions = append(vehiclePositions, entity.toVehiclePosition())
	}

	return vehiclePositions, nil
}
