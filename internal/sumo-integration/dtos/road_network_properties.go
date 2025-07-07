package dtos

import (
	"spoke7-go/internal/sumo-integration/models"
	"strconv"
)

func ToRoadNetworkPropertiesDTOs(prop models.NetworkProperties) map[string]string {
	return map[string]string{
		"offsetX":    strconv.FormatFloat(prop.NetOffsetX, 'f', -1, 64),
		"offsetY":    strconv.FormatFloat(prop.NetOffsetY, 'f', -1, 64),
		"projection": strconv.Itoa(prop.Projection),
	}
}
