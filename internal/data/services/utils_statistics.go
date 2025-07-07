package services

import (
	"sort"
	"spoke7-go/internal/data/models"
	"time"
)

type DetectionTimestampGetter interface {
	GetDetectionTimestamp() time.Time
}

func computeCurrentLaneStatistics(listModels []*models.CurrentTrafficDataByDetectionPointByLaneModel) models.CurrentTrafficDataByDetectionPointByLaneStatistics {
	var statistics models.CurrentTrafficDataByDetectionPointByLaneStatistics
	statistics.RecordsCount = uint32(len(listModels))

	if len(listModels) == 0 {
		return statistics
	}

	// order by ascending DetectionTimestamp
	getTimestamp := func(m *models.CurrentTrafficDataByDetectionPointByLaneModel) time.Time {
		return m.DetectionTimestamp
	}

	first, last := computeCommonTimestamps(listModels, getTimestamp)
	statistics.FirstRecordTimestamp = first
	statistics.LastRecordTimestamp = last

	// detect longest data gap
	gap, gapStart, gapEnd := computeLongestDataGap(listModels, getTimestamp)
	statistics.LongestDataGap = float32(gap)
	statistics.LongestDataGapStartTimestamp = gapStart
	statistics.LongestDataGapEndTimestamp = gapEnd

	// compute missing interval detection rate
	statistics.MissingIntervalDetectionRate = computeMissingIntervalDetectionRate(len(listModels), first, last, 300)

	// compute total
	statistics.TotalCountAllVehicles = 0
	statistics.TotalHarmonicMeanSpeedAllRecords = 0
	for _, m := range listModels {
		statistics.TotalCountAllVehicles += m.CountVehicleClassAll
		statistics.TotalHarmonicMeanSpeedAllRecords += m.HarmonicMeanSpeedVehicleClassAll
	}
	statistics.TotalHarmonicMeanSpeedAllRecords /= float32(len(listModels))

	// compute rate
	statistics.PercentageRecordsWithCounts = 0
	statistics.PercentageRecordsWithPositiveAverageSpeed = 0
	for _, m := range listModels {
		if m.CountVehicleClassAll > 0 {
			statistics.PercentageRecordsWithCounts++
		}
		if m.HarmonicMeanSpeedVehicleClassAll > 0 {
			statistics.PercentageRecordsWithPositiveAverageSpeed++
		}
	}
	statistics.PercentageRecordsWithCounts = (float32(statistics.PercentageRecordsWithCounts) / float32(len(listModels)) * 100)
	statistics.PercentageRecordsWithPositiveAverageSpeed = (float32(statistics.PercentageRecordsWithPositiveAverageSpeed) / float32(len(listModels)) * 100)

	return statistics
}

func computeStatisticsDetectionSection[T DetectionTimestampGetter](listModels []*T) models.TrafficStatisticsDetectionSection {
	var statistics models.TrafficStatisticsDetectionSection

	statistics.RecordsCount = uint32(len(listModels))

	if len(listModels) == 0 {
		return statistics
	}

	// order by ascending DetectionTimestamp
	getTimestamp := func(m *T) time.Time {
		return (*m).GetDetectionTimestamp()
	}

	first, last := computeCommonTimestamps(listModels, getTimestamp)
	statistics.FirstRecordTimestamp = first
	statistics.LastRecordTimestamp = last

	// detect longest data gap
	gap, gapStart, gapEnd := computeLongestDataGap(listModels, getTimestamp)
	statistics.LongestDataGap = float32(gap)
	statistics.LongestDataGapStartTimestamp = gapStart
	statistics.LongestDataGapEndTimestamp = gapEnd

	return statistics
}

