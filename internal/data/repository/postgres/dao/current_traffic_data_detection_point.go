package dao

import (
	"time"

	"spoke7-go/internal/data/models"
)

// CurrentTrafficDataByDetectionPoint represents the table current_traffic_data_by_detection_point
type CurrentTrafficDataByDetectionPointDao struct {
	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`

	DataSourceName     string    `gorm:"primary_key;index;not null"`
	DetectionPointID   string    `gorm:"primary_key;type:varchar(255);index;not null"`
	DetectionTimestamp time.Time `gorm:"primary_key;type:timestamp;index;not null"`
	DetectionInterval  uint32    `gorm:"type:integer;not null"`

	CountVehicleClass1 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass2 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass3 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass4 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass5 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass6 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass7 uint32 `gorm:"type:integer;not null"`
	CountVehicleClass8 uint32 `gorm:"type:integer;not null"`

	CountVehicleClassEquivalent uint32 `gorm:"type:integer;not null"`

	HarmonicMeanSpeedVehicleClass1 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass2 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass3 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass4 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass5 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass6 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass7 float32 `gorm:"type:float;not null"`
	HarmonicMeanSpeedVehicleClass8 float32 `gorm:"type:float;not null"`

	HarmonicMeanSpeedVehicleClassAll float32 `gorm:"type:float;not null"`

	CountDetectedSpeedVehicleUnder50       uint32 `gorm:"type:integer;not null"`
	CountDetectedSpeedVehicleBetween50_100 uint32 `gorm:"type:integer;not null"`
	CountDetectedSpeedVehicleOver100       uint32 `gorm:"type:integer;not null"`

	AverageVehicleLength float32 `gorm:"type:float;not null"`

	AverageHeadway float32 `gorm:"type:float;not null"`
	StdHeadway     float32 `gorm:"type:float;not null"`

	AverageTimeToCollision float32 `gorm:"type:float;not null"`
	StdTimeToCollision     float32 `gorm:"type:float;not null"`
}

// TableName returns the table name of the CurrentTrafficDataByDetectionPoint
func (CurrentTrafficDataByDetectionPointDao) TableName() string {
	return "current_traffic_data_by_detection_point"
}

func (c CurrentTrafficDataByDetectionPointDao) FromCurrentTrafficDataByDetectionPointDaoToModel() models.CurrentTrafficDataByDetectionPointModel {
	return models.CurrentTrafficDataByDetectionPointModel{
		CreatedAt:                              c.CreatedAt,
		ModifiedAt:                             c.ModifiedAt,
		DataSourceName:                         c.DataSourceName,
		DetectionPointID:                       c.DetectionPointID,
		DetectionTimestamp:                     c.DetectionTimestamp,
		DetectionInterval:                      c.DetectionInterval,
		CountVehicleClass1:                     c.CountVehicleClass1,
		CountVehicleClass2:                     c.CountVehicleClass2,
		CountVehicleClass3:                     c.CountVehicleClass3,
		CountVehicleClass4:                     c.CountVehicleClass4,
		CountVehicleClass5:                     c.CountVehicleClass5,
		CountVehicleClass6:                     c.CountVehicleClass6,
		CountVehicleClass7:                     c.CountVehicleClass7,
		CountVehicleClass8:                     c.CountVehicleClass8,
		CountVehicleClassEquivalent:            c.CountVehicleClassEquivalent,
		HarmonicMeanSpeedVehicleClass1:         c.HarmonicMeanSpeedVehicleClass1,
		HarmonicMeanSpeedVehicleClass2:         c.HarmonicMeanSpeedVehicleClass2,
		HarmonicMeanSpeedVehicleClass3:         c.HarmonicMeanSpeedVehicleClass3,
		HarmonicMeanSpeedVehicleClass4:         c.HarmonicMeanSpeedVehicleClass4,
		HarmonicMeanSpeedVehicleClass5:         c.HarmonicMeanSpeedVehicleClass5,
		HarmonicMeanSpeedVehicleClass6:         c.HarmonicMeanSpeedVehicleClass6,
		HarmonicMeanSpeedVehicleClass7:         c.HarmonicMeanSpeedVehicleClass7,
		HarmonicMeanSpeedVehicleClass8:         c.HarmonicMeanSpeedVehicleClass8,
		HarmonicMeanSpeedVehicleClassAll:       c.HarmonicMeanSpeedVehicleClassAll,
		CountDetectedSpeedVehicleUnder50:       c.CountDetectedSpeedVehicleUnder50,
		CountDetectedSpeedVehicleBetween50_100: c.CountDetectedSpeedVehicleBetween50_100,
		CountDetectedSpeedVehicleOver100:       c.CountDetectedSpeedVehicleOver100,
		AverageVehicleLength:                   c.AverageVehicleLength,
		AverageHeadway:                         c.AverageHeadway,
		StdHeadway:                             c.StdHeadway,
		AverageTimeToCollision:                 c.AverageTimeToCollision,
		StdTimeToCollision:                     c.StdTimeToCollision,
	}
}

func FromCurrentTrafficDataByDetectionPointModelToDao(traffic models.CurrentTrafficDataByDetectionPointModel) CurrentTrafficDataByDetectionPointDao {
	return CurrentTrafficDataByDetectionPointDao{
		CreatedAt:                              traffic.CreatedAt,
		ModifiedAt:                             traffic.ModifiedAt,
		DataSourceName:                         traffic.DataSourceName,
		DetectionPointID:                       traffic.DetectionPointID,
		DetectionTimestamp:                     traffic.DetectionTimestamp,
		DetectionInterval:                      traffic.DetectionInterval,
		CountVehicleClass1:                     traffic.CountVehicleClass1,
		CountVehicleClass2:                     traffic.CountVehicleClass2,
		CountVehicleClass3:                     traffic.CountVehicleClass3,
		CountVehicleClass4:                     traffic.CountVehicleClass4,
		CountVehicleClass5:                     traffic.CountVehicleClass5,
		CountVehicleClass6:                     traffic.CountVehicleClass6,
		CountVehicleClass7:                     traffic.CountVehicleClass7,
		CountVehicleClass8:                     traffic.CountVehicleClass8,
		CountVehicleClassEquivalent:            traffic.CountVehicleClassEquivalent,
		HarmonicMeanSpeedVehicleClass1:         traffic.HarmonicMeanSpeedVehicleClass1,
		HarmonicMeanSpeedVehicleClass2:         traffic.HarmonicMeanSpeedVehicleClass2,
		HarmonicMeanSpeedVehicleClass3:         traffic.HarmonicMeanSpeedVehicleClass3,
		HarmonicMeanSpeedVehicleClass4:         traffic.HarmonicMeanSpeedVehicleClass4,
		HarmonicMeanSpeedVehicleClass5:         traffic.HarmonicMeanSpeedVehicleClass5,
		HarmonicMeanSpeedVehicleClass6:         traffic.HarmonicMeanSpeedVehicleClass6,
		HarmonicMeanSpeedVehicleClass7:         traffic.HarmonicMeanSpeedVehicleClass7,
		HarmonicMeanSpeedVehicleClass8:         traffic.HarmonicMeanSpeedVehicleClass8,
		HarmonicMeanSpeedVehicleClassAll:       traffic.HarmonicMeanSpeedVehicleClassAll,
		CountDetectedSpeedVehicleUnder50:       traffic.CountDetectedSpeedVehicleUnder50,
		CountDetectedSpeedVehicleBetween50_100: traffic.CountDetectedSpeedVehicleBetween50_100,
		CountDetectedSpeedVehicleOver100:       traffic.CountDetectedSpeedVehicleOver100,
		AverageVehicleLength:                   traffic.AverageVehicleLength,
		AverageHeadway:                         traffic.AverageHeadway,
		StdHeadway:                             traffic.StdHeadway,
		AverageTimeToCollision:                 traffic.AverageTimeToCollision,
		StdTimeToCollision:                     traffic.StdTimeToCollision,
	}
}
