package models

import (
	"testing"
	"time"
)

func TestDataSource(t *testing.T) {
	now := time.Now()

	dataSource := DataSource{
		Name:              "Test Source",
		Description:       "This is a test source",
		Type:              Simulator,
		RoadNetworkId:     nil,
		Owner:             "",
		Groups:            []string{},
		ModifiedBy:        "",
		CreatedAt:         time.Time{},
		ModifiedAt:        time.Time{},
		DetectionSections: []DetectionSection{},
		DetectionPoints:   []DetectionPoint{},
	}
	if dataSource.Description != "This is a test source" {
		t.Errorf("expected Description to be 'This is a test source', got %s", dataSource.Description)
	}
	if dataSource.Type != Simulator {
		t.Errorf("expected Type to be 'simulator', got %s", dataSource.Type)
	}
	if dataSource.RoadNetworkId != nil {
		t.Errorf("expected RoadNetworkId to be '12345', got %s", dataSource.RoadNetworkId)
	}
	if dataSource.Owner != "Test Owner" {
		t.Errorf("expected Owner to be 'Test Owner', got %s", dataSource.Owner)
	}
	if dataSource.Groups != nil {
		t.Errorf("expected Groups to be 'Test Groups', got %s", dataSource.Groups)
	}
	if dataSource.ModifiedBy != "Test Modifier" {
		t.Errorf("expected ModifiedBy to be 'Test Modifier', got %s", dataSource.ModifiedBy)
	}
	if !dataSource.CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt to be '%v', got %v", now, dataSource.CreatedAt)
	}
	if !dataSource.ModifiedAt.Equal(now) {
		t.Errorf("expected ModifiedAt to be '%v', got %v", now, dataSource.ModifiedAt)
	}
}

func TestUpdateDataSource(t *testing.T) {
	description := "Updated description"
	sourceType := Real
	roadNetworkId := "67890"
	owner := "Updated Owner"
	groups := []string{"group1"}
	modifiedBy := "Updated Modifier"

	updateDataSource := UpdateDataSource{
		Name:          "Updated Source",
		Description:   &description,
		Type:          &sourceType,
		RoadNetworkId: &roadNetworkId,
		Owner:         &owner,
		Groups:        groups,
		ModifiedBy:    &modifiedBy,
	}

	if updateDataSource.Name != "Updated Source" {
		t.Errorf("expected Name to be 'Updated Source', got %s", updateDataSource.Name)
	}
	if *updateDataSource.Description != "Updated description" {
		t.Errorf("expected Description to be 'Updated description', got %s", *updateDataSource.Description)
	}
	if *updateDataSource.Type != Real {
		t.Errorf("expected Type to be 'real', got %s", *updateDataSource.Type)
	}
	if *updateDataSource.RoadNetworkId != "67890" {
		t.Errorf("expected RoadNetworkId to be '67890', got %s", *updateDataSource.RoadNetworkId)
	}
	if *updateDataSource.Owner != "Updated Owner" {
		t.Errorf("expected Owner to be 'Updated Owner', got %s", *updateDataSource.Owner)
	}
	if updateDataSource.Groups != nil {
		t.Errorf("expected Groups to be 'Updated Groups', got %s", updateDataSource.Groups)
	}
	if *updateDataSource.ModifiedBy != "Updated Modifier" {
		t.Errorf("expected ModifiedBy to be 'Updated Modifier', got %s", *updateDataSource.ModifiedBy)
	}
}

func TestDataSource_DetectionSections(t *testing.T) {
	now := time.Now()
	dataSource := DataSource{
		Name: "Test Source",
		DetectionSections: []DetectionSection{
			{
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
			},
		},
	}

	if dataSource.Name != "Test Source" {
		t.Errorf("expected Name to be 'Test Source', got %s", dataSource.Name)
	}

	if len(dataSource.DetectionSections) != 1 {
		t.Errorf("expected DetectionSections to have 1 element, got %d", len(dataSource.DetectionSections))
	}

	if dataSource.DetectionSections[0].Id != "1" {
		t.Errorf("expected Id to be '1', got %s", dataSource.DetectionSections[0].Id)
	}
	if dataSource.DetectionSections[0].DataSourceName != "Source1" {
		t.Errorf("expected DataSourceName to be 'Source1', got %s", dataSource.DetectionSections[0].DataSourceName)
	}
	if dataSource.DetectionSections[0].Description != "Description1" {
		t.Errorf("expected Description to be 'Description1', got %s", dataSource.DetectionSections[0].Description)
	}
	if dataSource.DetectionSections[0].StartLatitude != 1.0 {
		t.Errorf("expected StartLatitude to be 1.0, got %f", dataSource.DetectionSections[0].StartLatitude)
	}
	if dataSource.DetectionSections[0].StartLongitude != 1.0 {
		t.Errorf("expected StartLongitude to be 1.0, got %f", dataSource.DetectionSections[0].StartLongitude)
	}
	if dataSource.DetectionSections[0].EndLatitude != 2.0 {
		t.Errorf("expected EndLatitude to be 2.0, got %f", dataSource.DetectionSections[0].EndLatitude)
	}
	if dataSource.DetectionSections[0].EndLongitude != 2.0 {
		t.Errorf("expected EndLongitude to be 2.0, got %f", dataSource.DetectionSections[0].EndLongitude)
	}
	if dataSource.DetectionSections[0].Direction != 1 {
		t.Errorf("expected Direction to be 1, got %d", dataSource.DetectionSections[0].Direction)
	}
	if dataSource.DetectionSections[0].RoadNetworkId != "RN123" {
		t.Errorf("expected RoadNetworkId to be 'RN123', got %s", dataSource.DetectionSections[0].RoadNetworkId)
	}
	if !dataSource.DetectionSections[0].CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt to be %v, got %v", now, dataSource.DetectionSections[0].CreatedAt)
	}
	if !dataSource.DetectionSections[0].ModifiedAt.Equal(now) {
		t.Errorf("expected ModifiedAt to be %v, got %v", now, dataSource.DetectionSections[0].ModifiedAt)
	}
}

func TestDataSource_DetectionPoints(t *testing.T) {
	now := time.Now()
	dataSource := DataSource{
		Name: "Test Source",
		DetectionPoints: []DetectionPoint{
			{
				Id: "1",

				Description: "Main Street Camera",

				CreatedAt:  now,
				ModifiedAt: now,
			},
		},
	}

	if dataSource.Name != "Test Source" {
		t.Errorf("expected Name to be 'Test Source', got %s", dataSource.Name)
	}

	if len(dataSource.DetectionPoints) != 1 {
		t.Errorf("expected DetectionPoints to have 1 element, got %d", len(dataSource.DetectionPoints))
	}

	if dataSource.DetectionPoints[0].Id != "1" {
		t.Errorf("expected Id to be '1', got %s", dataSource.DetectionPoints[0].Id)
	}

	if dataSource.DetectionPoints[0].Description != "Main Street Camera" {
		t.Errorf("expected Description to be 'Main Street Camera', got %s", dataSource.DetectionPoints[0].Description)
	}

	if !dataSource.DetectionPoints[0].CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt to be %v, got %v", now, dataSource.DetectionPoints[0].CreatedAt)
	}
	if !dataSource.DetectionPoints[0].ModifiedAt.Equal(now) {
		t.Errorf("expected ModifiedAt to be %v, got %v", now, dataSource.DetectionPoints[0].ModifiedAt)
	}
}
