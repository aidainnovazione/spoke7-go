package models

import "time"

type DataSourceType string

const (
	Simulator DataSourceType = "simulator"
	Real      DataSourceType = "real"
)

// Fonte dati
type DataSource struct {
	Name              string
	Description       string
	Type              DataSourceType
	RoadNetworkId     *string
	Owner             string
	Groups            []string
	ModifiedBy        string
	CreatedAt         time.Time
	ModifiedAt        time.Time
	DetectionSections []DetectionSection
	DetectionPoints   []DetectionPoint
}

type UpdateDataSource struct {
	Name          string
	Description   *string
	Type          *DataSourceType
	RoadNetworkId *string
	Owner         *string
	Groups        []string
	ModifiedBy    *string
}

type DataSourceListParams struct {
	DetectionSections bool
	DetectionPoints   bool
}

type DataSourceGetParams struct {
	DetectionSections bool
	DetectionPoints   bool
}
