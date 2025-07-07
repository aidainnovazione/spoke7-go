package dao

import (
	"spoke7-go/internal/data/models"
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

type RealTimeTrafficDataByDetectionPointByLaneDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	LaneID             string    `gorm:"primary_key;type:varchar(255);index;not null"`
	DetectionTimestamp time.Time `gorm:"primary_key;type:timestamp;index;not null"`

	DetectionType        string           `gorm:"type:varchar(255);not null"`
	DetectionTechnology  string           `gorm:"type:varchar(255);not null"`
	VehicleClass         VehicleClassType `gorm:"type:integer;not null"`
	VehicleSpeed         float32          `gorm:"type:float;not null"`
	VehicleLength        float32          `gorm:"type:float;not null"`
	VehicleHeadway       float32          `gorm:"type:float;not null"`
	QueuePresent         bool             `gorm:"type:boolean;not null"`
	CorrectFlowDirection bool             `gorm:"type:boolean;not null"`
}

func (RealTimeTrafficDataByDetectionPointByLaneDao) TableName() string {
	return "real_time_traffic_data_by_detection_point_by_lane"
}

func (c RealTimeTrafficDataByDetectionPointByLaneDao) FromRealTimeTrafficDataByDetectionPointByLaneDaoToModel() models.RealTimeTrafficDataByDetectionPointByLaneModel {
	return models.RealTimeTrafficDataByDetectionPointByLaneModel{
		CreatedAt:            c.CreatedAt,
		ModifiedAt:           c.ModifiedAt,
		DataSourceName:       c.DataSourceName,
		LaneID:               c.LaneID,
		DetectionTimestamp:   c.DetectionTimestamp,
		DetectionType:        c.DetectionType,
		DetectionTechnology:  c.DetectionTechnology,
		VehicleClass:         models.VehicleClassType(c.VehicleClass),
		VehicleSpeed:         c.VehicleSpeed,
		VehicleLength:        c.VehicleLength,
		VehicleHeadway:       c.VehicleHeadway,
		QueuePresent:         c.QueuePresent,
		CorrectFlowDirection: c.CorrectFlowDirection,
	}
}

func FromRealTimeTrafficDataByDetectionPointByLaneModelToDao(traffic models.RealTimeTrafficDataByDetectionPointByLaneModel) RealTimeTrafficDataByDetectionPointByLaneDao {
	return RealTimeTrafficDataByDetectionPointByLaneDao{
		CreatedAt:            traffic.CreatedAt,
		ModifiedAt:           traffic.ModifiedAt,
		DataSourceName:       traffic.DataSourceName,
		LaneID:               traffic.LaneID,
		DetectionTimestamp:   traffic.DetectionTimestamp,
		DetectionType:        traffic.DetectionType,
		DetectionTechnology:  traffic.DetectionTechnology,
		VehicleClass:         VehicleClassType(traffic.VehicleClass),
		VehicleSpeed:         traffic.VehicleSpeed,
		VehicleLength:        traffic.VehicleLength,
		VehicleHeadway:       traffic.VehicleHeadway,
		QueuePresent:         traffic.QueuePresent,
		CorrectFlowDirection: traffic.CorrectFlowDirection,
	}
}
