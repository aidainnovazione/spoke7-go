package dtos

import (
	"encoding/xml"
	"fmt"
	"math"
	"spoke7-go/internal/metadata/models"
	"time"
)

// Additional represents the root XML element <additional>
type Additional struct {
	XMLName        xml.Name        `xml:"additional"`
	InductionLoops []InductionLoop `xml:"inductionLoop"`
}

// InductionLoop represents each <inductionLoop> element
type InductionLoop struct {
	ID     string  `xml:"id,attr"`
	Lane   string  `xml:"lane,attr"`
	Pos    float64 `xml:"pos,attr"`
	Period float64 `xml:"period,attr"`
	File   string  `xml:"file,attr"`
}

// extract induction loops from the additional file
func GetInductionLoops(network Network, additional Additional) []models.DetectionPoint {
	// Build a map from lane ID to shape

	//1 get network properties
	networkProp := GetNetworkProperties(network)

	//1. Parse the network to get lane shapes, this will be used to calculate the absolute position of the induction loops
	laneShapeMap := make(map[string][][2]float64)
	laneById := make(map[string]Lane)
	EdgeByLaneId := make(map[string]Edge)
	for _, edge := range network.Edges {
		for _, lane := range edge.Lanes {
			laneShapeMap[lane.ID] = parseShape(lane.Shape)
			laneById[lane.ID] = lane
			EdgeByLaneId[lane.ID] = edge
		}
	}

	//var loops []models.Lane
	detectionPoints := make(map[string]models.DetectionPoint)

	//2. get induction loops from the additional file
	for _, loop := range additional.InductionLoops {

		//2.1 compute the absolute position of the induction loop
		//2.1.1 get the lane shape from the lane ID
		shape, exists := laneShapeMap[loop.Lane]
		lane := laneById[loop.Lane]
		edge := EdgeByLaneId[loop.Lane]

		if !exists || len(shape) == 0 {
			continue // Skip if no shape is found for the lane
		}

		absX, absY := computePosition(shape, loop.Pos)
		lat, lon := convertToGeo(absX, absY, networkProp.Projection, networkProp.NetOffsetX, networkProp.NetOffsetY)

		detectionlane := models.Lane{
			Id:          loop.ID,
			Description: fmt.Sprintf("Induction Loop at position %d of lane %d ", loop.Pos, lane.Index),
			Index:       fmt.Sprintf("%d", lane.Index),
			Properties: map[string]string{
				"lane_id": loop.Lane,
				"period":  fmt.Sprintf("%f", loop.Period),
				"file":    loop.File,
			},
			Coordinates: &models.Coordinate{
				Latitude:  lat,
				Longitude: lon,
			},
			Position: &models.Position{
				X: absX,
				Y: absY,
			},
			CreatedAt:  time.Time{},
			ModifiedAt: time.Time{},
		}

		//check if detection point already exists, if not, create it
		detectionPoint, exists := detectionPoints[edge.ID]
		if !exists {
			detectionPoint = models.DetectionPoint{
				Id:          edge.ID,
				Description: fmt.Sprintf("Detection Point at edge %s", edge.ID),
				Lanes:       []models.Lane{},
				Properties:  map[string]string{},
				Coordinates: &models.Coordinate{
					Latitude:  lat,
					Longitude: lon,
				},
				Position: &models.Position{
					X: absX,
					Y: absY,
				},
				CreatedAt:  time.Time{},
				ModifiedAt: time.Time{},
			}
			detectionPoints[edge.ID] = detectionPoint
		}

		//2.2 add the induction loop to the detection point
		detectionPoint.Lanes = append(detectionPoint.Lanes, detectionlane)
		detectionPoints[edge.ID] = detectionPoint

	}

	//3. convert the detection points to a slice
	detectionPointsSlice := make([]models.DetectionPoint, 0, len(detectionPoints))
	for _, detectionPoint := range detectionPoints {
		detectionPointsSlice = append(detectionPointsSlice, detectionPoint)
	}

	return detectionPointsSlice
}

// Given a share and a position on the shape, compute the absolute position of the induction loop
func computePosition(shape [][2]float64, pos float64) (float64, float64) {
	var accumulated float64
	for i := 1; i < len(shape); i++ {
		dx := shape[i][0] - shape[i-1][0]
		dy := shape[i][1] - shape[i-1][1]
		dist := math.Sqrt(dx*dx + dy*dy)
		if accumulated+dist >= pos {
			ratio := (pos - accumulated) / dist
			x := shape[i-1][0] + ratio*dx
			y := shape[i-1][1] + ratio*dy
			return x, y
		}
		accumulated += dist
	}
	if len(shape) > 0 {
		return shape[len(shape)-1][0], shape[len(shape)-1][1]
	}
	return 0, 0
}
