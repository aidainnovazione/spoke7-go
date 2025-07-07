package dtos

import "encoding/xml"

type LaneDetectorModel struct {
	XMLName   xml.Name            `xml:"detector"`
	Intervals []LaneIntervalModel `xml:"interval"`
}

type LaneIntervalModel struct {
	Begin                       float64 `xml:"begin,attr"`
	End                         float64 `xml:"end,attr"`
	ID                          string  `xml:"id,attr"`
	SampledSeconds              float64 `xml:"sampledSeconds,attr"`
	NVehEntered                 int     `xml:"nVehEntered,attr"`
	NVehLeft                    int     `xml:"nVehLeft,attr"`
	NVehSeen                    int     `xml:"nVehSeen,attr"`
	MeanSpeed                   float64 `xml:"meanSpeed,attr"`
	MeanTimeLoss                float64 `xml:"meanTimeLoss,attr"`
	MeanOccupancy               float64 `xml:"meanOccupancy,attr"`
	MaxOccupancy                float64 `xml:"maxOccupancy,attr"`
	MeanMaxJamLengthInVehicles  float64 `xml:"meanMaxJamLengthInVehicles,attr"`
	MeanMaxJamLengthInMeters    float64 `xml:"meanMaxJamLengthInMeters,attr"`
	MaxJamLengthInVehicles      int     `xml:"maxJamLengthInVehicles,attr"`
	MaxJamLengthInMeters        float64 `xml:"maxJamLengthInMeters,attr"`
	JamLengthInVehiclesSum      int     `xml:"jamLengthInVehiclesSum,attr"`
	JamLengthInMetersSum        float64 `xml:"jamLengthInMetersSum,attr"`
	MeanHaltingDuration         float64 `xml:"meanHaltingDuration,attr"`
	MaxHaltingDuration          float64 `xml:"maxHaltingDuration,attr"`
	HaltingDurationSum          float64 `xml:"haltingDurationSum,attr"`
	MeanIntervalHaltingDuration float64 `xml:"meanIntervalHaltingDuration,attr"`
	MaxIntervalHaltingDuration  float64 `xml:"maxIntervalHaltingDuration,attr"`
	IntervalHaltingDurationSum  float64 `xml:"intervalHaltingDurationSum,attr"`
	StartedHalts                float64 `xml:"startedHalts,attr"`
	MeanVehicleNumber           float64 `xml:"meanVehicleNumber,attr"`
	MaxVehicleNumber            int     `xml:"maxVehicleNumber,attr"`
}
