package models

import (
	"time"
)

type HistoryTrafficDataByDetectionSectionModel struct {
	CreatedAt                 time.Time
	ModifiedAt                time.Time
	DataSourceName            string
	DetectionSectionID        string
	DetectionTimestamp        time.Time
	DetectionInterval         uint32
	ForwardSpeedCount         float32
	ForwardSpeedCountMaxFlow  float32
	ForwardSpeedCountMinFlow  float32
	BackwardSpeedCount        float32
	BackwardSpeedCountMaxFlow float32
	BackwardSpeedCountMinFlow float32
}

func (m HistoryTrafficDataByDetectionSectionModel) GetDetectionTimestamp() time.Time {
	return m.DetectionTimestamp
}
