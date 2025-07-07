package models

import (
	"time"
)

type CurrentTrafficDataByDetectionPointModel struct {
	CreatedAt  time.Time
	ModifiedAt time.Time

	DataSourceName     string
	DetectionPointID   string
	DetectionTimestamp time.Time
	DetectionInterval  uint32

	CountVehicleClass1 uint32
	CountVehicleClass2 uint32
	CountVehicleClass3 uint32
	CountVehicleClass4 uint32
	CountVehicleClass5 uint32
	CountVehicleClass6 uint32
	CountVehicleClass7 uint32
	CountVehicleClass8 uint32

	CountVehicleClassEquivalent uint32

	HarmonicMeanSpeedVehicleClass1 float32
	HarmonicMeanSpeedVehicleClass2 float32
	HarmonicMeanSpeedVehicleClass3 float32
	HarmonicMeanSpeedVehicleClass4 float32
	HarmonicMeanSpeedVehicleClass5 float32
	HarmonicMeanSpeedVehicleClass6 float32
	HarmonicMeanSpeedVehicleClass7 float32
	HarmonicMeanSpeedVehicleClass8 float32

	HarmonicMeanSpeedVehicleClassAll float32

	CountDetectedSpeedVehicleUnder50       uint32
	CountDetectedSpeedVehicleBetween50_100 uint32
	CountDetectedSpeedVehicleOver100       uint32

	AverageVehicleLength float32

	AverageHeadway float32
	StdHeadway     float32

	AverageTimeToCollision float32
	StdTimeToCollision     float32
}

type CurrentTrafficDataByDetectionPointStatisticsByDatasource struct {
	DataSourceName             string
	RecordsCount               uint32
	FirstRecordTimestamp       time.Time
	LastRecordTimestamp        time.Time
	StatisticsByDetectionPoint []CurrentTrafficDataByDetectionPointStatistics
}

type CurrentTrafficDataByDetectionPointStatistics struct {
	DetectionPointID             string
	RecordsCount                 uint32
	FirstRecordTimestamp         time.Time
	LastRecordTimestamp          time.Time
	LongestDataGap               float32
	LongestDataGapStartTimestamp time.Time
	LongestDataGapEndTimestamp   time.Time
	MissingIntervalDetectionRate float32 // Percentage of expected intervals that are missing data.
	//total
	TotalCountEquivalentVehicles     uint32  // all counted vehicles
	TotalHarmonicMeanSpeedAllRecords float32 // mean speed of all means
	// rate
	PercentageRecordsWithEquivalentCounts     float32 // how many records have at least one count
	PercentageRecordsWithPositiveAverageSpeed float32 // how many records have an average speed more than 0
}