func computeStatisticsLane[T DetectionTimestampGetter](listModels []*T) models.TrafficStatisticsLane {
	var statistics models.TrafficStatisticsLane

	statistics.RecordsCount = uint32(len(listModels))

	if len(listModels) == 0 {
		return statistics
	}

	getTimestamp := func(m *T) time.Time {
		return (*m).GetDetectionTimestamp()
	}

	first, last := computeCommonTimestamps(listModels, getTimestamp)
	statistics.FirstRecordTimestamp = first
	statistics.LastRecordTimestamp = last

	// detect longest data gap
	gap, gapStart, gapEnd := computeLongestDataGap(listModels, getTimestamp)
	statistics.LongestDataGap = float32(gap)
	statistics.LongestDataGapStartTimestamp = gapStart
	statistics.LongestDataGapEndTimestamp = gapEnd

	return statistics
}

func computeStatisticsDetectionPoint[T DetectionTimestampGetter](listModels []*T) models.TrafficStatisticsDetectionPoint {
	var statistics models.TrafficStatisticsDetectionPoint

	statistics.RecordsCount = uint32(len(listModels))

	if len(listModels) == 0 {
		return statistics
	}

	getTimestamp := func(m *T) time.Time {
		return (*m).GetDetectionTimestamp()
	}

	first, last := computeCommonTimestamps(listModels, getTimestamp)
	statistics.FirstRecordTimestamp = first
	statistics.LastRecordTimestamp = last

	// detect longest data gap
	gap, gapStart, gapEnd := computeLongestDataGap(listModels, getTimestamp)
	statistics.LongestDataGap = float32(gap)
	statistics.LongestDataGapStartTimestamp = gapStart
	statistics.LongestDataGapEndTimestamp = gapEnd

	return statistics
}

func computeCurrentSectionStatistics(listModels []*models.CurrentTrafficDataByDetectionSectionModel) models.CurrentTrafficDataByDetectionSectionStatistics {
	var statistics models.CurrentTrafficDataByDetectionSectionStatistics

	statistics.RecordsCount = uint32(len(listModels))

	if len(listModels) == 0 {
		return statistics
	}

	// order by ascending DetectionTimestamp
	getTimestamp := func(m *models.CurrentTrafficDataByDetectionSectionModel) time.Time {
		return m.DetectionTimestamp
	}

	first, last := computeCommonTimestamps(listModels, getTimestamp)
	statistics.FirstRecordTimestamp = first
	statistics.LastRecordTimestamp = last

	// detect longest data gap
	gap, gapStart, gapEnd := computeLongestDataGap(listModels, getTimestamp)
	statistics.LongestDataGap = float32(gap)
	statistics.LongestDataGapStartTimestamp = gapStart
	statistics.LongestDataGapEndTimestamp = gapEnd

	// compute missing interval detection rate
	statistics.MissingIntervalDetectionRate = computeMissingIntervalDetectionRate(len(listModels), first, last, 300)

	// compute total
	statistics.TotalAverageForwardSpeed = 0
	statistics.TotalAverageBackwardSpeed = 0
	for _, m := range listModels {
		statistics.TotalAverageForwardSpeed += m.ForwardSpeed
		statistics.TotalAverageBackwardSpeed += m.BackwardSpeed
	}
	statistics.TotalAverageForwardSpeed /= float32(len(listModels))
	statistics.TotalAverageBackwardSpeed /= float32(len(listModels))

	// compute rate
	statistics.PercentageRecordsWithForwardSpeed = 0
	statistics.PercentageRecordsWithBackwardSpeed = 0
	for _, m := range listModels {
		if m.ForwardSpeed > 0 {
			statistics.PercentageRecordsWithForwardSpeed++
		}
		if m.BackwardSpeed > 0 {
			statistics.PercentageRecordsWithBackwardSpeed++
		}
	}
	statistics.PercentageRecordsWithForwardSpeed = (float32(statistics.PercentageRecordsWithForwardSpeed) / float32(len(listModels)) * 100)
	statistics.PercentageRecordsWithBackwardSpeed = (float32(statistics.PercentageRecordsWithBackwardSpeed) / float32(len(listModels)) * 100)

	return statistics
}

