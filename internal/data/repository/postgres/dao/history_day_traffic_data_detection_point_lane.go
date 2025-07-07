package dao

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"spoke7-go/internal/data/models"
	"time"
)

type HistoryDayTrafficDataByDetectionPointByLaneDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	LaneID             string    `gorm:"primary_key;type:varchar(255);index;not null"`
	DetectionTimestamp time.Time `gorm:"primary_key;type:timestamp;index;not null"`
	DetectionInterval  uint32    `gorm:"type:integer;not null"`

	TrafficFlowVehicleClass1          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass2          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass3          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass4          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass5          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass6          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass7          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClass8          float32 `gorm:"type:float;not null"`
	TrafficFlowVehicleClassEquivalent float32 `gorm:"type:float;not null"`

	AverageSpeedVehicleClass1   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass2   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass3   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass4   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass5   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass6   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass7   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClass8   float32 `gorm:"type:float;not null"`
	AverageSpeedVehicleClassAll float32 `gorm:"type:float;not null"`

	AverageVehicleLength   float32 `gorm:"type:float;not null"`
	AverageHeadway         float32 `gorm:"type:float;not null"`
	AverageTimeToCollision float32 `gorm:"type:float;not null"`

	TrafficFlowParametersVehicleClass1 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass2 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass3 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass4 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass5 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass6 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass7 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`
	TrafficFlowParametersVehicleClass8 TrafficFlowParametersByDayDao `gorm:"type:jsonb"`

	InstantaneousSpeedVelClass1MaxFlow float32 `gorm:"type:float;not null"`
	InstantaneousSpeedVelClass2MaxFlow float32 `gorm:"type:float;not null"`
	InstantaneousSpeedVelClass3MaxFlow float32 `gorm:"type:float;not null"`

	InstantaneousSpeedVelClass1MinFlow float32 `gorm:"type:float;not null"`
	InstantaneousSpeedVelClass2MinFlow float32 `gorm:"type:float;not null"`
	InstantaneousSpeedVelClass3MinFlow float32 `gorm:"type:float;not null"`

	HeadwayVelClass1MaxFlow float32 `gorm:"type:float;not null"`
	HeadwayVelClass2MaxFlow float32 `gorm:"type:float;not null"`
	HeadwayVelClass3MaxFlow float32 `gorm:"type:float;not null"`

	HeadwayVelClass1MinFlow float32 `gorm:"type:float;not null"`
	HeadwayVelClass2MinFlow float32 `gorm:"type:float;not null"`
	HeadwayVelClass3MinFlow float32 `gorm:"type:float;not null"`
}

func (HistoryDayTrafficDataByDetectionPointByLaneDao) TableName() string {
	return "history_day_traffic_data_by_detection_point_by_lane"
}

func (c HistoryDayTrafficDataByDetectionPointByLaneDao) FromHistoryDayTrafficDataByDetectionPointByLaneDaoToModel() models.HistoryDayTrafficDataByDetectionPointByLaneModel {
	trafficFlowParametersVehicleClass1 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass1.Min,
		Max:     c.TrafficFlowParametersVehicleClass1.Max,
		Average: c.TrafficFlowParametersVehicleClass1.Average,
		Std:     c.TrafficFlowParametersVehicleClass1.Std,
	}
	trafficFlowParametersVehicleClass2 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass2.Min,
		Max:     c.TrafficFlowParametersVehicleClass2.Max,
		Average: c.TrafficFlowParametersVehicleClass2.Average,
		Std:     c.TrafficFlowParametersVehicleClass2.Std,
	}
	trafficFlowParametersVehicleClass3 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass3.Min,
		Max:     c.TrafficFlowParametersVehicleClass3.Max,
		Average: c.TrafficFlowParametersVehicleClass3.Average,
		Std:     c.TrafficFlowParametersVehicleClass3.Std,
	}
	trafficFlowParametersVehicleClass4 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass4.Min,
		Max:     c.TrafficFlowParametersVehicleClass4.Max,
		Average: c.TrafficFlowParametersVehicleClass4.Average,
		Std:     c.TrafficFlowParametersVehicleClass4.Std,
	}
	trafficFlowParametersVehicleClass5 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass5.Min,
		Max:     c.TrafficFlowParametersVehicleClass5.Max,
		Average: c.TrafficFlowParametersVehicleClass5.Average,
		Std:     c.TrafficFlowParametersVehicleClass5.Std,
	}
	trafficFlowParametersVehicleClass6 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass6.Min,
		Max:     c.TrafficFlowParametersVehicleClass6.Max,
		Average: c.TrafficFlowParametersVehicleClass6.Average,
		Std:     c.TrafficFlowParametersVehicleClass6.Std,
	}
	trafficFlowParametersVehicleClass7 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass7.Min,
		Max:     c.TrafficFlowParametersVehicleClass7.Max,
		Average: c.TrafficFlowParametersVehicleClass7.Average,
		Std:     c.TrafficFlowParametersVehicleClass7.Std,
	}
	trafficFlowParametersVehicleClass8 := models.TrafficFlowParametersByDay{
		Min:     c.TrafficFlowParametersVehicleClass8.Min,
		Max:     c.TrafficFlowParametersVehicleClass8.Max,
		Average: c.TrafficFlowParametersVehicleClass8.Average,
		Std:     c.TrafficFlowParametersVehicleClass8.Std,
	}
	return models.HistoryDayTrafficDataByDetectionPointByLaneModel{
		CreatedAt:          c.CreatedAt,
		ModifiedAt:         c.ModifiedAt,
		DataSourceName:     c.DataSourceName,
		LaneID:             c.LaneID,
		DetectionTimestamp: c.DetectionTimestamp,
		DetectionInterval:  c.DetectionInterval,

		TrafficFlowVehicleClass1:          c.TrafficFlowVehicleClass1,
		TrafficFlowVehicleClass2:          c.TrafficFlowVehicleClass2,
		TrafficFlowVehicleClass3:          c.TrafficFlowVehicleClass3,
		TrafficFlowVehicleClass4:          c.TrafficFlowVehicleClass4,
		TrafficFlowVehicleClass5:          c.TrafficFlowVehicleClass5,
		TrafficFlowVehicleClass6:          c.TrafficFlowVehicleClass6,
		TrafficFlowVehicleClass7:          c.TrafficFlowVehicleClass7,
		TrafficFlowVehicleClass8:          c.TrafficFlowVehicleClass8,
		TrafficFlowVehicleClassEquivalent: c.TrafficFlowVehicleClassEquivalent,

		AverageSpeedVehicleClass1:   c.AverageSpeedVehicleClass1,
		AverageSpeedVehicleClass2:   c.AverageSpeedVehicleClass2,
		AverageSpeedVehicleClass3:   c.AverageSpeedVehicleClass3,
		AverageSpeedVehicleClass4:   c.AverageSpeedVehicleClass4,
		AverageSpeedVehicleClass5:   c.AverageSpeedVehicleClass5,
		AverageSpeedVehicleClass6:   c.AverageSpeedVehicleClass6,
		AverageSpeedVehicleClass7:   c.AverageSpeedVehicleClass7,
		AverageSpeedVehicleClass8:   c.AverageSpeedVehicleClass8,
		AverageSpeedVehicleClassAll: c.AverageSpeedVehicleClassAll,

		AverageVehicleLength:   c.AverageVehicleLength,
		AverageHeadway:         c.AverageHeadway,
		AverageTimeToCollision: c.AverageTimeToCollision,

		TrafficFlowParametersVehicleClass1: trafficFlowParametersVehicleClass1,
		TrafficFlowParametersVehicleClass2: trafficFlowParametersVehicleClass2,
		TrafficFlowParametersVehicleClass3: trafficFlowParametersVehicleClass3,
		TrafficFlowParametersVehicleClass4: trafficFlowParametersVehicleClass4,
		TrafficFlowParametersVehicleClass5: trafficFlowParametersVehicleClass5,
		TrafficFlowParametersVehicleClass6: trafficFlowParametersVehicleClass6,
		TrafficFlowParametersVehicleClass7: trafficFlowParametersVehicleClass7,
		TrafficFlowParametersVehicleClass8: trafficFlowParametersVehicleClass8,

		InstantaneousSpeedVelClass1MaxFlow: c.InstantaneousSpeedVelClass1MaxFlow,
		InstantaneousSpeedVelClass2MaxFlow: c.InstantaneousSpeedVelClass2MaxFlow,
		InstantaneousSpeedVelClass3MaxFlow: c.InstantaneousSpeedVelClass3MaxFlow,

		InstantaneousSpeedVelClass1MinFlow: c.InstantaneousSpeedVelClass1MinFlow,
		InstantaneousSpeedVelClass2MinFlow: c.InstantaneousSpeedVelClass2MinFlow,
		InstantaneousSpeedVelClass3MinFlow: c.InstantaneousSpeedVelClass3MinFlow,

		HeadwayVelClass1MaxFlow: c.HeadwayVelClass1MaxFlow,
		HeadwayVelClass2MaxFlow: c.HeadwayVelClass2MaxFlow,
		HeadwayVelClass3MaxFlow: c.HeadwayVelClass3MaxFlow,

		HeadwayVelClass1MinFlow: c.HeadwayVelClass1MinFlow,
		HeadwayVelClass2MinFlow: c.HeadwayVelClass2MinFlow,
		HeadwayVelClass3MinFlow: c.HeadwayVelClass3MinFlow,
	}
}

func FromHistoryDayTrafficDataByDetectionPointByLaneModelToDao(traffic models.HistoryDayTrafficDataByDetectionPointByLaneModel) HistoryDayTrafficDataByDetectionPointByLaneDao {
	trafficFlowParametersVehicleClass1 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass1.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass1.Max,
		Average: traffic.TrafficFlowParametersVehicleClass1.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass1.Std,
	}
	trafficFlowParametersVehicleClass2 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass2.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass2.Max,
		Average: traffic.TrafficFlowParametersVehicleClass2.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass2.Std,
	}
	trafficFlowParametersVehicleClass3 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass3.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass3.Max,
		Average: traffic.TrafficFlowParametersVehicleClass3.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass3.Std,
	}
	trafficFlowParametersVehicleClass4 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass4.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass4.Max,
		Average: traffic.TrafficFlowParametersVehicleClass4.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass4.Std,
	}
	trafficFlowParametersVehicleClass5 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass5.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass5.Max,
		Average: traffic.TrafficFlowParametersVehicleClass5.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass5.Std,
	}
	trafficFlowParametersVehicleClass6 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass6.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass6.Max,
		Average: traffic.TrafficFlowParametersVehicleClass6.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass6.Std,
	}
	trafficFlowParametersVehicleClass7 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass7.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass7.Max,
		Average: traffic.TrafficFlowParametersVehicleClass7.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass7.Std,
	}
	trafficFlowParametersVehicleClass8 := TrafficFlowParametersByDayDao{
		Min:     traffic.TrafficFlowParametersVehicleClass8.Min,
		Max:     traffic.TrafficFlowParametersVehicleClass8.Max,
		Average: traffic.TrafficFlowParametersVehicleClass8.Average,
		Std:     traffic.TrafficFlowParametersVehicleClass8.Std,
	}
	return HistoryDayTrafficDataByDetectionPointByLaneDao{
		CreatedAt:                         traffic.CreatedAt,
		ModifiedAt:                        traffic.ModifiedAt,
		DataSourceName:                    traffic.DataSourceName,
		LaneID:                            traffic.LaneID,
		DetectionTimestamp:                traffic.DetectionTimestamp,
		DetectionInterval:                 traffic.DetectionInterval,
		TrafficFlowVehicleClass1:          traffic.TrafficFlowVehicleClass1,
		TrafficFlowVehicleClass2:          traffic.TrafficFlowVehicleClass2,
		TrafficFlowVehicleClass3:          traffic.TrafficFlowVehicleClass3,
		TrafficFlowVehicleClass4:          traffic.TrafficFlowVehicleClass4,
		TrafficFlowVehicleClass5:          traffic.TrafficFlowVehicleClass5,
		TrafficFlowVehicleClass6:          traffic.TrafficFlowVehicleClass6,
		TrafficFlowVehicleClass7:          traffic.TrafficFlowVehicleClass7,
		TrafficFlowVehicleClass8:          traffic.TrafficFlowVehicleClass8,
		TrafficFlowVehicleClassEquivalent: traffic.TrafficFlowVehicleClassEquivalent,

		AverageSpeedVehicleClass1:   traffic.AverageSpeedVehicleClass1,
		AverageSpeedVehicleClass2:   traffic.AverageSpeedVehicleClass2,
		AverageSpeedVehicleClass3:   traffic.AverageSpeedVehicleClass3,
		AverageSpeedVehicleClass4:   traffic.AverageSpeedVehicleClass4,
		AverageSpeedVehicleClass5:   traffic.AverageSpeedVehicleClass5,
		AverageSpeedVehicleClass6:   traffic.AverageSpeedVehicleClass6,
		AverageSpeedVehicleClass7:   traffic.AverageSpeedVehicleClass7,
		AverageSpeedVehicleClass8:   traffic.AverageSpeedVehicleClass8,
		AverageSpeedVehicleClassAll: traffic.AverageSpeedVehicleClassAll,

		AverageVehicleLength:   traffic.AverageVehicleLength,
		AverageHeadway:         traffic.AverageHeadway,
		AverageTimeToCollision: traffic.AverageTimeToCollision,

		TrafficFlowParametersVehicleClass1: trafficFlowParametersVehicleClass1,
		TrafficFlowParametersVehicleClass2: trafficFlowParametersVehicleClass2,
		TrafficFlowParametersVehicleClass3: trafficFlowParametersVehicleClass3,
		TrafficFlowParametersVehicleClass4: trafficFlowParametersVehicleClass4,
		TrafficFlowParametersVehicleClass5: trafficFlowParametersVehicleClass5,
		TrafficFlowParametersVehicleClass6: trafficFlowParametersVehicleClass6,
		TrafficFlowParametersVehicleClass7: trafficFlowParametersVehicleClass7,
		TrafficFlowParametersVehicleClass8: trafficFlowParametersVehicleClass8,

		InstantaneousSpeedVelClass1MaxFlow: traffic.InstantaneousSpeedVelClass1MaxFlow,
		InstantaneousSpeedVelClass2MaxFlow: traffic.InstantaneousSpeedVelClass2MaxFlow,
		InstantaneousSpeedVelClass3MaxFlow: traffic.InstantaneousSpeedVelClass3MaxFlow,

		InstantaneousSpeedVelClass1MinFlow: traffic.InstantaneousSpeedVelClass1MinFlow,
		InstantaneousSpeedVelClass2MinFlow: traffic.InstantaneousSpeedVelClass2MinFlow,
		InstantaneousSpeedVelClass3MinFlow: traffic.InstantaneousSpeedVelClass3MinFlow,

		HeadwayVelClass1MaxFlow: traffic.HeadwayVelClass1MaxFlow,
		HeadwayVelClass2MaxFlow: traffic.HeadwayVelClass2MaxFlow,
		HeadwayVelClass3MaxFlow: traffic.HeadwayVelClass3MaxFlow,

		HeadwayVelClass1MinFlow: traffic.HeadwayVelClass1MinFlow,
		HeadwayVelClass2MinFlow: traffic.HeadwayVelClass2MinFlow,
		HeadwayVelClass3MinFlow: traffic.HeadwayVelClass3MinFlow,
	}
}

// flow parameters

type TrafficFlowParametersByDayDao struct {
	Min     uint32
	Max     uint32
	Average uint32
	Std     uint32
}

func (p TrafficFlowParametersByDayDao) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TrafficFlowParametersByDayDao: %w", err)
	}
	return data, nil
}

func (p *TrafficFlowParametersByDayDao) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan TrafficFlowParametersByDayDao: expected []byte but got %T", value)
	}
	return json.Unmarshal(bytes, p)
}
