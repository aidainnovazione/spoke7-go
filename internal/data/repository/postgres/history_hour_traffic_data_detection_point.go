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

// CreateHistoryHourTrafficDataByDetectionPoint inserts new detection point traffic data into the database.
func (repo *postgresClient) CreateHistoryHourTrafficDataByDetectionPoint(ctx context.Context, historyTrafficDataByDetectionPoint *models.HistoryHourTrafficDataByDetectionPointModel) error {
	daoTraffic := dao.FromHistoryHourTrafficDataByDetectionPointModelToDao(*historyTrafficDataByDetectionPoint)

	if daoTraffic.DetectionInterval != 3600 {
		return fmt.Errorf("error inserting history traffic data: the given interval is not 3600 seconds")
	}
	if err := repo.db.WithContext(ctx).Create(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// UpdateHistoryHourTrafficDataByDetectionPoint updates existing detection point traffic data in the database.
func (repo *postgresClient) UpdateHistoryHourTrafficDataByDetectionPoint(ctx context.Context, historyTrafficDataByDetectionPoint *models.HistoryHourTrafficDataByDetectionPointModel) error {
	daoTraffic := dao.FromHistoryHourTrafficDataByDetectionPointModelToDao(*historyTrafficDataByDetectionPoint)

	if daoTraffic.DetectionInterval != 3600 {
		return fmt.Errorf("error updating history traffic data: the given interval is not 3600 seconds")
	}
	if err := repo.db.WithContext(ctx).Updates(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// DeleteHistoryHourTrafficDataByDetectionPoint removes detection point traffic data from the database by id.
func (repo *postgresClient) DeleteHistoryHourTrafficDataByDetectionPoint(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error {
	if err := repo.db.WithContext(ctx).Delete(&dao.HistoryHourTrafficDataByDetectionPointDao{}, "data_source_name = ? and  detection_point_id IN ? and detection_timestamp >= ? and detection_timestamp <= ?", params.DataSourceName, params.DetectionPointIDs, params.StartTimestamp, params.EndTimestamp).Error; err != nil {
		return err
	}
	return nil
}

// GetHistoryHourTrafficDataByDetectionPoint retrieves detection point traffic data from the database by id and optional params.
func (repo *postgresClient) GetHistoryHourTrafficDataByDetectionPoint(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error) {
	daoTraffic := []dao.HistoryHourTrafficDataByDetectionPointDao{}
	err := repo.db.Find(&daoTraffic, "data_source_name = ? and  detection_point_id IN ? and detection_timestamp = ?", params.DataSourceName, params.DetectionPointIDs, params.DetectionTimestamp).Error
	if err != nil {
		return nil, err
	}
	models := make([]*models.HistoryHourTrafficDataByDetectionPointModel, 0)
	for _, dao := range daoTraffic {
		model := dao.FromHistoryHourTrafficDataByDetectionPointDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

// ListHistoryHourTrafficDataByDetectionPoint lists all detection point traffic data in the database that match the optional params.
func (repo *postgresClient) ListHistoryHourTrafficDataByDetectionPoint(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error) {
	var daosTraffic []dao.HistoryHourTrafficDataByDetectionPointDao

	query := repo.db.WithContext(ctx).Where("data_source_name = ? and detection_point_id IN ?", params.DataSourceName, params.DetectionPointIDs)

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

	models := make([]*models.HistoryHourTrafficDataByDetectionPointModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromHistoryHourTrafficDataByDetectionPointDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) BulkCreateHistoryHourTrafficDataByDetectionPoint(ctx context.Context, models []*models.HistoryHourTrafficDataByDetectionPointModel) error {
	if len(models) == 0 {
		return errors.ErrWrongHistoryHourTrafficDetectionPointRequest
	}

	daosTraffic := make([]dao.HistoryHourTrafficDataByDetectionPointDao, 0, len(models))
	for _, model := range models {
		daosTraffic = append(daosTraffic, dao.FromHistoryHourTrafficDataByDetectionPointModelToDao(*model))
	}
	db := repo.db.WithContext(ctx).Session(&gorm.Session{CreateBatchSize: repo.conf.BatchSize})

	return db.Save(&daosTraffic).Error
}

func (repo *postgresClient) ListHistoryHourTrafficDataByDetectionPointAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error) {

	var daosTraffic []dao.HistoryHourTrafficDataByDetectionPointDao

	query := `
		SELECT 
			date_trunc('day', detection_timestamp) AS day,
			MIN(detection_timestamp) AS detection_timestamp,
			
			AVG(traffic_flow_vehicle_class1) AS traffic_flow_vehicle_class1,
			AVG(traffic_flow_vehicle_class2) AS traffic_flow_vehicle_class2,
			AVG(traffic_flow_vehicle_class3) AS traffic_flow_vehicle_class3,
			AVG(traffic_flow_vehicle_class4) AS traffic_flow_vehicle_class4,
			AVG(traffic_flow_vehicle_class5) AS traffic_flow_vehicle_class5,
			AVG(traffic_flow_vehicle_class6) AS traffic_flow_vehicle_class6,
			AVG(traffic_flow_vehicle_class7) AS traffic_flow_vehicle_class7,
			AVG(traffic_flow_vehicle_class8) AS traffic_flow_vehicle_class8,
			AVG(traffic_flow_vehicle_class_equivalent) AS traffic_flow_vehicle_class_equivalent,

			AVG(average_speed_vehicle_class1) AS average_speed_vehicle_class1,
			AVG(average_speed_vehicle_class2) AS average_speed_vehicle_class2,
			AVG(average_speed_vehicle_class3) AS average_speed_vehicle_class3,
			AVG(average_speed_vehicle_class4) AS average_speed_vehicle_class4,
			AVG(average_speed_vehicle_class5) AS average_speed_vehicle_class5,
			AVG(average_speed_vehicle_class6) AS average_speed_vehicle_class6,
			AVG(average_speed_vehicle_class7) AS average_speed_vehicle_class7,
			AVG(average_speed_vehicle_class8) AS average_speed_vehicle_class8,
			AVG(average_speed_vehicle_class_all) AS average_speed_vehicle_class_all,

			AVG(average_vehicle_length) AS average_vehicle_length,
			AVG(average_headway) AS average_headway,
			AVG(average_time_to_collision) AS average_time_to_collision

		FROM history_hour_traffic_data_by_detection_point
		WHERE data_source_name = ?		
		`
	args := []interface{}{dataSourceName} //, from, to}

	if from != nil && to != nil {
		query += "AND detection_timestamp >= ? AND detection_timestamp <= ?"
		args = append(args, *from, *to)
	}

	// Aggiungi filtro per i detectionPointID, se presenti
	if len(detectionPointIDs) > 0 {
		query += " AND detection_point_id IN ?"
		args = append(args, detectionPointIDs)
	}

	query += " GROUP BY day ORDER BY day ASC"

	err := repo.db.WithContext(ctx).Raw(query, args...).Scan(&daosTraffic).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.HistoryHourTrafficDataByDetectionPointModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromHistoryHourTrafficDataByDetectionPointDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) ListHistoryHourTrafficDataByDetectionPointAggregatedByPoint(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionPoint, error) {
	var results []models.TrafficStatisticsDetectionPoint

	query := `
	WITH ordered_data AS (
		SELECT
			detection_point_id,
			detection_timestamp,
			LAG(detection_timestamp) OVER (PARTITION BY detection_point_id ORDER BY detection_timestamp) AS prev_timestamp
		FROM history_hour_traffic_data_by_detection_point
		WHERE data_source_name = ?
		-- optional lane filter
		AND (array_length(?::text[], 1) IS NULL OR detection_point_id = ANY(?::text[]))
		 -- optional time filter with explicit casting
	    AND detection_timestamp >= ?::timestamp
	    AND detection_timestamp <= ?::timestamp
	),

	gaps AS (
			SELECT
				detection_point_id,
				EXTRACT(EPOCH FROM detection_timestamp - prev_timestamp) AS gap_seconds,
				prev_timestamp AS gap_start,
				detection_timestamp AS gap_end
			FROM ordered_data
			WHERE prev_timestamp IS NOT NULL
		),

	aggregated AS (
		SELECT
			ct.detection_point_id,
			COUNT(*) AS records_count,
			MIN(ct.detection_timestamp) AS first_record_timestamp,
			MAX(ct.detection_timestamp) AS last_record_timestamp
		FROM history_hour_traffic_data_by_detection_point ct
		WHERE data_source_name = ?
		AND ct.detection_timestamp >= ?::timestamp
		AND ct.detection_timestamp <= ?::timestamp
		AND (array_length(?::text[], 1) IS NULL OR detection_point_id = ANY(?::text[]))
		GROUP BY ct.detection_point_id
	),

	max_gaps AS (
		SELECT DISTINCT ON (detection_point_id)
			detection_point_id,
			gap_seconds AS longest_data_gap,
			gap_start AS longest_data_gap_start_timestamp,
			gap_end AS longest_data_gap_end_timestamp
		FROM gaps
		ORDER BY detection_point_id, gap_seconds DESC
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
		LEFT JOIN max_gaps mg ON a.detection_point_id = mg.detection_point_id
	)

	SELECT * FROM final;
`
	args := []interface{}{
		dataSourceName,
		pq.Array(detectionPointIDs),
		pq.Array(detectionPointIDs),
		startTimestamp,
		endTimestamp,
		dataSourceName,
		startTimestamp,
		endTimestamp,
		pq.Array(detectionPointIDs),
		pq.Array(detectionPointIDs),
	}

	err := repo.db.WithContext(ctx).Raw(query, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, err
}
