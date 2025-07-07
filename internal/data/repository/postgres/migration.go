package postgres

import (
	"context"
	"fmt"
	"spoke7-go/internal/data/repository/postgres/dao"
)

func (pc *postgresClient) Migrate(ctx context.Context) error {
	err := pc.db.WithContext(ctx).AutoMigrate(
		&dao.CurrentTrafficDataByDetectionPointDao{},
		&dao.CurrentTrafficDataByDetectionPointByLaneDao{},
		&dao.CurrentTrafficDataByDetectionSectionDao{},
		&dao.RealTimeTrafficDataByDetectionPointByLaneDao{},
		&dao.RealTimeTrafficDataByDetectionSectionDao{},
		&dao.HistoryDayTrafficDataByDetectionPointDao{},
		&dao.HistoryDayTrafficDataByDetectionPointByLaneDao{},
		&dao.HistoryDayTrafficDataByDetectionSectionDao{},
		&dao.HistoryHourTrafficDataByDetectionPointDao{},
		&dao.HistoryHourTrafficDataByDetectionPointByLaneDao{},
		&dao.HistoryHourTrafficDataByDetectionSectionDao{},
	)
	if err != nil {
		return fmt.Errorf("error automigrating tables: %v", err)
	}

	// Create the hypertables

	// real time
	queryCreateHypertableByRealTimeLane := `SELECT create_hypertable('real_time_traffic_data_by_detection_point_by_lane', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByRealTimeLane).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable real_time_traffic_data_by_detection_point_by_lane: %v", err)
	}

	queryCreateHypertableRealTimeDetectionSection := `SELECT create_hypertable('real_time_traffic_data_by_detection_section', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableRealTimeDetectionSection).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable real_time_traffic_data_by_detection_section: %v", err)
	}

	// current
	queryCreateHypertableByCurrentLane := `SELECT create_hypertable('current_traffic_data_by_detection_point_by_lane', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByCurrentLane).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable current_traffic_data_by_detection_point_by_lane: %v", err)
	}

	queryCreateHypertableDetectionPoint := `SELECT create_hypertable('current_traffic_data_by_detection_point', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableDetectionPoint).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable current_traffic_data_by_detection_point: %v", err)
	}

	queryCreateHypertableDetectionSection := `SELECT create_hypertable('current_traffic_data_by_detection_section', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableDetectionSection).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable current_traffic_data_by_detection_section: %v", err)
	}

	// history by day
	queryCreateHypertableByHistoryDayByLane := `SELECT create_hypertable('history_day_traffic_data_by_detection_point_by_lane', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByHistoryDayByLane).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable history_day_traffic_data_by_detection_point_by_lane: %v", err)
	}

	queryCreateHypertableByHistoryDayByDetectionPoint := `SELECT create_hypertable('history_day_traffic_data_by_detection_point', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByHistoryDayByDetectionPoint).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable history_day_traffic_data_by_detection_point: %v", err)
	}

	queryCreateHypertableByHistoryDayByDetectionSection := `SELECT create_hypertable('history_day_traffic_data_by_detection_section', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByHistoryDayByDetectionSection).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable history_day_traffic_data_by_detection_section: %v", err)
	}

	// history by hour
	queryCreateHypertableByHistoryHourByLane := `SELECT create_hypertable('history_hour_traffic_data_by_detection_point_by_lane', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByHistoryHourByLane).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable history_hour_traffic_data_by_detection_point_by_lane: %v", err)
	}

	queryCreateHypertableByHistoryHourByDetectionPoint := `SELECT create_hypertable('history_hour_traffic_data_by_detection_point', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByHistoryHourByDetectionPoint).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable history_hour_traffic_data_by_detection_point: %v", err)
	}

	queryCreateHypertableByHistoryHourByDetectionSection := `SELECT create_hypertable('history_hour_traffic_data_by_detection_section', 'detection_timestamp', if_not_exists => TRUE);`
	err = pc.db.Exec(queryCreateHypertableByHistoryHourByDetectionSection).Error
	if err != nil {
		return fmt.Errorf("unable to create hypertable history_hour_traffic_data_by_detection_section: %v", err)
	}

	return nil
}
