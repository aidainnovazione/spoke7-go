package models

import "time"

type TrafficDataByDetectionSectionStatisticsByDatasource struct {
	DataSourceName       string
	RecordsCount         uint32
	FirstRecordTimestamp time.Time
	LastRecordTimestamp  time.Time
	StatisticsBySection  []TrafficStatisticsDetectionSection
}

type TrafficDataByDetectionPointStatisticsByDatasource struct {
	DataSourceName       string
	RecordsCount         uint32
	FirstRecordTimestamp time.Time
	LastRecordTimestamp  time.Time
	StatisticsByPoint    []TrafficStatisticsDetectionPoint
}

type TrafficDataByDetectionPointByLaneStatisticsByDatasource struct {
	DataSourceName       string
	RecordsCount         uint32
	FirstRecordTimestamp time.Time
	LastRecordTimestamp  time.Time
	StatisticsByLane     []TrafficStatisticsLane
}
