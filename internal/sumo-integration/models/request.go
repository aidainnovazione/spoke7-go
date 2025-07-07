package models

import "time"

type SumoIntegrationCreateCurrentTrafficDataByDetectionPointByLaneFromXmlRequestModel struct {
	DataSourceName string
	StartTime      time.Time
	Xml            XmlFile
}

type SumoIntegrationCreateCurrentTrafficDataByDetectionSectionFromXmlRequestModel struct {
	DataSourceName string
	StartTime      time.Time
	Xml            XmlFile
}

type XmlFile struct {
	Filename    string
	ContentType string
	Content     []byte
}