func computeCurrentDetectionPointStatistics(listModels []*models.CurrentTrafficDataByDetectionPointModel) models.CurrentTrafficDataByDetectionPointStatistics {
	var statistics models.CurrentTrafficDataByDetectionPointStatistics

	statistics.RecordsCount = uint32(len(listModels))

	if len(listModels) == 0 {
		return statistics
	}

	// order by ascending DetectionTimestamp
	getTimestamp := func(m *models.CurrentTrafficDataByDetectionPointModel) time.Time {
		return m.DetectionTimestamp
	}

	first, last := computeCommonTimestamps(listModels, getTimestamp)
	statistics.FirstRecordTimestamp = first
	statistics.LastRecordTimestamp = last

	// detect longest data gap
	gap, gapStart, gapEnd := computeLongestDataGap(listModels, getTimestamp)
	statistics.LongestDataGap = float32(gap)
	statistics.LongestDataGapStartTimestamp = gapStart
	statistics.LongestDataGapEndTimestamp = gapEnd

	// compute missing interval detection rate
	statistics.MissingIntervalDetectionRate = computeMissingIntervalDetectionRate(len(listModels), first, last, 300)

	// compute total
	statistics.TotalCountEquivalentVehicles = 0
	statistics.TotalHarmonicMeanSpeedAllRecords = 0
	for _, m := range listModels {
		statistics.TotalCountEquivalentVehicles += m.CountVehicleClassEquivalent
		statistics.TotalHarmonicMeanSpeedAllRecords += m.HarmonicMeanSpeedVehicleClassAll
	}
	statistics.TotalHarmonicMeanSpeedAllRecords /= float32(len(listModels))

	// compute rate
	statistics.PercentageRecordsWithEquivalentCounts = 0
	statistics.PercentageRecordsWithPositiveAverageSpeed = 0
	for _, m := range listModels {
		if m.CountVehicleClassEquivalent > 0 {
			statistics.PercentageRecordsWithEquivalentCounts++
		}
		if m.HarmonicMeanSpeedVehicleClassAll > 0 {
			statistics.PercentageRecordsWithPositiveAverageSpeed++
		}
	}
	statistics.PercentageRecordsWithEquivalentCounts = (float32(statistics.PercentageRecordsWithEquivalentCounts) / float32(len(listModels)) * 100)
	statistics.PercentageRecordsWithPositiveAverageSpeed = (float32(statistics.PercentageRecordsWithPositiveAverageSpeed) / float32(len(listModels)) * 100)

	return statistics
}

func computeCommonTimestamps[T any](list []T, getTime func(T) time.Time) (time.Time, time.Time) {
	if len(list) == 0 {
		return time.Time{}, time.Time{}
	}
	sort.Slice(list, func(i, j int) bool {
		return getTime(list[i]).Before(getTime(list[j]))
	})
	return getTime(list[0]), getTime(list[len(list)-1])
}

func computeLongestDataGap[T any](list []T, getTime func(T) time.Time) (gap uint32, start, end time.Time) {
	for i := 1; i < len(list); i++ {
		prev := getTime(list[i-1])
		curr := getTime(list[i])
		seconds := uint32(curr.Sub(prev).Seconds())
		if seconds > gap {
			gap = seconds
			start = prev
			end = curr
		}
	}
	return
}

func computeMissingIntervalDetectionRate(records int, first, last time.Time, intervalSec float64) float32 {
	expected := last.Sub(first).Seconds() / intervalSec
	missing := max(0, expected-float64(records))
	if expected == 0 {
		return 0
	}
	return float32(missing) * 100 / float32(expected)
}

func AggregateStatistics[T models.TimeStampedStatistic](stats []T) (totalCount uint32, first, last time.Time) {
	for _, stat := range stats {
		if stat.GetFirstRecordTimestamp().IsZero() || stat.GetLastRecordTimestamp().IsZero() {
			continue
		}
		if first.IsZero() || stat.GetFirstRecordTimestamp().Before(first) {
			first = stat.GetFirstRecordTimestamp()
		}
		if last.IsZero() || stat.GetLastRecordTimestamp().After(last) {
			last = stat.GetLastRecordTimestamp()
		}
		totalCount += stat.GetRecordsCount()
	}
	return
}
