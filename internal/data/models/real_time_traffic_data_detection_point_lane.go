package models

import (
	"time"
)

type VehicleClassType uint8

const (
	VEHICLE_CLASS_1     VehicleClassType = iota // (moto)
	VEHICLE_CLASS_2                             // (auto)
	VEHICLE_CLASS_3                             // (auto con rimorchio)
	VEHICLE_CLASS_4                             // (furgone)
	VEHICLE_CLASS_5                             // (camion)
	VEHICLE_CLASS_6                             // (autotreno)
	VEHICLE_CLASS_7                             // (autoarticolato)
	VEHICLE_CLASS_8                             // (autobus)
	VEHICLE_CLASS_OTHER                         // (non classificato)
)

type RealTimeTrafficDataByDetectionPointByLaneModel struct {
	CreatedAt            time.Time
	ModifiedAt           time.Time
	DataSourceName       string
	LaneID               string
	DetectionTimestamp   time.Time
	DetectionType        string
	DetectionTechnology  string
	VehicleClass         VehicleClassType
	VehicleSpeed         float32
	VehicleLength        float32
	VehicleHeadway       float32
	QueuePresent         bool
	CorrectFlowDirection bool
}

func (m RealTimeTrafficDataByDetectionPointByLaneModel) GetDetectionTimestamp() time.Time {
	return m.DetectionTimestamp
}
