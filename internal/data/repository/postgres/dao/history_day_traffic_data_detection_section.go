package dao

import (
	"spoke7-go/internal/data/models"
	"time"
)

type HistoryDayTrafficDataByDetectionSectionDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	DetectionSectionID string    `gorm:"primary_key;index;not null"`
	DetectionTimestamp time.Time `gorm:"primary_key;index;not null"`
	DetectionInterval  uint32    `gorm:"type_integer;not null"`

	ForwardSpeedCount         float32 `gorm:"type:float8"`
	ForwardSpeedCountMaxFlow  float32 `gorm:"type:float8"`
	ForwardSpeedCountMinFlow  float32 `gorm:"type:float8"`
	BackwardSpeedCount        float32 `gorm:"type:float8"`
	BackwardSpeedCountMaxFlow float32 `gorm:"type:float8"`
	BackwardSpeedCountMinFlow float32 `gorm:"type:float8"`
}

func (HistoryDayTrafficDataByDetectionSectionDao) TableName() string {
	return "history_day_traffic_data_by_detection_section"
}

func (c HistoryDayTrafficDataByDetectionSectionDao) FromHistoryDayTrafficDataByDetectionSectionDaoToModel() models.HistoryTrafficDataByDetectionSectionModel {
	return models.HistoryTrafficDataByDetectionSectionModel{
		CreatedAt:                 c.CreatedAt,
		ModifiedAt:                c.ModifiedAt,
		DataSourceName:            c.DataSourceName,
		DetectionSectionID:        c.DetectionSectionID,
		DetectionTimestamp:        c.DetectionTimestamp,
		DetectionInterval:         c.DetectionInterval,
		ForwardSpeedCount:         c.ForwardSpeedCount,
		ForwardSpeedCountMaxFlow:  c.ForwardSpeedCountMaxFlow,
		ForwardSpeedCountMinFlow:  c.ForwardSpeedCountMinFlow,
		BackwardSpeedCount:        c.BackwardSpeedCount,
		BackwardSpeedCountMaxFlow: c.BackwardSpeedCountMaxFlow,
		BackwardSpeedCountMinFlow: c.BackwardSpeedCountMinFlow,
	}
}

func FromHistoryDayTrafficDataByDetectionSectionModelToDao(traffic models.HistoryTrafficDataByDetectionSectionModel) HistoryDayTrafficDataByDetectionSectionDao {
	return HistoryDayTrafficDataByDetectionSectionDao{
		CreatedAt:                 traffic.CreatedAt,
		ModifiedAt:                traffic.ModifiedAt,
		DataSourceName:            traffic.DataSourceName,
		DetectionSectionID:        traffic.DetectionSectionID,
		DetectionTimestamp:        traffic.DetectionTimestamp,
		DetectionInterval:         traffic.DetectionInterval,
		ForwardSpeedCount:         traffic.ForwardSpeedCount,
		ForwardSpeedCountMaxFlow:  traffic.ForwardSpeedCountMaxFlow,
		ForwardSpeedCountMinFlow:  traffic.ForwardSpeedCountMinFlow,
		BackwardSpeedCount:        traffic.BackwardSpeedCount,
		BackwardSpeedCountMaxFlow: traffic.BackwardSpeedCountMaxFlow,
		BackwardSpeedCountMinFlow: traffic.BackwardSpeedCountMinFlow,
	}
}
