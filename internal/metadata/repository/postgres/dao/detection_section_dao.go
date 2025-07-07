package dao

import (
	"spoke7-go/internal/metadata/models"
	"time"
)

type DetectionSectionRoadNetwork struct {
	Id                 string `gorm:"primaryKey"`
	DataSourceName     string `gorm:"primaryKey;index;constraint:OnDelete:CASCADE,OnUpdate:CASCADE,foreignKey:DataSourceName;references:Name"`
	DetectionSectionId string
	RoadNetworkId      string
	StartMeters        float64
	EndMeters          float64
	CreatedAt          time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt         time.Time `gorm:"autoUpdateTime:milli"`
}

func (dao DetectionSectionRoadNetwork) TableName() string {
	return "detection_section_road_networks"
}

func (dao *DetectionSectionRoadNetwork) ToModel() models.DetectionSectionRoadNetwork {
	return models.DetectionSectionRoadNetwork{
		Id:                 dao.Id,
		DetectionSectionId: dao.DetectionSectionId,
		RoadNetworkId:      dao.RoadNetworkId,
		StartMeters:        dao.StartMeters,
		EndMeters:          dao.EndMeters,
		CreatedAt:          dao.CreatedAt,
		ModifiedAt:         dao.ModifiedAt,
	}
}

func NewDetectionSectionRoadNetworkDaoFromModel(model models.DetectionSectionRoadNetwork) DetectionSectionRoadNetwork {
	return DetectionSectionRoadNetwork{
		Id:                 model.Id,
		DetectionSectionId: model.DetectionSectionId,
		RoadNetworkId:      model.RoadNetworkId,
		StartMeters:        model.StartMeters,
		EndMeters:          model.EndMeters,
		CreatedAt:          model.CreatedAt,
		ModifiedAt:         model.ModifiedAt,
	}
}

type DetectionSection struct {
	Id             string `gorm:"primaryKey"`
	DataSourceName string
	Description    string
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
	Direction      int
	// Shape          interface{}
	RoadNetworkId string
	CreatedAt     time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt    time.Time `gorm:"autoUpdateTime:milli"`

	DetectionSectionRoadNetworks []DetectionSectionRoadNetwork
}

func (dao DetectionSection) TableName() string {
	return "detection_sections"
}

func (dao *DetectionSection) ToModel() models.DetectionSection {
	detectionSectionRoadNetworks := make([]models.DetectionSectionRoadNetwork, 0, len(dao.DetectionSectionRoadNetworks))
	for _, detectionSectionRoadNetwork := range dao.DetectionSectionRoadNetworks {
		detectionSectionRoadNetworks = append(detectionSectionRoadNetworks, detectionSectionRoadNetwork.ToModel())
	}

	return models.DetectionSection{
		Id:             dao.Id,
		DataSourceName: dao.DataSourceName,
		Description:    dao.Description,
		StartLatitude:  dao.StartLatitude,
		StartLongitude: dao.StartLongitude,
		EndLatitude:    dao.EndLatitude,
		EndLongitude:   dao.EndLongitude,
		Direction:      dao.Direction,
		// Shape:          dao.Shape,
		RoadNetworkId:                dao.RoadNetworkId,
		CreatedAt:                    dao.CreatedAt,
		ModifiedAt:                   dao.ModifiedAt,
		DetectionSectionRoadNetworks: detectionSectionRoadNetworks,
	}
}

func NewDetectionSectionDaoFromModel(model models.DetectionSection) DetectionSection {
	detectionSectionRoadNetworks := make([]DetectionSectionRoadNetwork, 0, len(model.DetectionSectionRoadNetworks))
	for _, detectionSectionRoadNetwork := range model.DetectionSectionRoadNetworks {
		detectionSectionRoadNetworks = append(detectionSectionRoadNetworks, NewDetectionSectionRoadNetworkDaoFromModel(detectionSectionRoadNetwork))
	}
	return DetectionSection{
		Id:             model.Id,
		DataSourceName: model.DataSourceName,
		Description:    model.Description,
		StartLatitude:  model.StartLatitude,
		StartLongitude: model.StartLongitude,
		EndLatitude:    model.EndLatitude,
		EndLongitude:   model.EndLongitude,
		Direction:      model.Direction,
		// Shape:          model.Shape,
		RoadNetworkId:                model.RoadNetworkId,
		DetectionSectionRoadNetworks: detectionSectionRoadNetworks,
	}
}
