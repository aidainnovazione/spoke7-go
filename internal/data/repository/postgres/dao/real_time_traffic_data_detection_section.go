package dao

import (
	"spoke7-go/internal/data/models"
	"time"
)

type DetectionSectionNodeType int

const (
	TYPE_START DetectionSectionNodeType = iota // nodo start
	TYPE_END                                   // nodo end
)

type RealTimeTrafficDataByDetectionSectionDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	DetectionSectionID string    `gorm:"primary_key;index;not null"`
	DetectionTimestamp time.Time `gorm:"primary_key;index;not null"`

	DetectionSectionNodeType DetectionSectionNodeType `gorm:"type:integer;not null;column:detection_section_node_type"`

	DetectionType        string `gorm:"type:varchar(255);not null;column:detection_type"`
	DetectionTechnology  string `gorm:"type:varchar(255);not null;column:detection_technology"` // e.g., floating car data
	AnonymousDetectionID string `gorm:"type:varchar(255);not null;column:anonymous_detection_id"`
}

func (RealTimeTrafficDataByDetectionSectionDao) TableName() string {
	return "real_time_traffic_data_by_detection_section"
}

func (c RealTimeTrafficDataByDetectionSectionDao) FromRealTimeTrafficDataByDetectionSectionDaoToModel() models.RealTimeTrafficDataByDetectionSectionModel {
	return models.RealTimeTrafficDataByDetectionSectionModel{
		CreatedAt:                c.CreatedAt,
		ModifiedAt:               c.ModifiedAt,
		DataSourceName:           c.DataSourceName,
		DetectionSectionID:       c.DetectionSectionID,
		DetectionTimestamp:       c.DetectionTimestamp,
		DetectionSectionNodeType: models.DetectionSectionNodeType(c.DetectionSectionNodeType),
		DetectionType:            c.DetectionType,
		DetectionTechnology:      c.DetectionTechnology,
		AnonymousDetectionID:     c.AnonymousDetectionID,
	}
}

func FromRealTimeTrafficDataByDetectionSectionModelToDao(traffic models.RealTimeTrafficDataByDetectionSectionModel) RealTimeTrafficDataByDetectionSectionDao {
	return RealTimeTrafficDataByDetectionSectionDao{
		CreatedAt:                traffic.CreatedAt,
		ModifiedAt:               traffic.ModifiedAt,
		DataSourceName:           traffic.DataSourceName,
		DetectionSectionID:       traffic.DetectionSectionID,
		DetectionTimestamp:       traffic.DetectionTimestamp,
		DetectionSectionNodeType: DetectionSectionNodeType(traffic.DetectionSectionNodeType),
		DetectionType:            traffic.DetectionType,
		DetectionTechnology:      traffic.DetectionTechnology,
		AnonymousDetectionID:     traffic.AnonymousDetectionID,
	}
}
