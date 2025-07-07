package mocks

import (
	"context"
	"spoke7-go/internal/storage/models"

	"github.com/stretchr/testify/mock"
)

type MockUploadFilesService struct {
	mock.Mock
}

func (m *MockUploadFilesService) UploadFile(ctx context.Context, dataSourceName string, fileName string, fileContent []byte, description string, fileFormat string, fileType models.FileType) error {
	args := m.Called(ctx, dataSourceName, fileName, fileContent, description, fileFormat)
	return args.Error(0)
}
