package models

import (
	"time"

	"github.com/paulmach/orb/geojson"
)

type RoadNetwork struct {
	ID         string
	Name       string
	Geom       *geojson.FeatureCollection
	Properties map[string]string
	Owner      string
	Groups     []string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
