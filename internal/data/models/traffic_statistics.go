package models

import "time"

type TrafficStatisticsDetectionSection struct {
	DetectionSectionID           string
	RecordsCount                 uint32
	FirstRecordTimestamp         time.Time
	LastRecordTimestamp          time.Time
	LongestDataGap               float32
	LongestDataGapStartTimestamp time.Time
	LongestDataGapEndTimestamp   time.Time
}

type TrafficStatisticsLane struct {
	LaneID                       string
	RecordsCount                 uint32
	FirstRecordTimestamp         time.Time
	LastRecordTimestamp          time.Time
	LongestDataGap               float32
	LongestDataGapStartTimestamp time.Time
	LongestDataGapEndTimestamp   time.Time
}

type TrafficStatisticsDetectionPoint struct {
	DetectionPointID             string
	RecordsCount                 uint32
	FirstRecordTimestamp         time.Time
	LastRecordTimestamp          time.Time
	LongestDataGap               float32
	LongestDataGapStartTimestamp time.Time
	LongestDataGapEndTimestamp   time.Time
}

type TimeStampedStatistic interface {
	GetFirstRecordTimestamp() time.Time
	GetLastRecordTimestamp() time.Time
	GetRecordsCount() uint32
}

// real time - section
func (s TrafficStatisticsDetectionSection) GetFirstRecordTimestamp() time.Time {
	return s.FirstRecordTimestamp
}

func (s TrafficStatisticsDetectionSection) GetLastRecordTimestamp() time.Time {
	return s.LastRecordTimestamp
}

func (s TrafficStatisticsDetectionSection) GetRecordsCount() uint32 {
	return s.RecordsCount
}

// real time- lane
func (s TrafficStatisticsLane) GetFirstRecordTimestamp() time.Time {
	return s.FirstRecordTimestamp
}

func (s TrafficStatisticsLane) GetLastRecordTimestamp() time.Time {
	return s.LastRecordTimestamp
}

func (s TrafficStatisticsLane) GetRecordsCount() uint32 {
	return s.RecordsCount
}

func (s TrafficStatisticsDetectionPoint) GetFirstRecordTimestamp() time.Time {
	return s.FirstRecordTimestamp
}

func (s TrafficStatisticsDetectionPoint) GetLastRecordTimestamp() time.Time {
	return s.LastRecordTimestamp
}

func (s TrafficStatisticsDetectionPoint) GetRecordsCount() uint32 {
	return s.RecordsCount
}

// current - detection section
func (s CurrentTrafficDataByDetectionSectionStatistics) GetFirstRecordTimestamp() time.Time {
	return s.FirstRecordTimestamp
}

func (s CurrentTrafficDataByDetectionSectionStatistics) GetLastRecordTimestamp() time.Time {
	return s.LastRecordTimestamp
}

func (s CurrentTrafficDataByDetectionSectionStatistics) GetRecordsCount() uint32 {
	return s.RecordsCount
}

// current- dtection point
func (s CurrentTrafficDataByDetectionPointStatistics) GetFirstRecordTimestamp() time.Time {
	return s.FirstRecordTimestamp
}

func (s CurrentTrafficDataByDetectionPointStatistics) GetLastRecordTimestamp() time.Time {
	return s.LastRecordTimestamp
}

func (s CurrentTrafficDataByDetectionPointStatistics) GetRecordsCount() uint32 {
	return s.RecordsCount
}

// current- lane
func (s CurrentTrafficDataByDetectionPointByLaneStatistics) GetFirstRecordTimestamp() time.Time {
	return s.FirstRecordTimestamp
}

func (s CurrentTrafficDataByDetectionPointByLaneStatistics) GetLastRecordTimestamp() time.Time {
	return s.LastRecordTimestamp
}

func (s CurrentTrafficDataByDetectionPointByLaneStatistics) GetRecordsCount() uint32 {
	return s.RecordsCount
}
