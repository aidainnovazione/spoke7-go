package dtos

import "encoding/xml"

type InductionDetectorModel struct {
	XMLName   xml.Name                 `xml:"detector"`
	Intervals []InductionIntervalModel `xml:"interval"`
}

type InductionIntervalModel struct {
	Begin             float64 `xml:"begin,attr"`
	End               float64 `xml:"end,attr"`
	ID                string  `xml:"id,attr"`
	NVehContrib       int     `xml:"nVehContrib,attr"`
	Flow              float64 `xml:"flow,attr"`
	Occupancy         float64 `xml:"occupancy,attr"`
	Speed             float64 `xml:"speed,attr"`
	HarmonicMeanSpeed float64 `xml:"harmonicMeanSpeed,attr"`
	Length            float64 `xml:"length,attr"`
	NVehEntered       int     `xml:"nVehEntered,attr"`
}
