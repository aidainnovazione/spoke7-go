package dtos

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"spoke7-go/internal/sumo-integration/models"
	"strconv"
	"strings"

	"github.com/im7mortal/UTM"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type Network struct {
	XMLName     xml.Name     `xml:"net"`
	Location    Location     `xml:"location"`
	Edges       []Edge       `xml:"edge"`
	Connections []Connection `xml:"connection"`
	Junctions   []Junction   `xml:"junction"`
}

type Junction struct {
	ID       string    `xml:"id,attr"`
	Type     string    `xml:"type,attr"`
	X        float64   `xml:"x,attr"`
	Y        float64   `xml:"y,attr"`
	IncLanes string    `xml:"incLanes,attr"`
	IntLanes string    `xml:"intLanes,attr"`
	Shape    string    `xml:"shape,attr"`
	Requests []Request `xml:"request"`
}

type Request struct {
	Index    int    `xml:"index,attr"`
	Response string `xml:"response,attr"`
	Foes     string `xml:"foes,attr"`
	Cont     int    `xml:"cont,attr"`
}

type Location struct {
	NetOffset     string `xml:"netOffset,attr"`
	ConvBoundary  string `xml:"convBoundary,attr"`
	OrigBoundary  string `xml:"origBoundary,attr"`
	ProjParameter string `xml:"projParameter,attr"`
}

type Edge struct {
	ID    string `xml:"id,attr"`
	Lanes []Lane `xml:"lane"`
}

type Lane struct {
	ID     string  `xml:"id,attr"`
	Index  int     `xml:"index,attr"`
	Shape  string  `xml:"shape,attr"`
	Speed  float32 `xml:"speed,attr"`
	Length float32 `xml:"length,attr"`
}

type Connection struct {
	From     string `xml:"from,attr"`
	To       string `xml:"to,attr"`
	FromLane int    `xml:"fromLane,attr"`
	ToLane   int    `xml:"toLane,attr"`
	Dir      string `xml:"dir,attr"`
	State    string `xml:"state,attr"`
}

func GetNetworkProperties(network Network) models.NetworkProperties {
	offsets := strings.Split(network.Location.NetOffset, ",")
	var netOffsetX, netOffsetY float64
	if len(offsets) == 2 {
		netOffsetX, _ = strconv.ParseFloat(offsets[0], 64)
		netOffsetY, _ = strconv.ParseFloat(offsets[1], 64)
	}

	proj, err := extractZone(network.Location.ProjParameter)
	//use default zone as 33
	if err != nil {
		proj = 33
	}

	return models.NetworkProperties{
		NetOffsetX: netOffsetX,
		NetOffsetY: netOffsetY,
		Projection: proj,
	}
}

func NetworkToGeoJSON(network Network) geojson.FeatureCollection {
	// Parse net offset
	offsets := strings.Split(network.Location.NetOffset, ",")
	var netOffsetX, netOffsetY float64
	if len(offsets) == 2 {
		netOffsetX, _ = strconv.ParseFloat(offsets[0], 64)
		netOffsetY, _ = strconv.ParseFloat(offsets[1], 64)
	}

	proj, err := extractZone(network.Location.ProjParameter)
	//use default zone as 33
	if err != nil {
		proj = 33
	}

	// Parse lane shapes
	laneShapes := make(map[string][][2]float64)
	laneMap := make(map[string]Lane)
	for _, edge := range network.Edges {
		for _, lane := range edge.Lanes {
			laneShapes[lane.ID] = parseShape(lane.Shape)
			laneMap[lane.ID] = lane
		}
	}

	// Gather features
	var features []*geojson.Feature
	features = make([]*geojson.Feature, 0)
	features = append(features, laneFeatures(laneMap, laneShapes, proj, netOffsetX, netOffsetY)...)
	features = append(features, connectionFeatures(network.Connections, laneShapes, laneMap, proj, netOffsetX, netOffsetY)...)
	features = append(features, junctionFeatures(network.Junctions, proj, netOffsetX, netOffsetY)...)

	return geojson.FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
}

