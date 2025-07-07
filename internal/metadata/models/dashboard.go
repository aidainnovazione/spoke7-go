package models

import "time"

// Dashboard
type Dashboard struct {
	ID             string
	Name           string
	Description    string
	DataSourceName string
	Sections       []Section
	CreatedAt      time.Time
	ModifiedAt     time.Time
	Owner          string
	Groups         []string
	Bottombar      Bottombar
	Sidebar        Sidebar
	BaseMap        BaseMap
}

// LayerType
type LayerType int

const (
	LAYER_TYPE_CURRENT_TRAFFIC_LANE LayerType = iota
	LAYER_TYPE_CURRENT_TRAFFIC_POINT
	LAYER_TYPE_CURRENT_TRAFFIC_SECTION
	LAYER_TYPE_NETWORK
	LAYER_TYPE_SPIRA
)

// Section
type Section struct {
	Name       string
	IsVisible  bool
	IsExpanded bool
	LayerType  LayerType
	Params     map[string]interface{}
}

// Bottombar
type Bottombar struct {
	IsOpen       bool
	SelectedTime time.Time
	StartTime    time.Time
	EndTime      time.Time
	Interval     int64
}

// Sidebar
type Sidebar struct {
	IsOpen            bool
	InjectedComponent string
}

type MapTheme int

const (
	LIGHT MapTheme = iota
	DARK
)

type BaseMap struct {
	MapTheme MapTheme // possible values: "light", "dark"
	Center   []float64
	Zoom     float64
}
