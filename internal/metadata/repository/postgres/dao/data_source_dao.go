package dao

import (
	"spoke7-go/internal/metadata/models"
	"time"

	"github.com/lib/pq"
)

// DAO for DataSource
type DataSource struct {
	Name          string `gorm:"primaryKey"`
	Description   string `gorm:"type:text"`
	Type          string
	RoadNetworkId *string      `gorm:"index"` // Foreign key reference
	RoadNetwork   *RoadNetwork `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE;foreignKey:RoadNetworkId;references:ID"`
	CreatedAt     time.Time    `gorm:"autoCreateTime:milli"`
	ModifiedAt    time.Time    `gorm:"autoUpdateTime:milli"`

	Owner      string
	Groups     pq.StringArray `gorm:"type:text[]"`
	ModifiedBy string
	// relations
	DetectionSections []DetectionSection `gorm:"foreignKey:DataSourceName;references:Name"`
	DetectionPoints   []DetectionPoint   `gorm:"foreignKey:DataSourceName;references:Name"`
}

// TableName sets the table name for the DataSource struct
func (DataSource) TableName() string {
	return "data_source"
}

func (dao *DataSource) ToModel() models.DataSource {
	detectionSections := make([]models.DetectionSection, len(dao.DetectionSections))
	for i, detectionSection := range dao.DetectionSections {
		detectionSections[i] = detectionSection.ToModel()
	}

	detectionPoints := make([]models.DetectionPoint, len(dao.DetectionPoints))
	for i, detectionPoint := range dao.DetectionPoints {
		detectionPoints[i] = FromDetectionPointDaoToModel(detectionPoint)
	}

	return models.DataSource{
		Name:              dao.Name,
		Description:       dao.Description,
		Type:              models.DataSourceType(dao.Type),
		RoadNetworkId:     dao.RoadNetworkId,
		Owner:             dao.Owner,
		Groups:            dao.Groups,
		ModifiedBy:        dao.ModifiedBy,
		CreatedAt:         dao.CreatedAt,
		ModifiedAt:        dao.ModifiedAt,
		DetectionSections: detectionSections,
		DetectionPoints:   detectionPoints,
	}
}

func NewDataSourceDaoFromModel(model models.DataSource) DataSource {
	detectionSections := make([]DetectionSection, len(model.DetectionSections))
	for i, detectionSection := range model.DetectionSections {
		detectionSections[i] = NewDetectionSectionDaoFromModel(detectionSection)
	}

	detectionPoints := make([]DetectionPoint, len(model.DetectionPoints))
	for i, detectionPoint := range model.DetectionPoints {
		detectionPoints[i] = FromDetectionPointModelToDao(model.Name, detectionPoint)
	}
	return DataSource{
		Name:          model.Name,
		Description:   model.Description,
		Type:          string(model.Type),
		RoadNetworkId: model.RoadNetworkId,
		Owner:         model.Owner,
		Groups:        model.Groups,
		ModifiedBy:    model.ModifiedBy,
		// CreatedAt:         nil,
		// ModifiedAt:        nil,
		DetectionSections: detectionSections,
		DetectionPoints:   detectionPoints,
	}
}

func NewUpdateDataSourceDaoFromModel(model models.UpdateDataSource) DataSource {
	dataSourceType := ""
	if model.Type != nil {
		dataSourceType = string(*model.Type)
	}

	owner := ""
	if model.Owner != nil {
		owner = *model.Owner
	}

	groups := []string{}
	if model.Groups != nil {
		groups = model.Groups
	}

	modifiedBy := ""
	if model.ModifiedBy != nil {
		modifiedBy = *model.ModifiedBy
	}

	return DataSource{
		Name:          model.Name,
		Description:   *model.Description,
		Type:          dataSourceType,
		RoadNetworkId: model.RoadNetworkId,
		Owner:         owner,
		Groups:        groups,
		ModifiedBy:    modifiedBy,
	}
}
