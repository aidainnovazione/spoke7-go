package models

import (
	"testing"
	"time"
)

func TestDetectionPoint(t *testing.T) {
	now := time.Now()
	dp := DetectionPoint{
		Id: "1",

		Description: "Main Street Camera",

		CreatedAt:  now,
		ModifiedAt: now,
	}

	if dp.Id != "1" {
		t.Errorf("expected Id to be '1', got %s", dp.Id)
	}

	if dp.Description != "Main Street Camera" {
		t.Errorf("expected Description to be 'Main Street Camera', got %s", dp.Description)
	}

	if !dp.CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt to be %v, got %v", now, dp.CreatedAt)
	}
	if !dp.ModifiedAt.Equal(now) {
		t.Errorf("expected ModifiedAt to be %v, got %v", now, dp.ModifiedAt)
	}
}

func TestLane(t *testing.T) {
	now := time.Now()
	lane := Lane{
		Id: "1",

		CreatedAt:  now,
		ModifiedAt: now,
	}

	if lane.Id != "1" {
		t.Errorf("expected Id to be '1', got %s", lane.Id)
	}

	if !lane.CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt to be %v, got %v", now, lane.CreatedAt)
	}
	if !lane.ModifiedAt.Equal(now) {
		t.Errorf("expected ModifiedAt to be %v, got %v", now, lane.ModifiedAt)
	}
}
