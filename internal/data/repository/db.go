package repository

import (
	"context"
	"errors"
	"fmt"
	"spoke7-go/internal/data/models"
	"spoke7-go/internal/data/repository/postgres"
	"spoke7-go/pkg/logger"
	"time"
)

var (
	ErrNotFound            = errors.New("item not found")
	ErrUnsupportedDatabase = errors.New("unsupported database type! supported( postgres )")
	ErrInvalidObjectId     = errors.New("invalid object ID")
	ErrNotUnique           = errors.New("resource already exists")
	ErrCommandStillInUse   = errors.New("command is still in use by device profiles")
	ErrSlugEmpty           = errors.New("slug is nil or empty")
	ErrNameEmpty           = errors.New("name is required")
)

// DBClient defines the interface for interacting with the database.
// It provides methods to create, update, delete, retrieve, and list source data.
type DBClient interface {
	Close() error
	Migrate(ctx context.Context) error

	// CURRENT TRAFFIC DATA BY DETECTION POINT
	//CreateCurrentTrafficDataByDetectionPoint inserts new detection point traffic data into the database.
	CreateCurrentTrafficDataByDetectionPoint(ctx context.Context, currentTrafficDataByDetectionPoint *models.CurrentTrafficDataByDetectionPointModel) error
	// UpdateCurrentTrafficDataByDetectionPoint updates existing detection point traffic data in the database.
	UpdateCurrentTrafficDataByDetectionPoint(ctx context.Context, currentTrafficDataByDetectionPoint *models.CurrentTrafficDataByDetectionPointModel) error
	// DeleteCurrentTrafficDataByDetectionPoint removes detection point traffic data from the database by id.
	DeleteCurrentTrafficDataByDetectionPoint(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error
	// GetCurrentTrafficDataByDetectionPoint retrieves detection point traffic data from the database by id and optional params.
	GetCurrentTrafficDataByDetectionPoint(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.CurrentTrafficDataByDetectionPointModel, error)
	//ListCurrentTrafficDataByDetectionPoint lists all detection point traffic data in the database that match the optional params.
	ListCurrentTrafficDataByDetectionPoint(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.CurrentTrafficDataByDetectionPointModel, error)
	//ListAggregatedByDay
	ListCurrentTrafficDataByDetectionPointAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, from *time.Time, to *time.Time) ([]*models.CurrentTrafficDataByDetectionPointModel, error)
	//ListAggregatedByDetectionPoint
	ListCurrentTrafficDataByDetectionPointAggregatedByPoint(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.CurrentTrafficDataByDetectionPointStatistics, error)
	// BulkCreateCurrentTrafficDataByDetectionPoint inserts multiple detection point traffic data into the database.
	BulkCreateCurrentTrafficDataByDetectionPoint(ctx context.Context, models []*models.CurrentTrafficDataByDetectionPointModel) error

	// CURRENT TRAFFIC DATA BY DETECTION POINT BY LANES
	//CreateCurrentTrafficDataByDetectionPointByLane inserts new detection point traffic data into the database.
	CreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.CurrentTrafficDataByDetectionPointByLaneModel) error
	// UpdateCurrentTrafficDataByDetectionPointByLane updates existing detection point traffic data in the database.
	UpdateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.CurrentTrafficDataByDetectionPointByLaneModel) error
	// DeleteCurrentTrafficDataByDetectionPointByLane removes detection point traffic data from the database by id.
	DeleteCurrentTrafficDataByDetectionPointByLane(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	// GetCurrentTrafficDataByDetectionPointByLane retrieves detection point traffic data from the database by id and optional params.
	GetCurrentTrafficDataByDetectionPointByLane(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error)
	//ListCurrentTrafficDataByDetectionPointByLane lists all detection point traffic data in the database that match the optional params.
	ListCurrentTrafficDataByDetectionPointByLane(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByDay
	ListCurrentTrafficDataByDetectionPointByLaneAggregatedByDay(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp *time.Time, endTimestamp *time.Time) ([]*models.CurrentTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByLane
	ListCurrentTrafficDataByDetectionPointByLaneAggregatedByLane(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.CurrentTrafficDataByDetectionPointByLaneStatistics, error)
	// BulkCreateCurrentTrafficDataByDetectionPointByLane inserts multiple detection point traffic data into the database.
	BulkCreateCurrentTrafficDataByDetectionPointByLane(ctx context.Context, models []*models.CurrentTrafficDataByDetectionPointByLaneModel) error

	// // CURRENT TRAFFIC DATA BY DETECTION SECTION
	//CreateCurrentTrafficDataByDetectionSection inserts new detection point traffic data into the database.
	CreateCurrentTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.CurrentTrafficDataByDetectionSectionModel) error
	// UpdateCurrentTrafficDataByDetectionSection updates existing detection point traffic data in the database.
	UpdateCurrentTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.CurrentTrafficDataByDetectionSectionModel) error
	// DeleteCurrentTrafficDataByDetectionSection removes detection point traffic data from the database by id.
	DeleteCurrentTrafficDataByDetectionSection(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	// GetCurrentTrafficDataByDetectionSection retrieves detection point traffic data from the database by id and optional params.
	GetCurrentTrafficDataByDetectionSection(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.CurrentTrafficDataByDetectionSectionModel, error)
	//ListCurrentTrafficDataByDetectionSection lists all detection point traffic data in the database that match the optional params.
	ListCurrentTrafficDataByDetectionSection(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.CurrentTrafficDataByDetectionSectionModel, error)
	//ListAggregatedByDay
	ListCurrentTrafficDataByDetectionSectionAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, from *time.Time, to *time.Time) ([]*models.CurrentTrafficDataByDetectionSectionModel, error)
	//ListAggregatedByDetectionSection
	ListCurrentTrafficDataByDetectionSectionAggregatedBySection(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.CurrentTrafficDataByDetectionSectionStatistics, error)
	// BulkCreateCurrentTrafficDataByDetectionSection inserts multiple detection point traffic data into the database.
	BulkCreateCurrentTrafficDataByDetectionSection(ctx context.Context, models []*models.CurrentTrafficDataByDetectionSectionModel) error

	// REAL TIME TRAFFIC DATA BY DETECTION POINT BY LANES
	//CreateRealTimeTrafficDataByDetectionPointByLane inserts new detection point traffic data into the database.
	CreateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.RealTimeTrafficDataByDetectionPointByLaneModel) error
	// UpdateRealTimeTrafficDataByDetectionPointByLane updates existing detection point traffic data in the database.
	UpdateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.RealTimeTrafficDataByDetectionPointByLaneModel) error
	// DeleteRealTimeTrafficDataByDetectionPointByLane removes detection point traffic data from the database by id.
	DeleteRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	// GetRealTimeTrafficDataByDetectionPointByLane retrieves detection point traffic data from the database by id and optional params.
	GetRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error)
	//ListRealTimeTrafficDataByDetectionPointByLane lists all detection point traffic data in the database that match the optional params.
	ListRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByDay
	ListRealTimeTrafficDataByDetectionPointByLaneAggregatedByDay(ctx context.Context, dataSourceName string, laneIDs []string, from *time.Time, to *time.Time) ([]*models.RealTimeTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByLane
	ListRealTimeTrafficDataByDetectionPointByLaneAggregatedByLane(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsLane, error)
	// BulkCreateRealTimeTrafficDataByDetectionPointByLane inserts multiple detection point traffic data into the database.
	BulkCreateRealTimeTrafficDataByDetectionPointByLane(ctx context.Context, models []*models.RealTimeTrafficDataByDetectionPointByLaneModel) error

	// REAL TIME TRAFFIC DATA BY DETECTION SECTION
	//CreateRealTimeTrafficDataByDetectionSection inserts new detection point traffic data into the database.
	CreateRealTimeTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.RealTimeTrafficDataByDetectionSectionModel) error
	// UpdateRealTimeTrafficDataByDetectionSection updates existing detection point traffic data in the database.
	UpdateRealTimeTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.RealTimeTrafficDataByDetectionSectionModel) error
	// DeleteRealTimeTrafficDataByDetectionSection removes detection point traffic data from the database by id.
	DeleteRealTimeTrafficDataByDetectionSection(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	// GetRealTimeTrafficDataByDetectionSection retrieves detection point traffic data from the database by id and optional params.
	GetRealTimeTrafficDataByDetectionSection(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error)
	//ListRealTimeTrafficDataByDetectionSection lists all detection point traffic data in the database that match the optional params.
	ListRealTimeTrafficDataByDetectionSection(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error)
	//ListAggregatedByDay
	ListRealTimeTrafficDataByDetectionSectionAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, from *time.Time, to *time.Time) ([]*models.RealTimeTrafficDataByDetectionSectionModel, error)
	//ListAggregatedBySection
	ListRealTimeTrafficDataByDetectionSectionAggregatedBySection(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionSection, error)
	// BulkCreateRealTimeTrafficDataByDetectionSection inserts multiple detection point traffic data into the database.
	BulkCreateRealTimeTrafficDataByDetectionSection(ctx context.Context, models []*models.RealTimeTrafficDataByDetectionSectionModel) error

	// HISTORY DAY TRAFFIC DATA BY DETECTION POINT BY LANES
	//CreateHistoryDayTrafficDataByDetectionPointByLane inserts new detection point traffic data into the database.
	CreateHistoryDayTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.HistoryDayTrafficDataByDetectionPointByLaneModel) error
	// UpdateHistoryDayTrafficDataByDetectionPointByLane updates existing detection point traffic data in the database.
	UpdateHistoryDayTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.HistoryDayTrafficDataByDetectionPointByLaneModel) error
	// DeleteHistoryDayTrafficDataByDetectionPointByLane removes detection point traffic data from the database by id.
	DeleteHistoryDayTrafficDataByDetectionPointByLane(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	// GetHistoryDayTrafficDataByDetectionPointByLane retrieves detection point traffic data from the database by id and optional params.
	GetHistoryDayTrafficDataByDetectionPointByLane(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error)
	//ListHistoryDayTrafficDataByDetectionPointByLane lists all detection point traffic data in the database that match the optional params.
	ListHistoryDayTrafficDataByDetectionPointByLane(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByDay
	ListHistoryDayTrafficDataByDetectionPointByLaneAggregatedByDay(ctx context.Context, dataSourceName string, laneIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryDayTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByLane
	ListHistoryDayTrafficDataByDetectionPointByLaneAggregatedByLane(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsLane, error)
	// BulkCreateHistoryDayTrafficDataByDetectionPointByLane inserts multiple detection point traffic data into the database.
	BulkCreateHistoryDayTrafficDataByDetectionPointByLane(ctx context.Context, models []*models.HistoryDayTrafficDataByDetectionPointByLaneModel) error

	// HISTORY DAY TRAFFIC DATA BY DETECTION POINT
	//CreateHistoryDayTrafficDataByDetectionPoint inserts new detection point traffic data into the database.
	CreateHistoryDayTrafficDataByDetectionPoint(ctx context.Context, currentTrafficDataByDetectionPoint *models.HistoryDayTrafficDataByDetectionPointModel) error
	// UpdateHistoryDayTrafficDataByDetectionPoint updates existing detection point traffic data in the database.
	UpdateHistoryDayTrafficDataByDetectionPoint(ctx context.Context, currentTrafficDataByDetectionPoint *models.HistoryDayTrafficDataByDetectionPointModel) error
	// DeleteHistoryDayTrafficDataByDetectionPoint removes detection point traffic data from the database by id.
	DeleteHistoryDayTrafficDataByDetectionPoint(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error
	// GetHistoryDayTrafficDataByDetectionPoint retrieves detection point traffic data from the database by id and optional params.
	GetHistoryDayTrafficDataByDetectionPoint(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error)
	//ListHistoryDayTrafficDataByDetectionPoint lists all detection point traffic data in the database that match the optional params.
	ListHistoryDayTrafficDataByDetectionPoint(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error)
	//ListAggregatedByDay
	ListHistoryDayTrafficDataByDetectionPointAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryDayTrafficDataByDetectionPointModel, error)
	//ListAggregatedByPoint
	ListHistoryDayTrafficDataByDetectionPointAggregatedByPoint(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionPoint, error)
	// BulkCreateHistoryDayTrafficDataByDetectionPoint inserts multiple detection point traffic data into the database.
	BulkCreateHistoryDayTrafficDataByDetectionPoint(ctx context.Context, models []*models.HistoryDayTrafficDataByDetectionPointModel) error

	// HISTORY DAY TRAFFIC DATA BY DETECTION SECTION
	//CreateHistoryDayTrafficDataByDetectionSection inserts new detection point traffic data into the database.
	CreateHistoryDayTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	// UpdateHistoryDayTrafficDataByDetectionSection updates existing detection point traffic data in the database.
	UpdateHistoryDayTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	// DeleteHistoryDayTrafficDataByDetectionSection removes detection point traffic data from the database by id.
	DeleteHistoryDayTrafficDataByDetectionSection(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	// GetHistoryDayTrafficDataByDetectionSection retrieves detection point traffic data from the database by id and optional params.
	GetHistoryDayTrafficDataByDetectionSection(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	//ListHistoryDayTrafficDataByDetectionSection lists all detection point traffic data in the database that match the optional params.
	ListHistoryDayTrafficDataByDetectionSection(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	//ListAggregatedByDay
	ListHistoryDayTrafficDataByDetectionSectionAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	//ListAggregatedBySection
	ListHistoryDayTrafficDataByDetectionSectionAggregatedBySection(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionSection, error)
	// BulkCreateHistoryDayTrafficDataByDetectionSection inserts multiple detection point traffic data into the database.
	BulkCreateHistoryDayTrafficDataByDetectionSection(ctx context.Context, models []*models.HistoryTrafficDataByDetectionSectionModel) error

	// HISTORY HOUR TRAFFIC DATA BY DETECTION POINT BY LANES
	//CreateHistoryHourTrafficDataByDetectionPointByLane inserts new detection point traffic data into the database.
	CreateHistoryHourTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.HistoryHourTrafficDataByDetectionPointByLaneModel) error
	// UpdateHistoryHourTrafficDataByDetectionPointByLane updates existing detection point traffic data in the database.
	UpdateHistoryHourTrafficDataByDetectionPointByLane(ctx context.Context, currentTrafficDataByDetectionPointByLane *models.HistoryHourTrafficDataByDetectionPointByLaneModel) error
	// DeleteHistoryHourTrafficDataByDetectionPointByLane removes detection point traffic data from the database by id.
	DeleteHistoryHourTrafficDataByDetectionPointByLane(ctx context.Context, params models.DeleteTrafficDataByDetectionPointByLaneParams) error
	// GetHistoryHourTrafficDataByDetectionPointByLane retrieves detection point traffic data from the database by id and optional params.
	GetHistoryHourTrafficDataByDetectionPointByLane(ctx context.Context, params models.GetTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error)
	//ListHistoryHourTrafficDataByDetectionPointByLane lists all detection point traffic data in the database that match the optional params.
	ListHistoryHourTrafficDataByDetectionPointByLane(ctx context.Context, params models.ListTrafficDataByDetectionPointByLaneParams) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByDay
	ListHistoryHourTrafficDataByDetectionPointByLaneAggregatedByDay(ctx context.Context, dataSourceName string, laneIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointByLaneModel, error)
	//ListAggregatedByLane
	ListHistoryHourTrafficDataByDetectionPointByLaneAggregatedByLane(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsLane, error)
	// BulkCreateHistoryHourTrafficDataByDetectionPointByLane inserts multiple detection point traffic data into the database.
	BulkCreateHistoryHourTrafficDataByDetectionPointByLane(ctx context.Context, models []*models.HistoryHourTrafficDataByDetectionPointByLaneModel) error

	// HISTORY HOUR TRAFFIC DATA BY DETECTION POINT
	//CreateHistoryHourTrafficDataByDetectionPoint inserts new detection point traffic data into the database.
	CreateHistoryHourTrafficDataByDetectionPoint(ctx context.Context, currentTrafficDataByDetectionPoint *models.HistoryHourTrafficDataByDetectionPointModel) error
	// UpdateHistoryHourTrafficDataByDetectionPoint updates existing detection point traffic data in the database.
	UpdateHistoryHourTrafficDataByDetectionPoint(ctx context.Context, currentTrafficDataByDetectionPoint *models.HistoryHourTrafficDataByDetectionPointModel) error
	// DeleteHistoryHourTrafficDataByDetectionPoint removes detection point traffic data from the database by id.
	DeleteHistoryHourTrafficDataByDetectionPoint(ctx context.Context, params models.DeleteTrafficDataByDetectionPointParams) error
	// GetHistoryHourTrafficDataByDetectionPoint retrieves detection point traffic data from the database by id and optional params.
	GetHistoryHourTrafficDataByDetectionPoint(ctx context.Context, params models.GetTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error)
	//ListHistoryHourTrafficDataByDetectionPoint lists all detection point traffic data in the database that match the optional params.
	ListHistoryHourTrafficDataByDetectionPoint(ctx context.Context, params models.ListTrafficDataByDetectionPointParams) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error)
	//ListAggregatedByDay
	ListHistoryHourTrafficDataByDetectionPointAggregatedByDay(ctx context.Context, dataSourceName string, detectionPointIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryHourTrafficDataByDetectionPointModel, error)
	//ListAggregatedByPoint
	ListHistoryHourTrafficDataByDetectionPointAggregatedByPoint(ctx context.Context, dataSourceName string, detectionPointIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionPoint, error)
	// BulkCreateHistoryHourTrafficDataByDetectionPoint inserts multiple detection point traffic data into the database.
	BulkCreateHistoryHourTrafficDataByDetectionPoint(ctx context.Context, models []*models.HistoryHourTrafficDataByDetectionPointModel) error

	// HISTORY HOUR TRAFFIC DATA BY DETECTION SECTION
	//CreateHistoryHourTrafficDataByDetectionSection inserts new detection point traffic data into the database.
	CreateHistoryHourTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	// UpdateHistoryHourTrafficDataByDetectionSection updates existing detection point traffic data in the database.
	UpdateHistoryHourTrafficDataByDetectionSection(ctx context.Context, currentTrafficDataByDetectionSection *models.HistoryTrafficDataByDetectionSectionModel) error
	// DeleteHistoryHourTrafficDataByDetectionSection removes detection point traffic data from the database by id.
	DeleteHistoryHourTrafficDataByDetectionSection(ctx context.Context, params models.DeleteTrafficDataByDetectionSectionParams) error
	// GetHistoryHourTrafficDataByDetectionSection retrieves detection point traffic data from the database by id and optional params.
	GetHistoryHourTrafficDataByDetectionSection(ctx context.Context, params models.GetTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	//ListHistoryHourTrafficDataByDetectionSection lists all detection point traffic data in the database that match the optional params.
	ListHistoryHourTrafficDataByDetectionSection(ctx context.Context, params models.ListTrafficDataByDetectionSectionParams) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	//ListAggregatedByDay
	ListHistoryHourTrafficDataByDetectionSectionAggregatedByDay(ctx context.Context, dataSourceName string, detectionSectionIDs []string, from *time.Time, to *time.Time) ([]*models.HistoryTrafficDataByDetectionSectionModel, error)
	//ListAggregatedBySection
	ListHistoryHourTrafficDataByDetectionSectionAggregatedBySection(ctx context.Context, dataSourceName string, laneIDs []string, startTimestamp time.Time, endTimestamp time.Time) ([]models.TrafficStatisticsDetectionSection, error)
	// BulkCreateHistoryHourTrafficDataByDetectionSection inserts multiple detection point traffic data into the database.
	BulkCreateHistoryHourTrafficDataByDetectionSection(ctx context.Context, models []*models.HistoryTrafficDataByDetectionSectionModel) error
}

type DBConfig struct {
	Type         string `toml:"type" yaml:"type"`
	Host         string `toml:"host" yaml:"host"`
	Port         int    `toml:"port" yaml:"port"`
	Timeout      int    `toml:"timeout" yaml:"timeout"`
	DatabaseName string `toml:"databaseName" yaml:"databaseName"`
	Username     string `toml:"username" yaml:"username"`
	Password     string `toml:"password" yaml:"password"`
	BatchSize    int    `toml:"batchSize" yaml:"batchSize"`
	SSLMode      bool   `toml:"sslMode" yaml:"sslMode"`
	SSLRootCert  string `toml:"sslRootCert" yaml:"sslRootCert"`
}

func (c *DBConfig) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("Database.Type is required")
	}
	if c.Host == "" {
		return fmt.Errorf("Database.Host is required")
	}
	if c.Port == 0 {
		return fmt.Errorf("Database.Port is required")
	}
	if c.DatabaseName == "" {
		return fmt.Errorf("Database.DatabaseName is required")
	}
	if c.Username == "" {
		return fmt.Errorf("Database.Username is required")
	}
	if c.Password == "" {
		return fmt.Errorf("Database.Password is required")
	}

	if c.SSLMode {
		if c.SSLRootCert == "" {
			return fmt.Errorf("Database.SSLRootCert is required")
		}
	}
	return nil
}

// NewDBClient creates a new database client.
func NewClient(conf DBConfig, logger logger.Logger) (DBClient, error) {
	switch conf.Type {
	case "postgres":
		postgresConf := postgres.PostgresConf{
			Host:         conf.Host,
			Port:         uint16(conf.Port),
			DatabaseName: conf.DatabaseName,
			Username:     conf.Username,
			Password:     conf.Password,
			SSLMode:      conf.SSLMode,
			SSLRootCert:  conf.SSLRootCert,
			BatchSize:    conf.BatchSize,
		}
		return postgres.NewPostgresClient(postgresConf, logger)
	default:
		return nil, ErrUnsupportedDatabase
	}
}
