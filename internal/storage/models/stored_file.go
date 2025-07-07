package models

import "time"

type FileType int

const (
	FileTypeUnknown FileType = iota
	FileTypeRealTimeTrafficByLane
	FileTypeRealTimeTrafficByDetectionSection
	FileTypeAggregatedTraffic5MinByLane
	FileTypeAggregatedTraffic5MinByDetectionPoint
	FileTypeAggregatedTraffic5MinByDetectionSection
	FileTypeAggregatedTraffic1HourByLane
	FileTypeAggregatedTraffic1HourByDetectionPoint
	FileTypeAggregatedTraffic1HourByDetectionSection
	FileTypeAggregatedTrafficDayByLane
	FileTypeAggregatedTrafficDayByDetectionPoint
	FileTypeAggregatedTrafficDayByDetectionSection
	FileTypeSumoNetwork
	FileTypeSumoAdditional
	FileTypeSumoRoutes
	FileTypeDetectionSections
	FileTypeDetectionPoints
	FileTypeNetwork
)

type StoredFile struct {
	ID             string
	DataSourceName string
	Description    string
	Tag            string
	Owner          string
	Groups         []string
	CreatedAt      time.Time
	ModifiedAt     time.Time
	FileName       string
	FileSize       uint32
	FileType       FileType
	FileFormat     string
	FileContent    []byte
}

type StoredFileUpload struct {
	DataSourceName string
	Description    string
	Tag            string
	Owner          string
	Groups         []string
	FileType       FileType
	FileName       string
	FileSize       uint32
	FileFormat     string
	FileContent    []byte
}

type StoredFileUpdate struct {
	ID             string
	DataSourceName string
	Description    string
	Tag            string
	Owner          string
	Groups         []string
	FileType       FileType
	FileName       string
	FileSize       uint32
	FileFormat     string
	FileContent    []byte
}
