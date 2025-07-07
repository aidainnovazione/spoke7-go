package dao

import (
	"encoding/json"
	"fmt"
	"spoke7-go/internal/metadata/models"
	"time"

	"github.com/lib/pq"
	"github.com/paulmach/orb/geojson"
	"gorm.io/datatypes"
)

type RoadNetwork struct {
	ID         string         `gorm:"type:uuid;primaryKey"`
	Name       string         `gorm:"type:varchar(255);not null"`
	Geom       datatypes.JSON `gorm:"type:jsonb;default:'[]';not null"`
	Properties datatypes.JSON `gorm:"type:jsonb"`
	Owner      string
	Groups     pq.StringArray `gorm:"type:text[]"`
	CreatedAt  time.Time      `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time      `gorm:"autoUpdateTime:milli"`
}

func NewRoadNetworkDaoFromModel(m *models.RoadNetwork) *RoadNetwork {
	geomBytes, _ := m.Geom.MarshalJSON() // Handle error appropriately in production

	jsonBytes, _ := json.Marshal(m.Properties)

	return &RoadNetwork{
		ID:         m.ID,
		Name:       m.Name,
		Geom:       geomBytes,
		Properties: jsonBytes,
		Owner:      m.Owner,
		Groups:     m.Groups,
		CreatedAt:  m.CreatedAt,
		ModifiedAt: m.ModifiedAt,
	}
}

func NewRoadNetworkModelFromDao(d *RoadNetwork) (*models.RoadNetwork, error) {
	var geojsonFeatureCollection *geojson.FeatureCollection = nil

	var err error
	if len(d.Geom) > 0 {

		geojsonFeatureCollection, err = geojson.UnmarshalFeatureCollection(d.Geom)

		if err != nil {
			geojsonFeatureCollection = nil
		}

	}

	var prop map[string]string
	json.Unmarshal(d.Properties, &prop)

	return &models.RoadNetwork{
		ID:         d.ID,
		Name:       d.Name,
		Geom:       geojsonFeatureCollection,
		Properties: prop,
		Owner:      d.Owner,
		Groups:     d.Groups,
		CreatedAt:  d.CreatedAt,
		ModifiedAt: d.ModifiedAt,
	}, nil
}

func ToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v) // fallback generico
	}
}
