package postgres

import (
	"context"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository/postgres/dao"
	"spoke7-go/internal/errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// CreateRealTimeTrafficDataByDetectionPointByLane inserts new detection point traffic data into the database.
func (repo *postgresClient) CreateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, realTimeTrafficDataByDetectionPointByLane *models.RealTimeTrafficDataByDetectionPointByLaneModel) error {
	daoTraffic := dao.FromRealTimeTrafficDataByDetectionPointByLaneModelToDao(*realTimeTrafficDataByDetectionPointByLane)

	if err := repo.db.WithContext(ctx).Create(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// UpdateRealTimeTrafficDataByDetectionPointByLane updates existing detection point traffic data in the database.
func (repo *postgresClient) UpdateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, realTimeTrafficDataByDetectionPointByLane *models.RealTimeTrafficDataByDetectionPointByLaneModel) error {
	daoTraffic := dao.FromRealTimeTrafficDataByDetectionPointByLaneModelToDao(*realTimeTrafficDataByDetectionPointByLane)

	if err := repo.db.WithContext(ctx).Updates(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// DeleteRealTimeTrafficDataByDetectionPointByLane removes detection point traffic data from the database by id.
func (repo *postgresClient) DeleteRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error {
	if err := repo.db.WithContext(ctx).Delete(&dao.RealTimeTrafficDataByDetectionPointByLaneDao{}, "data_source_name = ? and lane_id IN ? and detection_timestamp >= ? and detection_timestamp <= ?", params.DataSourceName, params.LaneIDs, params.StartTimestamp, params.EndTimestamp).Error; err != nil {
		return err
	}
	return nil
}

// ListRealTimeTrafficDataByDetectionPointByLane lists all detection point traffic data in the database that match the optional params.
func (repo *postgresClient) ListRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	var daosTraffic []dao.RealTimeTrafficDataByDetectionPointByLaneDao

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

	models := make([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromRealTimeTrafficDataByDetectionPointByLaneDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) BulkCreateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, models []*models.RealTimeTrafficDataByDetectionPointByLaneModel) error {
	if len(models) == 0 {
		return errors.ErrWrongRealTimeTrafficDetectionPointByLaneRequest
	}

	daosTraffic := make([]dao.RealTimeTrafficDataByDetectionPointByLaneDao, 0, len(models))
	for _, model := range models {
		daosTraffic = append(daosTraffic, dao.FromRealTimeTrafficDataByDetectionPointByLaneModelToDao(*model))
	}
	db := repo.db.WithContext(ctx).Session(&gorm.Session{CreateBatchSize: repo.conf.BatchSize})

	return db.Save(&daosTraffic).Error
}

// GetRealTimeTrafficDataByDetectionPointByLane retrieves detection point traffic data from the database by id and optional params.
func (repo *postgresClient) GetRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {
	daosTraffic := []dao.RealTimeTrafficDataByDetectionPointByLaneDao{}
	err := repo.db.Find(&daosTraffic, "data_source_name = ? and detection_timestamp = ? and lane_id IN ?", params.DataSourceName, params.DetectionTimestamp, params.LaneIDs).Error
	if err != nil {
		return nil, err
	}
	models := make([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromRealTimeTrafficDataByDetectionPointByLaneDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) ListRealTimeTrafficDataByDetectionPointByLaneAggregatedByDay(ctx context.Context, dataSourceName string, laneIDs []string, from *time.Time, to *time.Time) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error) {

	var daosTraffic []dao.RealTimeTrafficDataByDetectionPointByLaneDao

	query := `
		SELECT 
			date_trunc('day', detection_timestamp) AS day,
			MIN(detection_timestamp) AS detection_timestamp,

			AVG(vehicle_speed) AS vehicle_speed,
			AVG(vehicle_length) AS vehicle_length,
			AVG(vehicle_headway) AS vehicle_headway,
	
			SUM(CASE WHEN queue_present THEN 1 ELSE 0 END) AS queue_present,
			SUM(CASE WHEN correct_flow_direction THEN 1 ELSE 0 END) AS correct_flow_direction,
				
			MODE() WITHIN GROUP (ORDER BY detection_type) AS detection_type,
			MODE() WITHIN GROUP (ORDER BY detection_technology) AS detection_technology,
			MODE() WITHIN GROUP (ORDER BY vehicle_class) AS vehicle_class

		FROM real_time_traffic_data_by_detection_point_by_lane
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

	models := make([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromRealTimeTrafficDataByDetectionPointByLaneDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) ListRealTimeTrafficDataByDetectionPointByLaneAggregatedByLane(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsLane, error) {
	var results []models.TrafficStatisticsLane

	query := `
	WITH ordered_data AS (
		SELECT
			lane_id,
			detection_timestamp,
			LAG(detection_timestamp) OVER (PARTITION BY lane_id ORDER BY detection_timestamp) AS prev_timestamp
		FROM real_time_traffic_data_by_detection_point_by_lane
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
			MAX(ct.detection_timestamp) AS last_record_timestamp
		FROM real_time_traffic_data_by_detection_point_by_lane ct
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
