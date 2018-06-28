package client

// Shape represents a Shape in the AT GTFS API
type Shape struct {
	ID                string  `json:"shape_id"`
	PointLat          float64 `json:"shape_pt_lat"`
	PointLong         float64 `json:"shape_pt_lon"`
	PointSequence     int64   `json:"shape_pt_sequence"`
	DistanceTravelled float64 `json:"shape_dist_traveled"`
}

// ShapesByID gets a list of Shapes, filtered by ID from the AT GTFS API
func (client *Client) ShapesByID(shapeID string) ([]*Shape, error) {
	url := baseURL + "/v2/gtfs/shapes/shapeId/" + shapeID

	var response []*Shape
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