func laneFeatures(laneMap map[string]Lane, laneShapes map[string][][2]float64, proj int, netOffsetX, netOffsetY float64) []*geojson.Feature {

	features := make([]*geojson.Feature, 0)

	for _, lane := range laneMap {
		shape := laneShapes[lane.ID]
		if len(shape) == 0 {
			continue
		}
		coords := make(orb.LineString, len(shape))
		for i, pt := range shape {
			lat, lon := convertToGeo(pt[0], pt[1], proj, netOffsetX, netOffsetY)
			//coords = append(coords, []float64{lon, lat})
			coords[i] = orb.Point{lon, lat}
		}

		// Convert to geojson feature
		features = append(features, &geojson.Feature{
			Type:     "Feature",
			Geometry: coords,
			Properties: map[string]interface{}{
				"type":   "lane",
				"id":     lane.ID,
				"speed":  lane.Speed,
				"length": lane.Length,
			},
		})
	}
	return features
}

func connectionFeatures(conns []Connection, laneShapes map[string][][2]float64, laneMap map[string]Lane, proj int, netOffsetX, netOffsetY float64) []*geojson.Feature {
	features := make([]*geojson.Feature, 0)

	for _, conn := range conns {
		fromLaneID := conn.From
		toLaneID := conn.To

		fromShape := laneShapes[fromLaneID]
		toShape := laneShapes[toLaneID]

		if len(fromShape) == 0 || len(toShape) == 0 {
			continue
		}

		fromPt := fromShape[len(fromShape)-1]
		toPt := toShape[0]

		lat1, lon1 := convertToGeo(fromPt[0], fromPt[1], proj, netOffsetX, netOffsetY)
		lat2, lon2 := convertToGeo(toPt[0], toPt[1], proj, netOffsetX, netOffsetY)

		coords := orb.LineString{
			orb.Point{lon1, lat1},
			orb.Point{lon2, lat2},
		}

		features = append(features, &geojson.Feature{
			Type:     "Feature",
			Geometry: coords,
			Properties: map[string]interface{}{
				"type":     "connection",
				"from":     conn.From,
				"to":       conn.To,
				"fromLane": conn.FromLane,
				"toLane":   conn.ToLane,
				"dir":      conn.Dir,
				"state":    conn.State,
			},
		})
	}
	return features
}

func junctionFeatures(junctions []Junction, proj int, netOffsetX, netOffsetY float64) []*geojson.Feature {
	features := make([]*geojson.Feature, 0)
	for _, junction := range junctions {
		shape := parseShape(junction.Shape)
		if len(shape) <= 1 {
			continue
		}
		coords := make(orb.LineString, len(shape))
		for i, pt := range shape {
			lat, lon := convertToGeo(pt[0], pt[1], proj, netOffsetX, netOffsetY)
			coords[i] = orb.Point{lon, lat}
		}

		features = append(features, &geojson.Feature{
			Type:     "Feature",
			Geometry: coords,
			Properties: map[string]interface{}{
				"type":          "junction",
				"id":            junction.ID,
				"junction_type": junction.Type,
			},
		})
	}
	return features
}

func parseShape(shapeStr string) [][2]float64 {
	var shape [][2]float64
	points := strings.Split(shapeStr, " ")
	for _, point := range points {
		coords := strings.Split(point, ",")
		if len(coords) == 2 {
			x, err1 := strconv.ParseFloat(coords[0], 64)
			y, err2 := strconv.ParseFloat(coords[1], 64)
			if err1 == nil && err2 == nil {
				shape = append(shape, [2]float64{x, y})
			}
		}
	}
	return shape
}

func convertToGeo(absX, absY float64, proj int, netOffsetX, netOffsetY float64) (float64, float64) {
	realX := absX - netOffsetX
	realY := absY - netOffsetY
	lat, lon, _ := UTM.ToLatLon(realX, realY, proj, "N")
	return lat, lon
}

func extractZone(projStr string) (int, error) {
	re := regexp.MustCompile(`\+zone=(\d+)`)
	match := re.FindStringSubmatch(projStr)
	if len(match) < 2 {
		return 0, fmt.Errorf("zone not found in projection string")
	}
	zone, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, fmt.Errorf("invalid zone value: %v", err)
	}
	return zone, nil
}
