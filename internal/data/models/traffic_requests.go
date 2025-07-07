package models

import "time"

// DETECTION POINT BY LANE
type GetTrafficDataByDetectionPointByLaneParams struct {
	DataSourceName     string
	DetectionTimestamp time.Time
	LaneIDs            []string
}

type ListTrafficDataByDetectionPointByLaneParams struct {
	DataSourceName string
	LaneIDs        []string
	StartTime      time.Time
	EndTime        time.Time
}

type DeleteTrafficDataByDetectionPointByLaneParams struct {
	DataSourceName string
	LaneIDs        []string
	StartTimestamp time.Time
	EndTimestamp   time.Time
}

// DETECTION POINT

type GetTrafficDataByDetectionPointParams struct {
	DataSourceName     string
	DetectionPointIDs  []string
	DetectionTimestamp time.Time
}

type ListTrafficDataByDetectionPointParams struct {
	DataSourceName    string
	DetectionPointIDs []string
	StartTime         time.Time
	EndTime           time.Time
}

type DeleteTrafficDataByDetectionPointParams struct {
	DataSourceName    string
	DetectionPointIDs []string
	StartTimestamp    time.Time
	EndTimestamp      time.Time
}

// DETECTION SECTION
type GetTrafficDataByDetectionSectionParams struct {
	DataSourceName      string
	DetectionSectionIDs []string
	DetectionTimestamp  time.Time
}

type ListTrafficDataByDetectionSectionParams struct {
	DataSourceName      string
	DetectionSectionIDs []string
	StartTime           time.Time
	EndTime             time.Time
}

type DeleteTrafficDataByDetectionSectionParams struct {
	DataSourceName      string
	DetectionSectionIDs []string
	StartTimestamp      time.Time
	EndTimestamp        time.Time
}
