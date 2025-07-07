package models

import (
	"time"
)

type HistoryHourTrafficDataByDetectionPointByLaneModel struct {
	CreatedAt          time.Time
	ModifiedAt         time.Time
	DataSourceName     string
	LaneID             string
	DetectionTimestamp time.Time
	DetectionInterval  uint32

	TrafficFlowVehicleClass1          float32
	TrafficFlowVehicleClass2          float32
	TrafficFlowVehicleClass3          float32
	TrafficFlowVehicleClass4          float32
	TrafficFlowVehicleClass5          float32
	TrafficFlowVehicleClass6          float32
	TrafficFlowVehicleClass7          float32
	TrafficFlowVehicleClass8          float32
	TrafficFlowVehicleClassEquivalent float32

	AverageSpeedVehicleClass1   float32
	AverageSpeedVehicleClass2   float32
	AverageSpeedVehicleClass3   float32
	AverageSpeedVehicleClass4   float32
	AverageSpeedVehicleClass5   float32
	AverageSpeedVehicleClass6   float32
	AverageSpeedVehicleClass7   float32
	AverageSpeedVehicleClass8   float32
	AverageSpeedVehicleClassAll float32

	AverageVehicleLength   float32
	AverageHeadway         float32
	AverageTimeToCollision float32
}

func (m HistoryHourTrafficDataByDetectionPointByLaneModel) GetDetectionTimestamp() time.Time {
	return m.DetectionTimestamp
}
