package dao

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/encoding/wkt"
	"github.com/paulmach/orb/geojson"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Geometry struct {
	orb.Geometry
}

func (loc Geometry) GormDataType() string {
	return loc.GeoJSONType()
}

func (loc Geometry) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	switch g := loc.Geometry.(type) {
	case orb.Point:
		return clause.Expr{
			SQL:  "ST_PointFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	case orb.MultiPoint:
		return clause.Expr{
			SQL:  "ST_MultiPointFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	case orb.LineString:
		return clause.Expr{
			SQL:  "ST_LineStringFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	case orb.MultiLineString:
		return clause.Expr{
			SQL:  "ST_MultiLineStringFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	case orb.Polygon:
		return clause.Expr{
			SQL:  "ST_PolygonFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	case orb.MultiPolygon:
		return clause.Expr{
			SQL:  "ST_MultiPolygonFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	case orb.Collection:
		// GeometryCollection is a bit more complex, and might not be supported directly by all databases in the same way.
		// Depending on your database, you may need a custom implementation. For now, we'll return an error.
		return clause.Expr{
			SQL:  "ST_GeomFromText(?)",
			Vars: []interface{}{wkt.MarshalString(g)},
		}
	default:
		return clause.Expr{
			SQL:  "ST_GeomFromText(?)",
			Vars: []interface{}{wkt.MarshalString(loc.Geometry)},
		}
	}
}

// Scan implements the sql.Scanner interface
func (loc *Geometry) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	if v == nil {
		loc.Geometry = nil // or set to a default value if needed
		return nil
	}
	switch val := v.(type) {
	case []byte:
		// Attempt to decode WKB
		geometry, err := wkb.Unmarshal(val)
		if err != nil {
			// Attempt to decode GeoJSON
			g, errGeoJSON := geojson.UnmarshalGeometry(val)
			if errGeoJSON != nil {
				return fmt.Errorf("failed to unmarshal WKB: %w, also failed to unmarshal GeoJSON: %w", err, errGeoJSON)
			}
			loc.Geometry = g.Geometry()
			return nil

		}
		loc.Geometry = geometry
		return nil

	case string:
		// Attempt to decode WKB hex string
		wkbBytes, err := hex.DecodeString(val)
		if err == nil {
			geometry, err := wkb.Unmarshal(wkbBytes)
			if err == nil {
				loc.Geometry = geometry
				return nil
			}
		}

		// Attempt to decode GeoJSON string
		g, errGeoJSON := geojson.UnmarshalGeometry([]byte(val))
		if errGeoJSON == nil {
			loc.Geometry = g.Geometry()
			return nil
		}

		return fmt.Errorf("failed to unmarshal WKB hex string or GeoJSON string: %v", err)

	default:
		return fmt.Errorf("unsupported type %T", v)
	}
}
