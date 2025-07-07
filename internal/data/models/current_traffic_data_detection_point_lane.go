package models

import (
	"time"
)

type CurrentTrafficDataByDetectionPointByLaneModel struct {
	CreatedAt                              time.Time
	ModifiedAt                             time.Time
	DataSourceName                         string
	LaneID                                 string
	DetectionTimestamp                     time.Time
	DetectionInterval                      uint32
	CountVehicleClass1                     uint32
	CountVehicleClass2                     uint32
	CountVehicleClass3                     uint32
	CountVehicleClass4                     uint32
	CountVehicleClass5                     uint32
	CountVehicleClass6                     uint32
	CountVehicleClass7                     uint32
	CountVehicleClass8                     uint32
	CountVehicleClassAll                   uint32
	HarmonicMeanSpeedVehicleClass1         float32
	HarmonicMeanSpeedVehicleClass2         float32
	HarmonicMeanSpeedVehicleClass3         float32
	HarmonicMeanSpeedVehicleClass4         float32
	HarmonicMeanSpeedVehicleClass5         float32
	HarmonicMeanSpeedVehicleClass6         float32
	HarmonicMeanSpeedVehicleClass7         float32
	HarmonicMeanSpeedVehicleClass8         float32
	HarmonicMeanSpeedVehicleClassAll       float32
	CountDetectedSpeedVehicleUnder50       uint32
	CountDetectedSpeedVehicleBetween50_100 uint32
	CountDetectedSpeedVehicleOver100       uint32
	AverageVehicleLength                   float32
	AverageHeadway                         float32
	StdHeadway                             float32
	AverageTimeToCollision                 float32
	StdTimeToCollision                     float32
}

type CurrentTrafficDataByDetectionPointByLaneStatisticsByDatasource struct {
	DataSourceName       string
	RecordsCount         uint32
	FirstRecordTimestamp time.Time
	LastRecordTimestamp  time.Time
	StatisticsByLane     []CurrentTrafficDataByDetectionPointByLaneStatistics
}

type CurrentTrafficDataByDetectionPointByLaneStatistics struct {
	LaneID                       string
	RecordsCount                 uint32
	FirstRecordTimestamp         time.Time
	LastRecordTimestamp          time.Time
	LongestDataGap               float32
	LongestDataGapStartTimestamp time.Time
	LongestDataGapEndTimestamp   time.Time
	MissingIntervalDetectionRate float32 // Percentage of expected intervals that are missing data.
	//total
	TotalCountAllVehicles            uint32  // all counted vehicles
	TotalHarmonicMeanSpeedAllRecords float32 // mean speed of all means
	// rate
	PercentageRecordsWithCounts               float32 // how many records have at least one count
	PercentageRecordsWithPositiveAverageSpeed float32 // how many records have an average speed more than 0
}

type CurrentTrafficDataByDetectionPointByLaneAggregate struct {
	DataSourceName                               string
	DetectionTimestamp                           time.Time
	RecordsCount                                 uint32
	AggPercentageRecords                         float32
	AggPercentageRecordsWithCounts               float32
	AggPercentageRecordsWithPositiveAverageSpeed float32
	AggCountVehicleClassAll                      float32
	AggHarmonicMeanSpeedVehicleClassAll          float32
	AggAverageVehicleLength                      float32
	AggAverageHeadway                            float32
	AggStdHeadway                                float32
	AggAverageTimeToCollision                    float32
	AggStdTimeToCollision                        float32
}
