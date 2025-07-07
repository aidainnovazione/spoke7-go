package models

import "time"

type CurrentTrafficDataByDetectionSectionModel struct {
	CreatedAt          time.Time
	ModifiedAt         time.Time
	DataSourceName     string
	DetectionSectionID string
	DetectionTimestamp time.Time
	DetectionInterval  uint32
	ForwardSpeed       float32
	BackwardSpeed      float32
}

type CurrentTrafficDataByDetectionSectionStatistics struct {
	DetectionSectionID           string
	RecordsCount                 uint32
	FirstRecordTimestamp         time.Time
	LastRecordTimestamp          time.Time
	LongestDataGap               float32
	LongestDataGapStartTimestamp time.Time
	LongestDataGapEndTimestamp   time.Time
	MissingIntervalDetectionRate float32 // Percentage of expected intervals that are missing data.

	TotalAverageForwardSpeed           float32
	TotalAverageBackwardSpeed          float32
	PercentageRecordsWithForwardSpeed  float32 // percentage of records where fw speed is present
	PercentageRecordsWithBackwardSpeed float32 // percentage of records where bw speed is present
}

type CurrentTrafficDataByDetectionSectionStatisticsByDatasource struct {
	DataSourceName       string
	RecordsCount         uint32
	FirstRecordTimestamp time.Time
	LastRecordTimestamp  time.Time
	StatisticsBySection  []CurrentTrafficDataByDetectionSectionStatistics
}
