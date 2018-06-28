package client

// ShapeGeom represents a shape geometry from the AT GTFS API
type ShapeGeom struct {
	ShapeID string `json:"shape_id"`
	Geom    string `json:"the_geom"`
}

// ShapeGeometryByID gets a list of shape geometries filtered by shape ID
func (client *Client) ShapeGeometryByID(shapeID string) ([]*ShapeGeom, error) {
	url := baseURL + "/v2/gtfs/shapes/geometry/" + shapeID

	var response []*ShapeGeom
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
