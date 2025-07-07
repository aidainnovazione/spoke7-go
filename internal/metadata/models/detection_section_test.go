package models

import (
	"testing"
	"time"
)

func TestDetectionSection(t *testing.T) {
	now := time.Now()
	detectionSection := DetectionSection{
		Id:             "1",
		DataSourceName: "Source1",
		Description:    "Description1",
		StartLatitude:  1.0,
		StartLongitude: 1.0,
		EndLatitude:    2.0,
		EndLongitude:   2.0,
		Direction:      1,
		Shape:          nil,
		RoadNetworkId:  "RN123",
		CreatedAt:      now,
		ModifiedAt:     now,
	}

	if detectionSection.Id != "1" {
		t.Errorf("Id is not set correctly")
	}

	if detectionSection.DataSourceName != "Source1" {
		t.Errorf("DataSourceName is not set correctly")
	}

	if detectionSection.Description != "Description1" {
		t.Errorf("Description is not set correctly")
	}

	if detectionSection.StartLatitude != 1.0 {
		t.Errorf("StartLatitude is not set correctly")
	}

	if detectionSection.StartLongitude != 1.0 {
		t.Errorf("StartLongitude is not set correctly")
	}

	if detectionSection.EndLatitude != 2.0 {
		t.Errorf("EndLatitude is not set correctly")
	}

	if detectionSection.EndLongitude != 2.0 {
		t.Errorf("EndLongitude is not set correctly")
	}

	if detectionSection.Direction != 1 {
		t.Errorf("Direction is not set correctly")
	}

	if detectionSection.Shape != nil {
		t.Errorf("Shape is not set correctly")
	}

	if detectionSection.RoadNetworkId != "RN123" {
		t.Errorf("RoadNetworkId is not set correctly")
	}

	if detectionSection.CreatedAt != now {
		t.Errorf("CreatedAt is not set correctly")
	}

	if detectionSection.ModifiedAt != now {
		t.Errorf("ModifiedAt is not set correctly")
	}

}

func TestDetectionSectionRoadNetwork(t *testing.T) {
	now := time.Now()
	detectionSectionRoadNetwork := DetectionSectionRoadNetwork{
		Id:                 "1",
		DetectionSectionId: "DS123",
		RoadNetworkId:      "RN123",
		StartMeters:        1.0,
		EndMeters:          2.0,
		CreatedAt:          now,
		ModifiedAt:         now,
	}

	if detectionSectionRoadNetwork.Id != "1" {
		t.Errorf("Id is not set correctly")
	}

	if detectionSectionRoadNetwork.DetectionSectionId != "DS123" {
		t.Errorf("DetectionSectionId is not set correctly")
	}

	if detectionSectionRoadNetwork.RoadNetworkId != "RN123" {
		t.Errorf("RoadNetworkId is not set correctly")
	}

	if detectionSectionRoadNetwork.StartMeters != 1.0 {
		t.Errorf("StartMeters is not set correctly")
	}

	if detectionSectionRoadNetwork.EndMeters != 2.0 {
		t.Errorf("EndMeters is not set correctly")
	}

	if detectionSectionRoadNetwork.CreatedAt != now {
		t.Errorf("CreatedAt is not set correctly")
	}

	if detectionSectionRoadNetwork.ModifiedAt != now {
		t.Errorf("ModifiedAt is not set correctly")
	}

}
