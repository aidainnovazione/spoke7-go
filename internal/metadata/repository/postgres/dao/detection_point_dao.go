package dao

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"log"
	"spoke7-go/internal/metadata/models"
	"time"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

// DetectionPoint represents a point in the road network where a detection device is installed.
type DetectionPoint struct {
	Id             string `gorm:"primaryKey"`
	DataSourceName string `gorm:"primaryKey;index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE,foreignKey:DataSourceName;references:Name"`
	Description    string
	Lanes          []Lane            `gorm:"constraint:OnDelete:CASCADE;foreignKey:DetectionPointId,DataSourceName;references:Id,DataSourceName"`
	Properties     map[string]string `gorm:"-"`
	GeojsonShape   string
	Coordinates    *Coordinate `gorm:"type:geometry(Point,4326)"`
	Position       *Position   `gorm:"type:geometry(Point,4326)"`
	CreatedAt      time.Time   `gorm:"autoCreateTime:milli"`
	ModifiedAt     time.Time   `gorm:"autoUpdateTime:milli"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Scan implements the Scanner interface for GeometryPoint.
func (g *Coordinate) Scan(val interface{}) error {

	if val == nil || val == "" {
		// Optionally return an error or just ignore the value (set default values if necessary)
		// For now, we just return nil to treat it as an empty value
		g = nil
		return nil
	}

	// Ensure the value is a string
	strVal, ok := val.(string)
	if !ok {
		return fmt.Errorf("expected string value, got %T", val)
	}

	// Decode the hex string to bytes
	wkbBytes, err := hex.DecodeString(strVal)
	if err != nil {
		return fmt.Errorf("failed to decode WKB hex string: %w", err)
	}

	// Unmarshal the WKB bytes into a Geometry object
	p, err := wkb.Unmarshal(wkbBytes)
	if err != nil {
		return fmt.Errorf("failed to unmarshal WKB: %w", err)
	}

	// Assert the Geometry type as Point
	if point, ok := p.(orb.Point); ok {
		// Assign coordinates to the struct
		g.Latitude = point[1]
		g.Longitude = point[0]
	} else {
		return fmt.Errorf("expected Point geometry, got %T", p)
	}

	return nil
}

// Value implements the driver Valuer interface for GeometryPoint
func (g Coordinate) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", g.Longitude, g.Latitude), nil
}

// Value implements the driver Valuer interface for GeometryPoint
func (g Position) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", g.X, g.Y), nil
}

// Scan implements the Scanner interface for GeometryPoint
func (g *Position) Scan(val interface{}) error {

	if val == nil || val == "" {
		// Optionally return an error or just ignore the value (set default values if necessary)
		// For now, we just return nil to treat it as an empty value
		g = nil
		return nil
	}

	var p orb.Geometry

	wkbBytes, err := hex.DecodeString(val.(string))
	if err != nil {
		log.Fatal(err)
	}

	p, err = wkb.Unmarshal(wkbBytes)
	if err != nil {
		return err

	}
	// Assert the type as Point and print it
	if p, ok := p.(orb.Point); ok {
		// Output the coordinates of the point
		g.X = p[1]
		g.Y = p[0]

	}

	return nil
}

type Lane struct {
	Id               string `gorm:"primaryKey"`
	DataSourceName   string `gorm:"primaryKey;index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE,foreignKey:DataSourceName;references:Name"`
	DetectionPointId string `gorm:"index"` // ForeignKey to DetectionPoint
	Description      string
	Index            string
	Properties       map[string]string `gorm:"-"`
	GeojsonShape     string
	Coordinates      *Coordinate `gorm:"type:geometry(Point,4326)"`
	Position         *Position   `gorm:"type:geometry(Point,4326)"`
	CreatedAt        time.Time   `gorm:"autoCreateTime:milli"`
	ModifiedAt       time.Time   `gorm:"autoUpdateTime:milli"`
}

func (DetectionPoint) TableName() string {
	return "detection_points"
}

func (Lane) TableName() string {
	return "lanes"
}

// NewDetectionPointDaoFromModel creates a new DetectionPoint DAO from a model.
func FromDetectionPointDaoToModel(dao DetectionPoint) models.DetectionPoint {

	lanes := make([]models.Lane, 0, len(dao.Lanes))
	for _, lane := range dao.Lanes {
		// Convert Lane DAO to Lane model
		lanes = append(lanes, FromLaneDaoToModel(lane))
	}

	var position *models.Position
	var coordinate *models.Coordinate
	if dao.Position != nil {
		position = &models.Position{
			X: dao.Position.X,
			Y: dao.Position.Y,
		}
	}
	if dao.Coordinates != nil {
		coordinate = &models.Coordinate{
			Latitude:  dao.Coordinates.Latitude,
			Longitude: dao.Coordinates.Longitude,
		}
	}

	return models.DetectionPoint{
		Id:           dao.Id,
		Description:  dao.Description,
		Lanes:        lanes,
		Properties:   dao.Properties,
		GeojsonShape: dao.GeojsonShape,
		Coordinates:  coordinate,
		Position:     position,
		CreatedAt:    dao.CreatedAt,
		ModifiedAt:   dao.ModifiedAt,
	}
}

func FromDetectionPointModelToDao(datasourceName string, dp models.DetectionPoint) DetectionPoint {

	lanes := make([]Lane, 0, len(dp.Lanes))
	for _, lane := range dp.Lanes {
		// Convert Lane model to Lane DAO
		lanes = append(lanes, FromLaneModelToDao(lane))
	}

	var position *Position
	if dp.Position != nil {
		position = &Position{
			X: dp.Position.X,
			Y: dp.Position.Y,
		}
	}
	var coordinate *Coordinate
	if dp.Coordinates != nil {
		coordinate = &Coordinate{
			Latitude:  dp.Coordinates.Latitude,
			Longitude: dp.Coordinates.Longitude,
		}
	}

	return DetectionPoint{
		Id:           dp.Id,
		Description:  dp.Description,
		Lanes:        lanes,
		Properties:   dp.Properties,
		GeojsonShape: dp.GeojsonShape,
		Coordinates:  coordinate,

		Position:       position,
		CreatedAt:      dp.CreatedAt,
		ModifiedAt:     dp.ModifiedAt,
		DataSourceName: datasourceName,
	}
}

func FromLaneModelToDao(lane models.Lane) Lane {

	var coordinate *Coordinate
	var position *Position
	if lane.Position != nil {
		position = &Position{
			X: lane.Position.X,
			Y: lane.Position.Y,
		}
	}
	if lane.Coordinates != nil {
		coordinate = &Coordinate{
			Latitude:  lane.Coordinates.Latitude,
			Longitude: lane.Coordinates.Longitude,
		}
	}

	return Lane{
		Id:           lane.Id,
		Description:  lane.Description,
		Index:        lane.Index,
		Properties:   lane.Properties,
		GeojsonShape: lane.GeojsonShape,
		Coordinates:  coordinate,
		Position:     position,

		CreatedAt:  lane.CreatedAt,
		ModifiedAt: lane.ModifiedAt,
	}
}

func FromLaneDaoToModel(dao Lane) models.Lane {

	var position *models.Position
	if dao.Position != nil {
		position = &models.Position{
			X: dao.Position.X,
			Y: dao.Position.Y,
		}
	}
	var coordinate *models.Coordinate
	if dao.Coordinates != nil {
		coordinate = &models.Coordinate{
			Latitude:  dao.Coordinates.Latitude,
			Longitude: dao.Coordinates.Longitude,
		}
	}

	return models.Lane{
		Id:           dao.Id,
		Description:  dao.Description,
		Index:        dao.Index,
		Properties:   dao.Properties,
		GeojsonShape: dao.GeojsonShape,
		Coordinates:  coordinate,
		Position:     position,
		CreatedAt:    dao.CreatedAt,
		ModifiedAt:   dao.ModifiedAt,
	}
}
