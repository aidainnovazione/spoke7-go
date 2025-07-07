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

// CreateHistoryDayTrafficDataByDetectionSection inserts new detection point traffic data into the database.
func (repo *postgresClient) CreateHistoryDayTrafficDataByDetectionSection(ctx context.Context, historyTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error {
	daoTraffic := dao.FromHistoryDayTrafficDataByDetectionSectionModelToDao(*historyTrafficDataByDetectionSection)

	if daoTraffic.DetectionInterval != 86400 {
		return fmt.Errorf("error inserting history traffic data: the given interval is not 86400 seconds")
	}

	if err := repo.db.WithContext(ctx).Create(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// UpdateHistoryDayTrafficDataByDetectionSection updates existing detection point traffic data in the database.
func (repo *postgresClient) UpdateHistoryDayTrafficDataByDetectionSection(ctx context.Context, historyTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error {
	daoTraffic := dao.FromHistoryDayTrafficDataByDetectionSectionModelToDao(*historyTrafficDataByDetectionSection)

	if daoTraffic.DetectionInterval != 86400 {
		return fmt.Errorf("error updating history traffic data: the given interval is not 86400 seconds")
	}

	if err := repo.db.WithContext(ctx).Updates(&daoTraffic).Error; err != nil {
		return err
	}

	return nil
}

// DeleteHistoryDayTrafficDataByDetectionSection removes detection point traffic data from the database by id.
func (repo *postgresClient) DeleteHistoryDayTrafficDataByDetectionSection(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error {
	if err := repo.db.WithContext(ctx).Delete(&dao.HistoryDayTrafficDataByDetectionSectionDao{}, "data_source_name = ? and  detection_section_id IN ? and detection_timestamp >= ? and detection_timestamp <= ?", params.DataSourceName, params.DetectionSectionIDs, params.StartTimestamp, params.EndTimestamp).Error; err != nil {
		return err
	}
	return nil
}

// GetHistoryDayTrafficDataByDetectionSection retrieves detection point traffic data from the database by id and optional params.
func (repo *postgresClient) GetHistoryDayTrafficDataByDetectionSection(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	daoTraffic := []dao.HistoryDayTrafficDataByDetectionSectionDao{}
	err := repo.db.Find(&daoTraffic, "data_source_name = ? and  detection_section_id IN ? and detection_timestamp = ?", params.DataSourceName, params.DetectionSectionIDs, params.DetectionTimestamp).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.HistoryTrafficDataByDetectionSectionModel, 0)
	for _, dao := range daoTraffic {
		model := dao.FromHistoryDayTrafficDataByDetectionSectionDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

// ListHistoryDayTrafficDataByDetectionSection lists all detection point traffic data in the database that match the optional params.
func (repo *postgresClient) ListHistoryDayTrafficDataByDetectionSection(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {
	var daosTraffic []dao.HistoryDayTrafficDataByDetectionSectionDao

	query := repo.db.WithContext(ctx).Where("data_source_name = ? and detection_section_id IN ?", params.DataSourceName, params.DetectionSectionIDs)

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

	models := make([]*models.HistoryTrafficDataByDetectionSectionModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromHistoryDayTrafficDataByDetectionSectionDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) BulkCreateHistoryDayTrafficDataByDetectionSection(ctx context.Context, models []*models.HistoryTrafficDataByDetectionSectionModel) error {
	if len(models) == 0 {
		return errors.ErrWrongHistoryDayTrafficDetectionSectionRequest
	}

	daosTraffic := make([]dao.HistoryDayTrafficDataByDetectionSectionDao, 0, len(models))
	for _, model := range models {
		if model.DetectionInterval != 86400 {
			return fmt.Errorf("error inserting history traffic data: the given interval is not 300 seconds")
		}
		daosTraffic = append(daosTraffic, dao.FromHistoryDayTrafficDataByDetectionSectionModelToDao(*model))
	}
	db := repo.db.WithContext(ctx).Session(&gorm.Session{CreateBatchSize: repo.conf.BatchSize})

	return db.Save(&daosTraffic).Error
}

func (repo *postgresClient) ListHistoryDayTrafficDataByDetectionSectionAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error) {

	var daosTraffic []dao.HistoryDayTrafficDataByDetectionSectionDao

	query := `
		SELECT 
			date_trunc('day', detection_timestamp) AS day,
			MIN(detection_timestamp) AS detection_timestamp,
	
			AVG(forward_speed_count) AS forward_speed_count,
			AVG(forward_speed_count_max_flow) AS forward_speed_count_max_flow,
			AVG(forward_speed_count_min_flow) AS forward_speed_count_min_flow,
			AVG(backward_speed_count) AS backward_speed_count,
			AVG(backward_speed_count_max_flow) AS backward_speed_count_max_flow,
			AVG(backward_speed_count_min_flow) AS backward_speed_count_min_flow

		FROM history_day_traffic_data_by_detection_section
		WHERE data_source_name = ?		
		`
	args := []interface{}{dataSourceName} //, from, to}

	if from != nil && to != nil {
		query += "AND detection_timestamp >= ? AND detection_timestamp <= ?"
		args = append(args, *from, *to)
	}

	// Aggiungi filtro per i detectionSectionID, se presenti
	if len(detectionSectionIDs) > 0 {
		query += " AND detection_section_id IN ?"
		args = append(args, detectionSectionIDs)
	}

	query += " GROUP BY day ORDER BY day ASC"

	err := repo.db.WithContext(ctx).Raw(query, args...).Scan(&daosTraffic).Error
	if err != nil {
		return nil, err
	}

	models := make([]*models.HistoryTrafficDataByDetectionSectionModel, 0)
	for _, dao := range daosTraffic {
		model := dao.FromHistoryDayTrafficDataByDetectionSectionDaoToModel()
		models = append(models, &model)
	}

	return models, nil
}

func (repo *postgresClient) ListHistoryDayTrafficDataByDetectionSectionAggregatedBySection(ctx context.Context, dataSourceName string, detectionSectionIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionSection, error) {
	var results []models.TrafficStatisticsDetectionSection

	query := `
	WITH ordered_data AS (
		SELECT
			detection_section_id,
			detection_timestamp,
			LAG(detection_timestamp) OVER (PARTITION BY detection_section_id ORDER BY detection_timestamp) AS prev_timestamp
		FROM history_day_traffic_data_by_detection_section
		WHERE data_source_name = ?
		-- optional lane filter
		AND (array_length(?::text[], 1) IS NULL OR detection_section_id = ANY(?::text[]))
		 -- optional time filter with explicit casting
	    AND detection_timestamp >= ?::timestamp
	    AND detection_timestamp <= ?::timestamp
	),

	gaps AS (
			SELECT
				detection_section_id,
				EXTRACT(EPOCH FROM detection_timestamp - prev_timestamp) AS gap_seconds,
				prev_timestamp AS gap_start,
				detection_timestamp AS gap_end
			FROM ordered_data
			WHERE prev_timestamp IS NOT NULL
		),

	aggregated AS (
		SELECT
			ct.detection_section_id,
			COUNT(*) AS records_count,
			MIN(ct.detection_timestamp) AS first_record_timestamp,
			MAX(ct.detection_timestamp) AS last_record_timestamp
		FROM history_day_traffic_data_by_detection_section ct
		WHERE data_source_name = ?
		AND ct.detection_timestamp >= ?::timestamp
		AND ct.detection_timestamp <= ?::timestamp
		AND (array_length(?::text[], 1) IS NULL OR detection_section_id = ANY(?::text[]))
		GROUP BY ct.detection_section_id
	),

	max_gaps AS (
		SELECT DISTINCT ON (detection_section_id)
			detection_section_id,
			gap_seconds AS longest_data_gap,
			gap_start AS longest_data_gap_start_timestamp,
			gap_end AS longest_data_gap_end_timestamp
		FROM gaps
		ORDER BY detection_section_id, gap_seconds DESC
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
		LEFT JOIN max_gaps mg ON a.detection_section_id = mg.detection_section_id
	)

	SELECT * FROM final;
`
	args := []interface{}{
		dataSourceName,
		pq.Array(detectionSectionIDs),
		pq.Array(detectionSectionIDs),
		startTimestamp,
		endTimestamp,
		dataSourceName,
		startTimestamp,
		endTimestamp,
		pq.Array(detectionSectionIDs),
		pq.Array(detectionSectionIDs),
	}

	err := repo.db.WithContext(ctx).Raw(query, args...).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, err
}
