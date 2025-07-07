package models

import "time"

// DetectionPoint(Sezione) represents a point in the road network where a detection device is installed.
type DetectionPoint struct {
	Id           string
	Description  string
	Lanes        []Lane
	Properties   map[string]string // Additional metadata properties as key-value pairs
	GeojsonShape string            // GeoJSON representation (Point, LineString, Polygon, etc.)
	Coordinates  *Coordinate       // Represents the central point
	Position     *Position         // Represents the central point with x and y coordinates
	CreatedAt    time.Time
	ModifiedAt   time.Time
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Lane struct {
	Id           string
	Description  string
	Index        string            // Maps to "id_of_lane" from JSON
	Properties   map[string]string // Additional metadata properties
	GeojsonShape string            // GeoJSON representation (Point, LineString, etc.)
	Coordinates  *Coordinate       // Represents the lane's coordinates
	Position     *Position         // Represents optional x and y position
	CreatedAt    time.Time
	ModifiedAt   time.Time
}
