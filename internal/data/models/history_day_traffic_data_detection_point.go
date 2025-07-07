package models

import (
	"time"
)

type HistoryDayTrafficDataByDetectionPointModel struct {
	CreatedAt  time.Time
	ModifiedAt time.Time

	DataSourceName   string
	DetectionPointID string

	DetectionTimestamp time.Time
	DetectionInterval  uint32

	TrafficFlowVehicleClass1 float32
	TrafficFlowVehicleClass2 float32
	TrafficFlowVehicleClass3 float32
	TrafficFlowVehicleClass4 float32
	TrafficFlowVehicleClass5 float32
	TrafficFlowVehicleClass6 float32
	TrafficFlowVehicleClass7 float32
	TrafficFlowVehicleClass8 float32

	TrafficFlowVehicleClassEquivalent float32

	AverageSpeedVehicleClass1 float32
	AverageSpeedVehicleClass2 float32
	AverageSpeedVehicleClass3 float32
	AverageSpeedVehicleClass4 float32
	AverageSpeedVehicleClass5 float32
	AverageSpeedVehicleClass6 float32
	AverageSpeedVehicleClass7 float32
	AverageSpeedVehicleClass8 float32

	AverageSpeedVehicleClassAll float32
	AverageVehicleLength        float32
	AverageHeadway              float32
	AverageTimeToCollision      float32

	TrafficFlowParametersVehicleClass1 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass2 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass3 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass4 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass5 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass6 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass7 TrafficFlowParametersByDay
	TrafficFlowParametersVehicleClass8 TrafficFlowParametersByDay

	InstantaneousSpeedVelClass1MaxFlow float32
	InstantaneousSpeedVelClass2MaxFlow float32
	InstantaneousSpeedVelClass3MaxFlow float32

	InstantaneousSpeedVelClass1MinFlow float32
	InstantaneousSpeedVelClass2MinFlow float32
	InstantaneousSpeedVelClass3MinFlow float32

	HeadwayVelClass1MaxFlow float32
	HeadwayVelClass2MaxFlow float32
	HeadwayVelClass3MaxFlow float32

	HeadwayVelClass1MinFlow float32
	HeadwayVelClass2MinFlow float32
	HeadwayVelClass3MinFlow float32
}

func (m HistoryDayTrafficDataByDetectionPointModel) GetDetectionTimestamp() time.Time {
	return m.DetectionTimestamp
}
