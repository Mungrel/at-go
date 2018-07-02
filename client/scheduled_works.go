package client

import (
	"encoding/json"
)

// ScheduledWorks represents an instance of scheduled works in the AT Locations API
type ScheduledWorks struct {
	ID                string     `json:"id"`
	Code              string     `json:"code"`
	Type              string     `json:"type"`
	Name              string     `json:"name"`
	Address           string     `json:"address"`
	Suburb            string     `json:"suburb"`
	Region            string     `json:"region"`
	Description       string     `json:"description"`
	StartDate         *Timestamp `json:"startDate"`
	EndDate           *Timestamp `json:"endDate"`
	LastUpdated       *Timestamp `json:"lastUpdated"`
	ContractorCompany string     `json:"contractorCompany"`
	Latitude          float64    `json:"latitude"`
	Longitude         float64    `json:"longitude"`
	Geometry          *Geometry  `json:"geometry"`
}

// Geometry wraps the type and coordinates of the works' geometry
type Geometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
	Encoded     bool        `json:"encoded"`
}

// UnmarshalJSON unmarshals into consistent 2D array of floats, even though the API may give us a 1D or 2D array
func (g *Geometry) UnmarshalJSON(b []byte) error {
	var v1 struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
		Encoded     bool      `json:"encoded"`
	}
	var v2 struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
		Encoded     bool        `json:"encoded"`
	}

	err := json.Unmarshal(b, &v1)
	if err != nil {
		err = json.Unmarshal(b, &v2)
		if err != nil {
			return err
		}
		// It's a 2D array / slice of points
		g.Coordinates = v2.Coordinates
		g.Type = v2.Type
		g.Encoded = v2.Encoded
	} else {
		// It's a 1D array / single point
		g.Coordinates = [][]float64{v1.Coordinates}
		g.Type = v1.Type
		g.Encoded = v1.Encoded
	}

	return err
}

// ScheduledWorks gets a list of scheduled works from the AT Locations API
func (client *Client) ScheduledWorks() ([]*ScheduledWorks, error) {
	url := baseURL + "/v2/locations/scheduledworks"

	var response []*ScheduledWorks
	err := client.get(url, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
