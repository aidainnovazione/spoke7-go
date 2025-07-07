package repository

import (
	"context"
	"errors"
	"fmt"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/repository/postgres"

	"spoke7-go/pkg/logger"
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

	// SOURCE DATA
	// CreateDataSource inserts new source data into the database.
	CreateDataSource(ctx context.Context, dataSource *models.DataSource) error
	// UpdateDataSource updates existing source data in the database.
	UpdateDataSource(ctx context.Context, dataSource *models.UpdateDataSource) error
	// DeleteDataSource removes source data from the database by name.
	DeleteDataSource(ctx context.Context, name string) error
	// GetDataSource retrieves source data from the database by name and optional params.
	GetDataSource(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error)
	// ListDataSource lists all source data in the database that match the optional params.
	ListDataSource(ctx context.Context, params models.DataSourceListParams, organizationName string) ([]*models.DataSource, error)

	// DETECTION SECTION
	// CreateDetection inserts new detection into the database.
	CreateDetectionSection(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error
	// CreateManyDetectionSection inserts multiple detection points into the database.
	CreateManyDetectionSection(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error
	// UpdateDetectionSection updates existing detection in the database.
	UpdateDetectionSection(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error
	// DeleteDetectionSection removes detection from the database by id.
	DeleteDetectionSection(ctx context.Context, datasourceName string, id string) error
	// GetDetectionSection retrieves detection from the database by id.
	GetDetectionSection(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error)
	// ListDetectionSection lists all detection in the database.
	ListDetectionSection(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error)

	// DETECTION POINT
	// CreateDetectionPoint inserts new detection point into the database.
	CreateDetectionPoint(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error
	// CreateManyDetectionPoint inserts multiple detection points into the database.
	CreateManyDetectionPoint(ctx context.Context, datasourceName string, detectionPoint []*models.DetectionPoint) error
	// UpdateDetectionPoint updates existing detection point in the database.
	UpdateDetectionPoint(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error
	// DeleteDetectionPoint removes detection point from the database by id.
	DeleteDetectionPoint(ctx context.Context, datasourceName string, id string) error
	// Delete All Detection Point by DatasourceName
	DeleteAllDetectionPointByDatasourceName(ctx context.Context, datasourceName string) error
	// GetDetectionPoint retrieves detection point from the database by id.
	GetDetectionPoint(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error)
	// ListDetectionPoint lists all detection point in the database.
	ListDetectionPoint(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error)

	// BulkCreateDetectionPoint inserts multiple detection points into the database.
	BulkCreateDetectionPoint(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error

	//
	CreateRoadNetwork(ctx context.Context, roadNetwork *models.RoadNetwork) (*models.RoadNetwork, error)
	GetRoadNetworkByID(ctx context.Context, id string) (*models.RoadNetwork, error)
	UpdateRoadNetwork(ctx context.Context, roadNetwork *models.RoadNetwork) (*models.RoadNetwork, error)
	ListRoadNetworks(ctx context.Context) ([]*models.RoadNetwork, error)
	DeleteRoadNetwork(ctx context.Context, id string) error

	//
	CreateDashboard(ctx context.Context, dashboard *models.Dashboard) error
	// UpdateDashboard updates existing source data in the database.
	UpdateDashboard(ctx context.Context, dashboard *models.Dashboard) error
	// DeleteDashboard removes source data from the database by name.
	DeleteDashboard(ctx context.Context, id string) error
	// GetDashboard retrieves source data from the database by name and optional params.
	GetDashboard(ctx context.Context, id string) (*models.Dashboard, error)
	// ListDashboard lists all source data in the database that match the optional params.
	ListDashboard(ctx context.Context, dataSourceName string, organizationName string) ([]*models.Dashboard, error)
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
