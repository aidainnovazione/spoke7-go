package postgres

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository/postgres/dao"
	"spoke7-go/internal/errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// CreateCurrentTrafficDataByDetectionPointByLane inserts new detection point traffic data into the database.
func (repo *postgresClient) CreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.CurrentTrafficDataByDetectionPointByLaneModel) error {
	daoTraffic := dao.FromCurrentTrafficDataByDetectionPointByLaneModelToDao(*currentTrafficDataByDetectionPointByLane)

	if daoTraffic.DetectionInterval != 300 {
		return fmt.Errorf("error inserting current traffic data: the given interval is not 300 seconds")
	}
	if err := repo.db.WithContext(ctx).Create(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// UpdateCurrentTrafficDataByDetectionPointByLane updates existing detection point traffic data in the database.
func (repo *postgresClient) UpdateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.CurrentTrafficDataByDetectionPointByLaneModel) error {
	daoTraffic := dao.FromCurrentTrafficDataByDetectionPointByLaneModelToDao(*currentTrafficDataByDetectionPointByLane)

	if daoTraffic.DetectionInterval != 300 {
		return fmt.Errorf("error updating current traffic data: the given interval is not 300 seconds")
	}
	if err := repo.db.WithContext(ctx).Updates(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCurrentTrafficDataByDetectionPointByLane removes detection point traffic data from the database by id.
func (repo *postgresClient) DeleteCurrentTrafficDataByDetectionPointByLane(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error {
	if err := repo.db.WithContext(ctx).Delete(&dao.CurrentTrafficDataByDetectionPointByLaneDao{}, "data_source_name = ? and lane_id IN ? and detection_timestamp >= ? and detection_timestamp <= ?", params.DataSourceName, params.LaneIDs, params.StartTimestamp, params.EndTimestamp).Error; err != nil {
		return err
	}
	return nil
}

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ListCurrentTrafficDataByDetectionPointByLane lists all detection point traffic data in the database that match the optional params.
func (repo *postgresClient) ListCurrentTrafficDataByDetectionPointByLane(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	var daosTraffic []dao.CurrentTrafficDataByDetectionPointByLaneDao

	query := repo.db.WithContext(ctx).Where("data_source_name = ? and lane_id IN ?", params.DataSourceName, params.LaneIDs)

	if !params.StartTime.IsZero() {
		query = query.Where("detection_timestamp >= ?", params.StartTime)
	}

	if !params.EndTime.IsZero() {
		query = query.Where("detection_timestamp <= ?", params.EndTime)
	}

	err := query.Find(&daosTraffic).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.CurrentTrafficDataByDetectionPointByLaneModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromCurrentTrafficDataByDetectionPointByLaneDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, models []*models.CurrentTrafficDataByDetectionPointByLaneModel) error {
	if len(models) == 0 {
		return errors.ErrWrongCurrentTrafficDetectionPointByLaneRequest
	}

	daosTraffic := make([]dao.CurrentTrafficDataByDetectionPointByLaneDao, 0, len(models))
	for _, model := range models {
		if model.DetectionInterval != 300 {
			return fmt.Errorf("error updating current traffic data: the given interval is not 300 seconds")
		}
		daosTraffic = append(daosTraffic, dao.FromCurrentTrafficDataByDetectionPointByLaneModelToDao(*model))
	}
	db := repo.db.WithContext(ctx).Session(&gorm.Session{CreateBatchSize: repo.conf.BatchSize})

	return db.Save(&daosTraffic).Error
}

func (repo *postgresClient) ListCurrentTrafficDataByDetectionPointByLaneAggregatedByDay(ctx context.Context, dataSourceName string, laneIDs []string, from *time.Time, to *time.Time) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {

	var daosTraffic []dao.CurrentTrafficDataByDetectionPointByLaneDao

	query := `
		SELECT 
			date_trunc('day', detection_timestamp) AS day,
			MIN(detection_timestamp) AS detection_timestamp,
			SUM(count_vehicle_class1) AS count_vehicle_class1,
			SUM(count_vehicle_class2) AS count_vehicle_class2,
			SUM(count_vehicle_class3) AS count_vehicle_class3,
			SUM(count_vehicle_class4) AS count_vehicle_class4,
			SUM(count_vehicle_class5) AS count_vehicle_class5,
			SUM(count_vehicle_class6) AS count_vehicle_class6,
			SUM(count_vehicle_class7) AS count_vehicle_class7,
			SUM(count_vehicle_class8) AS count_vehicle_class8,
			SUM(count_vehicle_class_all) AS count_vehicle_class_all,
	
			AVG(harmonic_mean_speed_vehicle_class1) AS harmonic_mean_speed_vehicle_class1,
			AVG(harmonic_mean_speed_vehicle_class2) AS harmonic_mean_speed_vehicle_class2,
			AVG(harmonic_mean_speed_vehicle_class3) AS harmonic_mean_speed_vehicle_class3,
			AVG(harmonic_mean_speed_vehicle_class4) AS harmonic_mean_speed_vehicle_class4,
			AVG(harmonic_mean_speed_vehicle_class5) AS harmonic_mean_speed_vehicle_class5,
			AVG(harmonic_mean_speed_vehicle_class6) AS harmonic_mean_speed_vehicle_class6,
			AVG(harmonic_mean_speed_vehicle_class7) AS harmonic_mean_speed_vehicle_class7,
			AVG(harmonic_mean_speed_vehicle_class8) AS harmonic_mean_speed_vehicle_class8,
			AVG(harmonic_mean_speed_vehicle_class_all) AS harmonic_mean_speed_vehicle_class_all,
	
			SUM(count_detected_speed_vehicle_under50) AS count_detected_speed_vehicle_under50,
			SUM(count_detected_speed_vehicle_between50_100) AS count_detected_speed_vehicle_between50_100,
			SUM(count_detected_speed_vehicle_over100) AS count_detected_speed_vehicle_over100,
	
			AVG(average_vehicle_length) AS average_vehicle_length,
			AVG(average_headway) AS average_headway,
			AVG(std_headway) AS std_headway,
			AVG(average_time_to_collision) AS average_time_to_collision,
			AVG(std_time_to_collision) AS std_time_to_collision
		FROM current_traffic_data_by_detection_point_by_lane
		WHERE data_source_name = ?		
		`
	args := []interface{}{dataSourceName} //, from, to}

	if from != nil && to != nil {
		query += "AND detection_timestamp >= ? AND detection_timestamp <= ?"
		args = append(args, *from, *to)
	}

	// Aggiungi filtro per i laneID, se presenti
	if len(laneIDs) > 0 {
		query += " AND lane_id IN ?"
		args = append(args, laneIDs)
	}

	query += " GROUP BY day ORDER BY day ASC"

	err := repo.db.WithContext(ctx).Raw(query, args...).Scan(&daosTraffic).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.CurrentTrafficDataByDetectionPointByLaneModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromCurrentTrafficDataByDetectionPointByLaneDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

// GetCurrentTrafficDataByDetectionPointByLane retrieves detection point traffic data from the database by id and optional params.
func (repo *postgresClient) GetCurrentTrafficDataByDetectionPointByLane(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error) {
	daosTraffic := []dao.CurrentTrafficDataByDetectionPointByLaneDao{}
	err := repo.db.Find(&daosTraffic, "data_source_name = ? and detection_timestamp = ? and lane_id IN ?", params.DataSourceName, params.DetectionTimestamp, params.LaneIDs).Error
	if err != nil {
		return nil, err
	}
	models := make([]*models.CurrentTrafficDataByDetectionPointByLaneModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromCurrentTrafficDataByDetectionPointByLaneDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) ListCurrentTrafficDataByDetectionPointByLaneAggregatedByLane(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.CurrentTrafficDataByDetectionPointByLaneStatistics, error) {
	var results []models.CurrentTrafficDataByDetectionPointByLaneStatistics

	query := `
	WITH ordered_data AS (
		SELECT
			lane_id,
			detection_timestamp,
			LAG(detection_timestamp) OVER (PARTITION BY lane_id ORDER BY detection_timestamp) AS prev_timestamp
		FROM current_traffic_data_by_detection_point_by_lane
		WHERE data_source_name = ?
		-- optional lane filter
		AND (array_length(?::text[], 1) IS NULL OR lane_id = ANY(?::text[]))
		 -- optional time filter with explicit casting
	    AND detection_timestamp >= ?::timestamp
	    AND detection_timestamp <= ?::timestamp
	),

	gaps AS (
			SELECT
				lane_id,
				EXTRACT(EPOCH FROM detection_timestamp - prev_timestamp) AS gap_seconds,
				prev_timestamp AS gap_start,
				detection_timestamp AS gap_end
			FROM ordered_data
			WHERE prev_timestamp IS NOT NULL
		),

	aggregated AS (
		SELECT
			ct.lane_id,
			COUNT(*) AS records_count,
			MIN(ct.detection_timestamp) AS first_record_timestamp,
			MAX(ct.detection_timestamp) AS last_record_timestamp,
			SUM(ct.count_vehicle_class_all) AS total_count_all_vehicles,
			AVG(ct.harmonic_mean_speed_vehicle_class_all) AS total_harmonic_mean_speed_all_records,
			SUM(CASE WHEN ct.count_vehicle_class_all > 0 THEN 1 ELSE 0 END) * 100.0 / COUNT(*) AS percentage_records_with_counts,
			SUM(CASE WHEN ct.harmonic_mean_speed_vehicle_class_all > 0 THEN 1 ELSE 0 END) * 100.0 / COUNT(*) AS percentage_records_with_positive_average_speed
		FROM current_traffic_data_by_detection_point_by_lane ct
		WHERE data_source_name = ?
		AND ct.detection_timestamp >= ?::timestamp
		AND ct.detection_timestamp <= ?::timestamp
		AND (array_length(?::text[], 1) IS NULL OR lane_id = ANY(?::text[]))
		GROUP BY ct.lane_id
	),

	max_gaps AS (
		SELECT DISTINCT ON (lane_id)
			lane_id,
			gap_seconds AS longest_data_gap,
			gap_start AS longest_data_gap_start_timestamp,
			gap_end AS longest_data_gap_end_timestamp
		FROM gaps
		ORDER BY lane_id, gap_seconds DESC
	),

	final AS (
		SELECT
			a.*,
			COALESCE(mg.longest_data_gap, 0) AS longest_data_gap,
			mg.longest_data_gap_start_timestamp,
			mg.longest_data_gap_end_timestamp,
			-- Missing Interval Detection Rate (5-minute intervals)
			CASE
				WHEN a.records_count <= 1 THEN 0
				ELSE
					ROUND(
						((
							1 -
							(a.records_count::FLOAT / GREATEST(
								EXTRACT(EPOCH FROM (a.last_record_timestamp - a.first_record_timestamp)), 1)
							)
						) * 100)::numeric, 2
					)
			END AS missing_interval_detection_rate
		FROM aggregated a
		LEFT JOIN max_gaps mg ON a.lane_id = mg.lane_id
	)

	SELECT * FROM final;
`
	args := []interface{}{
		dataSourceName,
		pq.Array(laneIDs),
		pq.Array(laneIDs),
		startTimestamp,
		endTimestamp,
		dataSourceName,
		startTimestamp,
		endTimestamp,
		pq.Array(laneIDs),
		pq.Array(laneIDs),
	}

	err := repo.db.WithContext(ctx).Raw(query, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, err
}
