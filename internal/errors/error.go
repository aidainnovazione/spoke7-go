package errors

import "errors"

var ErrUnauthorized = errors.New("unauthorized access")
var ErrForbidden = errors.New("forbidden access")
var ErrNotFound = errors.New("resource not found")
var ErrBadRequest = errors.New("bad request")
var ErrMissingDataSourceName = errors.New("missing data source name")

// DATASOURCE ERRORS
var ErrDataSourceNotFound = errors.New("data source not found")
var ErrWrongDataSourceRequest = errors.New("wrong data source request")

// DETECTION POINT ERRORS
var ErrDetectionPointNotFound = errors.New("detection point not found")
var ErrWrongDetectionPointRequest = errors.New("wrong detection point request")
var ErrMissingDetectionPointID = errors.New("missing detection point ID")

// DETECTION SECTION ERRORS
var ErrDetectionSectionNotFound = errors.New("detection section not found")
var ErrWrongDetectionSectionRequest = errors.New("wrong detection section request")
var ErrMissingDetectionSectionID = errors.New("missing detection section ID")

// CURRENT TRAFFIC
var ErrWrongCurrentTrafficDetectionPointRequest = errors.New("wrong current traffic by detection point request")
var ErrWrongCurrentTrafficDetectionPointByLaneRequest = errors.New("wrong current traffic by detection point by lanerequest")
var ErrWrongCurrentTrafficDetectionSectionRequest = errors.New("wrong current traffic by detection section request")

// REAL TIME
var ErrWrongRealTimeTrafficDetectionPointByLaneRequest = errors.New("wrong real time traffic by detection point by lanerequest")
var ErrWrongRealTimeTrafficDetectionSectionRequest = errors.New("wrong real time traffic by detection section request")

// HISTORY DAY
var ErrWrongHistoryDayTrafficDetectionPointRequest = errors.New("wrong history day traffic by detection point request")
var ErrWrongHistoryDayTrafficDetectionPointByLaneRequest = errors.New("wrong history day traffic by detection point by lanerequest")
var ErrWrongHistoryDayTrafficDetectionSectionRequest = errors.New("wrong history day traffic by detection section request")

// HISTORY HOUR
var ErrWrongHistoryHourTrafficDetectionPointRequest = errors.New("wrong history hour traffic by detection point request")
var ErrWrongHistoryHourTrafficDetectionPointByLaneRequest = errors.New("wrong history hour traffic by detection point by lanerequest")
var ErrWrongHistoryHourTrafficDetectionSectionRequest = errors.New("wrong history hour traffic by detection section request")

// STORED FILE
var ErrWrongStoredFileRequest = errors.New("wrong data source request")
var ErrMissingStoredFileContent = errors.New("missing file content")
var ErrMissingStoredFileID = errors.New("missing stored file id")
