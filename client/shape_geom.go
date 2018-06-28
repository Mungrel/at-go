package client

import "errors"

type shapeGeomResponse struct {
	Status   string       `json:"status"`
	Response []*ShapeGeom `json:"response"`
	Error    *string      `json:"error"`
}

// ShapeGeom represents a shape geometry from the AT GTFS API
type ShapeGeom struct {
	ShapeID string `json:"shape_id"`
	Geom    string `json:"the_geom"`
}

// ShapeGeometryByID gets a list of shape geometries filtered by shape ID
func (client *Client) ShapeGeometryByID(shapeID string) ([]*ShapeGeom, error) {
	url := baseURL + "/v2/gtfs/shapes/geometry/" + shapeID

	var response shapeGeomResponse
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(*response.Error)
	}

	return response.Response, nil
}
