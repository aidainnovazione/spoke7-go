package models

type InductionLoop struct {
	ID     string
	LaneID string
	Pos    float64
	Period float64
	File   string
	AbsX   float64
	AbsY   float64
	Shape  [][2]float64
	EdgeID string
}
