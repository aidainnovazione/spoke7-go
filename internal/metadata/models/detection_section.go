package models

import "time"

// DetectionSection(Tratta) represents a section of the road network delimited by two detection points.
type DetectionSection struct {
	Id                           string
	DataSourceName               string
	Description                  string
	StartLatitude                float64
	StartLongitude               float64
	EndLatitude                  float64
	EndLongitude                 float64
	Direction                    int
	Shape                        interface{}
	RoadNetworkId                string
	CreatedAt                    time.Time
	ModifiedAt                   time.Time
	DetectionSectionRoadNetworks []DetectionSectionRoadNetwork
}

// DetectionSectionRoadNetwork represents a section of the road network delimited by two detection points.
type DetectionSectionRoadNetwork struct {
	Id                 string
	DetectionSectionId string
	RoadNetworkId      string
	StartMeters        float64
	EndMeters          float64
	CreatedAt          time.Time
	ModifiedAt         time.Time
}
