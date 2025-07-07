package dao

import (
	"spoke7-go/internal/data/models"
	"time"
)

type HistoryHourTrafficDataByDetectionPointDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	DetectionPointID   string    `gorm:"primary_key;type:varchar(255);index;not null"`
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
}

func (HistoryHourTrafficDataByDetectionPointDao) TableName() string {
	return "history_hour_traffic_data_by_detection_point"
}

func (c HistoryHourTrafficDataByDetectionPointDao) FromHistoryHourTrafficDataByDetectionPointDaoToModel() models.HistoryHourTrafficDataByDetectionPointModel {
	return models.HistoryHourTrafficDataByDetectionPointModel{
		CreatedAt:          c.CreatedAt,
		ModifiedAt:         c.ModifiedAt,
		DataSourceName:     c.DataSourceName,
		DetectionPointID:   c.DetectionPointID,
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
		AverageSpeedVehicleClass1:         c.AverageSpeedVehicleClass1,
		AverageSpeedVehicleClass2:         c.AverageSpeedVehicleClass2,
		AverageSpeedVehicleClass3:         c.AverageSpeedVehicleClass3,
		AverageSpeedVehicleClass4:         c.AverageSpeedVehicleClass4,
		AverageSpeedVehicleClass5:         c.AverageSpeedVehicleClass5,
		AverageSpeedVehicleClass6:         c.AverageSpeedVehicleClass6,
		AverageSpeedVehicleClass7:         c.AverageSpeedVehicleClass7,
		AverageSpeedVehicleClass8:         c.AverageSpeedVehicleClass8,
		AverageSpeedVehicleClassAll:       c.AverageSpeedVehicleClassAll,
		AverageVehicleLength:              c.AverageVehicleLength,
		AverageHeadway:                    c.AverageHeadway,
		AverageTimeToCollision:            c.AverageTimeToCollision,
	}
}

func FromHistoryHourTrafficDataByDetectionPointModelToDao(traffic models.HistoryHourTrafficDataByDetectionPointModel) HistoryHourTrafficDataByDetectionPointDao {
	return HistoryHourTrafficDataByDetectionPointDao{
		CreatedAt:                         traffic.CreatedAt,
		ModifiedAt:                        traffic.ModifiedAt,
		DataSourceName:                    traffic.DataSourceName,
		DetectionPointID:                  traffic.DetectionPointID,
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
		AverageSpeedVehicleClass1:         traffic.AverageSpeedVehicleClass1,
		AverageSpeedVehicleClass2:         traffic.AverageSpeedVehicleClass2,
		AverageSpeedVehicleClass3:         traffic.AverageSpeedVehicleClass3,
		AverageSpeedVehicleClass4:         traffic.AverageSpeedVehicleClass4,
		AverageSpeedVehicleClass5:         traffic.AverageSpeedVehicleClass5,
		AverageSpeedVehicleClass6:         traffic.AverageSpeedVehicleClass6,
		AverageSpeedVehicleClass7:         traffic.AverageSpeedVehicleClass7,
		AverageSpeedVehicleClass8:         traffic.AverageSpeedVehicleClass8,
		AverageSpeedVehicleClassAll:       traffic.AverageSpeedVehicleClassAll,
		AverageVehicleLength:              traffic.AverageVehicleLength,
		AverageHeadway:                    traffic.AverageHeadway,
		AverageTimeToCollision:            traffic.AverageTimeToCollision,
	}
}
