package dao

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"spoke7-go/internal/metadata/models"
	"time"

	"github.com/lib/pq"
)

// DAO for Dashboard
type Dashboard struct {
	ID             string                `gorm:"primaryKey"`
	Name           string                `gorm:"type:text"`
	Description    string                `gorm:"type:text"`
	DataSourceName string                `gorm:"type:text"`
	CreatedAt      time.Time             `gorm:"autoCreateTime:milli"`
	ModifiedAt     time.Time             `gorm:"autoUpdateTime:milli"`
	Section        DashboardSectionArray `gorm:"type:jsonb" json:"section,omitempty"`
	Owner          string
	Groups         pq.StringArray     `gorm:"type:text[]"`
	Bottombar      DashboardBottombar `gorm:"type:jsonb" json:"bottombar,omitempty"`
	Sidebar        DashboardSidebar   `gorm:"type:jsonb" json:"sidebar,omitempty"`
	BaseMap        DashboardBaseMap   `gorm:"type:jsonb" json:"basemap,omitempty"`
}

// TableName sets the table name for the DataSource struct
func (Dashboard) TableName() string {
	return "dashboard"
}

func (dao *Dashboard) ToModel() models.Dashboard {
	modelSections := make([]models.Section, len(dao.Section))
	for i, section := range dao.Section {
		modelSections[i] = models.Section{
			Name:       section.Name,
			IsVisible:  section.IsVisible,
			IsExpanded: section.IsExpanded,
			LayerType:  models.LayerType(section.LayerType),
			Params:     section.Params,
		}
	}
	modelSidebar := models.Sidebar{
		IsOpen:            dao.Sidebar.IsOpen,
		InjectedComponent: dao.Sidebar.InjectedComponent,
	}
	modelBaseMap := models.BaseMap{
		Zoom:     dao.BaseMap.Zoom,
		Center:   dao.BaseMap.Center,
		MapTheme: models.MapTheme(dao.BaseMap.MapTheme),
	}

	modelBottombar := models.Bottombar{
		IsOpen:       dao.Bottombar.IsOpen,
		SelectedTime: dao.Bottombar.SelectedTime,
		StartTime:    dao.Bottombar.StartTime,
		EndTime:      dao.Bottombar.EndTime,
		Interval:     dao.Bottombar.Interval,
	}

	return models.Dashboard{
		Name:           dao.Name,
		Description:    dao.Description,
		ID:             dao.ID,
		DataSourceName: dao.DataSourceName,
		CreatedAt:      dao.CreatedAt,
		ModifiedAt:     dao.ModifiedAt,
		Sections:       modelSections,
		Owner:          dao.Owner,
		Groups:         dao.Groups,
		Sidebar:        modelSidebar,
		Bottombar:      modelBottombar,
		BaseMap:        modelBaseMap,
	}
}

func NewDashboardDaoFromModel(model models.Dashboard) Dashboard {
	var sections DashboardSectionArray
	for _, section := range model.Sections {
		sections = append(sections, DashboardSection{
			Name:       section.Name,
			IsVisible:  section.IsVisible,
			IsExpanded: section.IsExpanded,
			LayerType:  LayerType(section.LayerType),
			Params:     section.Params,
		})
	}
	bottombar := DashboardBottombar{
		IsOpen:       model.Bottombar.IsOpen,
		SelectedTime: model.Bottombar.SelectedTime,
		StartTime:    model.Bottombar.StartTime,
		EndTime:      model.Bottombar.EndTime,
		Interval:     model.Bottombar.Interval,
	}
	sidebar := DashboardSidebar{
		IsOpen:            model.Sidebar.IsOpen,
		InjectedComponent: model.Sidebar.InjectedComponent,
	}
	baseMap := DashboardBaseMap{
		Zoom:     model.BaseMap.Zoom,
		Center:   model.BaseMap.Center,
		MapTheme: MapTheme(model.BaseMap.MapTheme),
	}
	return Dashboard{
		ID:             model.ID,
		Name:           model.Name,
		Description:    model.Description,
		DataSourceName: model.DataSourceName,
		Section:        sections,
		Owner:          model.Owner,
		Groups:         model.Groups,
		Bottombar:      bottombar,
		Sidebar:        sidebar,
		BaseMap:        baseMap,
		// CreatedAt:         nil,
		// ModifiedAt:        nil,
	}
}

// section
type LayerType int

const (
	LAYER_TYPE_CURRENT_TRAFFIC_LANE LayerType = iota
	LAYER_TYPE_CURRENT_TRAFFIC_POINT
	LAYER_TYPE_CURRENT_TRAFFIC_SECTION
	LAYER_TYPE_NETWORK
	LAYER_TYPE_SPIRA
)

type DashboardSectionArray []DashboardSection

type DashboardSection struct {
	Name       string
	IsVisible  bool
	IsExpanded bool
	LayerType  LayerType
	Params     map[string]interface{}
}

func (p DashboardSectionArray) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DashboardSectionArray: %w", err)
	}
	return data, nil
}

func (p *DashboardSectionArray) Scan(value interface{}) error {
	if value == nil {
		*p = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DashboardSectionArray: expected []byte but got %T", value)
	}
	return json.Unmarshal(bytes, p)
}

// bottombar

type DashboardBottombar struct {
	IsOpen       bool
	SelectedTime time.Time
	StartTime    time.Time
	EndTime      time.Time
	Interval     int64
}

func (p DashboardBottombar) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DashboardBottombar: %w", err)
	}
	return data, nil
}

func (p *DashboardBottombar) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DashboardBottombar: expected []byte but got %T", value)
	}
	return json.Unmarshal(bytes, p)
}

// sidebar

type DashboardSidebar struct {
	IsOpen            bool
	InjectedComponent string
}

func (p DashboardSidebar) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DashboardSidebar: %w", err)
	}
	return data, nil
}

func (p *DashboardSidebar) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DashboardSidebar: expected []byte but got %T", value)
	}
	return json.Unmarshal(bytes, p)
}

// base map

type MapTheme int

const (
	LIGHT MapTheme = iota
	DARK
)

type DashboardBaseMap struct {
	MapTheme MapTheme // possible values: "light", "dark"
	Center   []float64
	Zoom     float64
}

func (p DashboardBaseMap) Value() (driver.Value, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DashboardBaseMap: %w", err)
	}
	return data, nil
}

func (p *DashboardBaseMap) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan DashboardBaseMap: expected []byte but got %T", value)
	}
	return json.Unmarshal(bytes, p)
}
