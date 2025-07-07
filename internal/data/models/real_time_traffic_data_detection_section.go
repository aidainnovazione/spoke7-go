package models

import "time"

type DetectionSectionNodeType int

const (
	TYPE_START DetectionSectionNodeType = iota // nodo start
	TYPE_END                                   // nodo end
)

type RealTimeTrafficDataByDetectionSectionModel struct {
	CreatedAt          time.Time
	ModifiedAt         time.Time
	DataSourceName     string
	DetectionSectionID string
	DetectionTimestamp time.Time

	DetectionSectionNodeType DetectionSectionNodeType

	DetectionType        string
	DetectionTechnology  string // e.g., floating car data
	AnonymousDetectionID string
}

func (m RealTimeTrafficDataByDetectionSectionModel) GetDetectionTimestamp() time.Time {
	return m.DetectionTimestamp
}
