package dao

import (
	"spoke7-go/internal/data/models"
	"time"
)

type CurrentTrafficDataByDetectionSectionDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	DetectionSectionID string    `gorm:"primary_key;index;not null"`
	DetectionTimestamp time.Time `gorm:"primary_key;index;not null"`
	DetectionInterval  uint32    `gorm:"type_integer;not null"`

	ForwardSpeed  float32 `gorm:"type:float8"`
	BackwardSpeed float32 `gorm:"type:float8"`
}

func (CurrentTrafficDataByDetectionSectionDao) TableName() string {
	return "current_traffic_data_by_detection_section"
}

func (c CurrentTrafficDataByDetectionSectionDao) FromCurrentTrafficDataByDetectionSectionDaoToModel() models.CurrentTrafficDataByDetectionSectionModel {
	return models.CurrentTrafficDataByDetectionSectionModel{
		CreatedAt:          c.CreatedAt,
		ModifiedAt:         c.ModifiedAt,
		DataSourceName:     c.DataSourceName,
		DetectionSectionID: c.DetectionSectionID,
		DetectionTimestamp: c.DetectionTimestamp,
		DetectionInterval:  c.DetectionInterval,
		ForwardSpeed:       c.ForwardSpeed,
		BackwardSpeed:      c.BackwardSpeed,
	}
}

func FromCurrentTrafficDataByDetectionSectionModelToDao(traffic models.CurrentTrafficDataByDetectionSectionModel) CurrentTrafficDataByDetectionSectionDao {
	return CurrentTrafficDataByDetectionSectionDao{
		CreatedAt:          traffic.CreatedAt,
		ModifiedAt:         traffic.ModifiedAt,
		DataSourceName:     traffic.DataSourceName,
		DetectionSectionID: traffic.DetectionSectionID,
		DetectionTimestamp: traffic.DetectionTimestamp,
		DetectionInterval:  traffic.DetectionInterval,
		ForwardSpeed:       traffic.ForwardSpeed,
		BackwardSpeed:      traffic.BackwardSpeed,
	}
}
