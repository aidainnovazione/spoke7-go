package mock

import (
	"context"
	"spoke7-go/internal/metadata/models"
)

type MockDBClient struct {
	CloseFunc    func() error
	MigrateFunc  func(ctx context.Context) error
	migrateCalls []struct {
		ctx context.Context
	}

	// Source Data
	CreateDataSourceFunc  func(ctx context.Context, dataSource *models.DataSource) error
	createDataSourceCalls []struct {
		ctx        context.Context
		dataSource *models.DataSource
	}

	UpdateDataSourceFunc  func(ctx context.Context, dataSource *models.UpdateDataSource) error
	updateDataSourceCalls []struct {
		ctx        context.Context
		dataSource *models.UpdateDataSource
	}

	DeleteDataSourceFunc  func(ctx context.Context, name string) error
	deleteDataSourceCalls []struct {
		ctx  context.Context
		name string
	}

	GetDataSourceFunc  func(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error)
	getDataSourceCalls []struct {
		ctx    context.Context
		name   string
		params models.DataSourceGetParams
	}

	ListDataSourceFunc  func(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error)
	listDataSourceCalls []struct {
		ctx    context.Context
		params models.DataSourceListParams
	}

	// Detection Section
	CreateDetectionSectionFunc  func(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error
	createDetectionSectionCalls []struct {
		ctx              context.Context
		detectionSection *models.DetectionSection
	}

	CreateManyDetectionSectionFunc  func(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error
	createManyDetectionSectionCalls []struct {
		ctx              context.Context
		detectionSection []*models.DetectionSection
	}

	UpdateDetectionSectionFunc  func(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error
	updateDetectionSectionCalls []struct {
		ctx              context.Context
		detectionSection *models.DetectionSection
	}

	DeleteDetectionSectionFunc  func(ctx context.Context, datasourceName string, id string) error
	deleteDetectionSectionCalls []struct {
		ctx context.Context
		id  string
	}

	GetDetectionSectionFunc  func(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error)
	getDetectionSectionCalls []struct {
		ctx context.Context
		id  string
	}

	ListDetectionSectionFunc  func(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error)
	listDetectionSectionCalls []struct {
		ctx context.Context
	}

	// Detection Point
	CreateDetectionPointFunc  func(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error
	createDetectionPointCalls []struct {
		ctx            context.Context
		detectionPoint *models.DetectionPoint
	}

	CreateManyDetectionPointFunc  func(ctx context.Context, datasourceName string, detectionPoint []*models.DetectionPoint) error
	createManyDetectionPointCalls []struct {
		ctx            context.Context
		detectionPoint []*models.DetectionPoint
	}

	UpdateDetectionPointFunc  func(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error
	updateDetectionPointCalls []struct {
		ctx            context.Context
		detectionPoint *models.DetectionPoint
	}

	DeleteDetectionPointFunc  func(ctx context.Context, datasourceName string, id string) error
	deleteDetectionPointCalls []struct {
		ctx context.Context
		id  string
	}

	GetDetectionPointFunc  func(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error)
	getDetectionPointCalls []struct {
		ctx context.Context
		id  string
	}

	ListDetectionPointFunc  func(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error)
	listDetectionPointCalls []struct {
		ctx context.Context
	}

	BulkCreateDetectionPointFunc  func(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error
	bulkCreateDetectionPointCalls []struct {
		ctx             context.Context
		detectionPoints []*models.DetectionPoint
	}

	// Dashboard

	CreateDashboardFunc  func(ctx context.Context, dashboard *models.Dashboard) error
	createDashboardCalls []struct {
		ctx       context.Context
		dashboard *models.Dashboard
	}

	UpdateDashboardFunc  func(ctx context.Context, dashboard *models.Dashboard) error
	updateDashboardCalls []struct {
		ctx       context.Context
		dashboard *models.Dashboard
	}

	DeleteDashboardFunc  func(ctx context.Context, name string) error
	deleteDashboardCalls []struct {
		ctx  context.Context
		name string
	}

	GetDashboardFunc  func(ctx context.Context, id string) (*models.Dashboard, error)
	getDashboardCalls []struct {
		ctx context.Context
		id  string
	}

	ListDashboardFunc  func(ctx context.Context, dataSourceName string) ([]*models.Dashboard, error)
	listDashboardCalls []struct {
		ctx            context.Context
		dataSourceName string
	}
}

// NewMockDBClient creates a new mock DBClient instance
func NewMockDBClient() *MockDBClient {
	return &MockDBClient{}
}

// Close
func (m *MockDBClient) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}

// Migration
func (m *MockDBClient) Migrate(ctx context.Context) error {
	m.migrateCalls = append(m.migrateCalls, struct{ ctx context.Context }{ctx})
	if m.MigrateFunc != nil {
		return m.MigrateFunc(ctx)
	}
	return nil
}

// Source Data Methods
func (m *MockDBClient) CreateDataSource(ctx context.Context, dataSource *models.DataSource) error {
	m.createDataSourceCalls = append(m.createDataSourceCalls, struct {
		ctx        context.Context
		dataSource *models.DataSource
	}{ctx, dataSource})
	if m.CreateDataSourceFunc != nil {
		return m.CreateDataSourceFunc(ctx, dataSource)
	}
	return nil
}

func (m *MockDBClient) UpdateDataSource(ctx context.Context, dataSource *models.UpdateDataSource) error {
	m.updateDataSourceCalls = append(m.updateDataSourceCalls, struct {
		ctx        context.Context
		dataSource *models.UpdateDataSource
	}{ctx, dataSource})
	if m.UpdateDataSourceFunc != nil {
		return m.UpdateDataSourceFunc(ctx, dataSource)
	}
	return nil
}

func (m *MockDBClient) DeleteDataSource(ctx context.Context, name string) error {
	m.deleteDataSourceCalls = append(m.deleteDataSourceCalls, struct {
		ctx  context.Context
		name string
	}{ctx, name})
	if m.DeleteDataSourceFunc != nil {
		return m.DeleteDataSourceFunc(ctx, name)
	}
	return nil
}

func (m *MockDBClient) GetDataSource(ctx context.Context, name string, params models.DataSourceGetParams) (*models.DataSource, error) {
	m.getDataSourceCalls = append(m.getDataSourceCalls, struct {
		ctx    context.Context
		name   string
		params models.DataSourceGetParams
	}{ctx, name, params})
	if m.GetDataSourceFunc != nil {
		return m.GetDataSourceFunc(ctx, name, params)
	}
	return nil, nil
}

func (m *MockDBClient) ListDataSource(ctx context.Context, params models.DataSourceListParams) ([]*models.DataSource, error) {
	m.listDataSourceCalls = append(m.listDataSourceCalls, struct {
		ctx    context.Context
		params models.DataSourceListParams
	}{ctx, params})
	if m.ListDataSourceFunc != nil {
		return m.ListDataSourceFunc(ctx, params)
	}
	return nil, nil
}

// Detection Section Methods
func (m *MockDBClient) CreateDetectionSection(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	m.createDetectionSectionCalls = append(m.createDetectionSectionCalls, struct {
		ctx              context.Context
		detectionSection *models.DetectionSection
	}{ctx, detectionSection})
	if m.CreateDetectionSectionFunc != nil {
		return m.CreateDetectionSectionFunc(ctx, datasourceName, detectionSection)
	}
	return nil
}

func (m *MockDBClient) CreateManyDetectionSection(ctx context.Context, datasourceName string, detectionSection []*models.DetectionSection) error {
	m.createManyDetectionSectionCalls = append(m.createManyDetectionSectionCalls, struct {
		ctx              context.Context
		detectionSection []*models.DetectionSection
	}{ctx, detectionSection})
	if m.CreateDetectionSectionFunc != nil {
		return m.CreateManyDetectionSectionFunc(ctx, datasourceName, detectionSection)
	}
	return nil
}

func (m *MockDBClient) UpdateDetectionSection(ctx context.Context, datasourceName string, detectionSection *models.DetectionSection) error {
	m.updateDetectionSectionCalls = append(m.updateDetectionSectionCalls, struct {
		ctx              context.Context
		detectionSection *models.DetectionSection
	}{ctx, detectionSection})
	if m.UpdateDetectionSectionFunc != nil {
		return m.UpdateDetectionSectionFunc(ctx, datasourceName, detectionSection)
	}
	return nil
}

func (m *MockDBClient) DeleteDetectionSection(ctx context.Context, datasourceName string, id string) error {
	m.deleteDetectionSectionCalls = append(m.deleteDetectionSectionCalls, struct {
		ctx context.Context
		id  string
	}{ctx, id})
	if m.DeleteDetectionSectionFunc != nil {
		return m.DeleteDetectionSectionFunc(ctx, datasourceName, id)
	}
	return nil
}

func (m *MockDBClient) GetDetectionSection(ctx context.Context, datasourceName string, id string) (*models.DetectionSection, error) {
	m.getDetectionSectionCalls = append(m.getDetectionSectionCalls, struct {
		ctx context.Context
		id  string
	}{ctx, id})
	if m.GetDetectionSectionFunc != nil {
		return m.GetDetectionSectionFunc(ctx, datasourceName, id)
	}
	return nil, nil
}

func (m *MockDBClient) ListDetectionSection(ctx context.Context, datasourceName string) ([]*models.DetectionSection, error) {
	m.listDetectionSectionCalls = append(m.listDetectionSectionCalls, struct {
		ctx context.Context
	}{ctx})
	if m.ListDetectionSectionFunc != nil {
		return m.ListDetectionSectionFunc(ctx, datasourceName)
	}
	return nil, nil
}

// Detection Point Methods
func (m *MockDBClient) CreateDetectionPoint(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	m.createDetectionPointCalls = append(m.createDetectionPointCalls, struct {
		ctx            context.Context
		detectionPoint *models.DetectionPoint
	}{ctx, detectionPoint})
	if m.CreateDetectionPointFunc != nil {
		return m.CreateDetectionPointFunc(ctx, datasourceName, detectionPoint)
	}
	return nil
}

func (m *MockDBClient) CreateManyDetectionPoint(ctx context.Context, datasourceName string, detectionPoint []*models.DetectionPoint) error {
	m.createManyDetectionPointCalls = append(m.createManyDetectionPointCalls, struct {
		ctx            context.Context
		detectionPoint []*models.DetectionPoint
	}{ctx, detectionPoint})
	if m.CreateDetectionPointFunc != nil {
		return m.CreateManyDetectionPointFunc(ctx, datasourceName, detectionPoint)
	}
	return nil
}

func (m *MockDBClient) UpdateDetectionPoint(ctx context.Context, datasourceName string, detectionPoint *models.DetectionPoint) error {
	m.updateDetectionPointCalls = append(m.updateDetectionPointCalls, struct {
		ctx            context.Context
		detectionPoint *models.DetectionPoint
	}{ctx, detectionPoint})
	if m.UpdateDetectionPointFunc != nil {
		return m.UpdateDetectionPointFunc(ctx, datasourceName, detectionPoint)
	}
	return nil
}

func (m *MockDBClient) DeleteDetectionPoint(ctx context.Context, datasourceName string, id string) error {
	m.deleteDetectionPointCalls = append(m.deleteDetectionPointCalls, struct {
		ctx context.Context
		id  string
	}{ctx, id})
	if m.DeleteDetectionPointFunc != nil {
		return m.DeleteDetectionPointFunc(ctx, datasourceName, id)
	}
	return nil
}

func (m *MockDBClient) GetDetectionPoint(ctx context.Context, datasourceName string, id string) (*models.DetectionPoint, error) {
	m.getDetectionPointCalls = append(m.getDetectionPointCalls, struct {
		ctx context.Context
		id  string
	}{ctx, id})
	if m.GetDetectionPointFunc != nil {
		return m.GetDetectionPointFunc(ctx, datasourceName, id)
	}
	return nil, nil
}

func (m *MockDBClient) ListDetectionPoint(ctx context.Context, datasourceName string) ([]*models.DetectionPoint, error) {
	m.listDetectionPointCalls = append(m.listDetectionPointCalls, struct {
		ctx context.Context
	}{ctx})
	if m.ListDetectionPointFunc != nil {
		return m.ListDetectionPointFunc(ctx, datasourceName)
	}
	return nil, nil
}

func (m *MockDBClient) BulkCreateDetectionPoint(ctx context.Context, datasourceName string, detectionPoints []*models.DetectionPoint) error {
	m.bulkCreateDetectionPointCalls = append(m.bulkCreateDetectionPointCalls, struct {
		ctx             context.Context
		detectionPoints []*models.DetectionPoint
	}{ctx, detectionPoints})
	if m.BulkCreateDetectionPointFunc != nil {
		return m.BulkCreateDetectionPointFunc(ctx, datasourceName, detectionPoints)
	}
	return nil
}

// Reset all mock calls and functions
func (m *MockDBClient) Reset() {
	// Reset all call tracking slices
	m.migrateCalls = nil
	m.createDataSourceCalls = nil
	m.updateDataSourceCalls = nil
	m.deleteDataSourceCalls = nil
	m.getDataSourceCalls = nil
	m.listDataSourceCalls = nil
	m.createDetectionSectionCalls = nil
	m.updateDetectionSectionCalls = nil
	m.deleteDetectionSectionCalls = nil
	m.getDetectionSectionCalls = nil
	m.listDetectionSectionCalls = nil
	m.createDetectionPointCalls = nil
	m.updateDetectionPointCalls = nil
	m.deleteDetectionPointCalls = nil
	m.getDetectionPointCalls = nil
	m.listDetectionPointCalls = nil

	// Reset all function implementations
	m.MigrateFunc = nil
	m.CreateDataSourceFunc = nil
	m.UpdateDataSourceFunc = nil
	m.DeleteDataSourceFunc = nil
	m.GetDataSourceFunc = nil
	m.ListDataSourceFunc = nil
	m.CreateDetectionSectionFunc = nil
	m.UpdateDetectionSectionFunc = nil
	m.DeleteDetectionSectionFunc = nil
	m.GetDetectionSectionFunc = nil
	m.ListDetectionSectionFunc = nil
	m.CreateDetectionPointFunc = nil
	m.UpdateDetectionPointFunc = nil
	m.DeleteDetectionPointFunc = nil
	m.GetDetectionPointFunc = nil
	m.ListDetectionPointFunc = nil
}

// Source Data Methods
func (m *MockDBClient) CreateDashboard(ctx context.Context, dashboard *models.Dashboard) error {
	m.createDashboardCalls = append(m.createDashboardCalls, struct {
		ctx       context.Context
		dashboard *models.Dashboard
	}{ctx, dashboard})
	if m.CreateDashboardFunc != nil {
		return m.CreateDashboardFunc(ctx, dashboard)
	}
	return nil
}

func (m *MockDBClient) UpdateDashboard(ctx context.Context, dashboard *models.Dashboard) error {
	m.updateDashboardCalls = append(m.updateDashboardCalls, struct {
		ctx       context.Context
		dashboard *models.Dashboard
	}{ctx, dashboard})
	if m.UpdateDashboardFunc != nil {
		return m.UpdateDashboardFunc(ctx, dashboard)
	}
	return nil
}

func (m *MockDBClient) DeleteDashboard(ctx context.Context, name string) error {
	m.deleteDashboardCalls = append(m.deleteDashboardCalls, struct {
		ctx  context.Context
		name string
	}{ctx, name})
	if m.DeleteDashboardFunc != nil {
		return m.DeleteDashboardFunc(ctx, name)
	}
	return nil
}

func (m *MockDBClient) GetDashboard(ctx context.Context, id string) (*models.Dashboard, error) {
	m.getDashboardCalls = append(m.getDashboardCalls, struct {
		ctx context.Context
		id  string
	}{ctx, id})
	if m.GetDashboardFunc != nil {
		return m.GetDashboardFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockDBClient) ListDashboard(ctx context.Context, dataSourceName string) ([]*models.Dashboard, error) {
	m.listDashboardCalls = append(m.listDashboardCalls, struct {
		ctx            context.Context
		dataSourceName string
	}{ctx, dataSourceName})
	if m.ListDashboardFunc != nil {
		return m.ListDashboardFunc(ctx, dataSourceName)
	}
	return nil, nil
}
