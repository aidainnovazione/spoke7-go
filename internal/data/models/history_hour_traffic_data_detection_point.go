package models

import (
	"time"
)

type HistoryHourTrafficDataByDetectionPointModel struct {
	CreatedAt  time.Time
	ModifiedAt time.Time

	DataSourceName     string
	DetectionPointID   string
	DetectionTimestamp time.Time
	DetectionInterval  uint32

	// flusso - a partire dai conteggi
	TrafficFlowVehicleClass1 float32
	TrafficFlowVehicleClass2 float32
	TrafficFlowVehicleClass3 float32
	TrafficFlowVehicleClass4 float32
	TrafficFlowVehicleClass5 float32
	TrafficFlowVehicleClass6 float32
	TrafficFlowVehicleClass7 float32
	TrafficFlowVehicleClass8 float32

	TrafficFlowVehicleClassEquivalent float32

	// velocità di transito

	// - media
	AverageSpeedVehicleClass1 float32
	AverageSpeedVehicleClass2 float32
	AverageSpeedVehicleClass3 float32
	AverageSpeedVehicleClass4 float32
	AverageSpeedVehicleClass5 float32
	AverageSpeedVehicleClass6 float32
	AverageSpeedVehicleClass7 float32
	AverageSpeedVehicleClass8 float32

	AverageSpeedVehicleClassAll float32

	// lunghezza
	AverageVehicleLength float32

	// headway fra veicoli
	AverageHeadway float32

	// time-to-collision
	AverageTimeToCollision float32
}

func (m HistoryHourTrafficDataByDetectionPointModel) GetDetectionTimestamp() time.Time {
	return m.DetectionTimestamp
}
