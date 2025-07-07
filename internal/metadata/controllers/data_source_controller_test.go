package controllers

import (
	"context"
	"errors"
	"log"
	"spoke7-go/internal/metadata/config"
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"
	"spoke7-go/internal/metadata/services/mocks"
	"spoke7-go/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceHttpController_List(t *testing.T) {

	t.Run("List Source Data Success", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("List", mock.Anything, mock.Anything).Return([]*models.DataSource{}, nil)

		req := &pb.DataSourceListParams{}
		resp, err := controller.List(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("List Source Data Failure", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("List", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		req := &pb.DataSourceListParams{}
		_, err = controller.List(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestDataSourceHttpController_Get(t *testing.T) {
	t.Run("Get Source Data Success", func(t *testing.T) {
		service := &mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(service, logger)

		expectedDataSource := &models.DataSource{Name: "test"} // FIX: Ensure this is a pointer
		service.On("Get", mock.Anything, "test", mock.Anything).Return(expectedDataSource, nil)

		req := &pb.DataSourceGetRequest{Name: "test"}
		resp, err := controller.Get(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Get Source Data Failure", func(t *testing.T) {
		service := &mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(service, logger)

		service.On("Get", mock.Anything, "test", mock.Anything).Return(nil, errors.New("error"))

		req := &pb.DataSourceGetRequest{Name: "test"}
		_, err = controller.Get(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestDataSourceHttpController_Create(t *testing.T) {

	t.Run("Create Source Data Success", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("Create", mock.Anything, mock.Anything).Return(nil)

		req := &pb.DataSource{Name: "test"}
		resp, err := controller.Create(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Create Source Data Failure", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("Create", mock.Anything, mock.Anything).Return(errors.New("error"))

		req := &pb.DataSource{Name: "test"}
		_, err = controller.Create(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestDataSourceHttpController_Update(t *testing.T) {

	req := &pb.UpdateDataSource{
		Name: "test",
	}
	t.Run("Update Source Data Success", func(t *testing.T) {
		service := &mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(service, logger)
		service.On("Update", mock.Anything, mock.Anything).Return(&models.DataSource{Name: "test"}, nil)

		resp, err := controller.Update(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Update Source Data Failure", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("error"))

		_, err = controller.Update(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestDataSourceHttpController_Delete(t *testing.T) {

	t.Run("Delete Source Data Success", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("Delete", mock.Anything, "test").Return(nil)

		req := &pb.DataSourceDeleteRequest{Name: "test"}
		resp, err := controller.Delete(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Delete Source Data Failure", func(t *testing.T) {
		service := mocks.MockDataSourceService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDataSourceController(&service, logger)
		service.On("Delete", mock.Anything, "test").Return(errors.New("error"))

		req := &pb.DataSourceDeleteRequest{Name: "test"}
		_, err = controller.Delete(context.Background(), req)
		assert.Error(t, err)
	})
}
