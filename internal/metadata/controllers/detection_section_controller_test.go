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

	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestDetectionSection_Create(t *testing.T) {
	now := time.Now().UTC()
	// var shape interface{} = nil // Ensure Shape matches
	t.Run("Create Detection Section Success", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)

		service.On("Create", mock.Anything, mock.Anything).Return(nil)

		req := &pb.CreateDetectionSectionRequest{
			DetectionSection: &pb.DetectionSection{
				Id:             "1",
				DataSourceName: "Source1",
				Description:    "Test description",
				StartLatitude:  12.3456,
				StartLongitude: 78.9101,
				EndLatitude:    34.5678,
				EndLongitude:   90.1234,
				Direction:      1,
				// Shape:          nil,
				RoadNetworkId: "RN123",
				CreatedAt:     timestamppb.New(now),
				ModifiedAt:    timestamppb.New(now),
			},
		}

		_, err = controller.CreateDetectionSection(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("Create Detection Section Failure", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)
		ctx := context.Background()
		req := &pb.CreateDetectionSectionRequest{
			DetectionSection: &pb.DetectionSection{
				Id:             "1",
				DataSourceName: "Source1",
				Description:    "Test description",
				StartLatitude:  12.3456,
				StartLongitude: 78.9101,
				EndLatitude:    34.5678,
				EndLongitude:   90.1234,
				Direction:      1,
				Shape:          nil,
				RoadNetworkId:  "RN123",
				CreatedAt:      timestamppb.New(time.Now()),
				ModifiedAt:     timestamppb.New(time.Now()),
			},
		}

		service.On("Create", mock.Anything, mock.Anything).Return(errors.New("service error"))
		_, err = controller.CreateDetectionSection(ctx, req)
		assert.Error(t, err)
	})

}

func TestDetectionSection_Update(t *testing.T) {
	now := time.Now().UTC()

	t.Run("Update Detection Section Success", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)

		service.On("Update", mock.Anything, mock.Anything).Return(nil)

		req := &pb.UpdateDetectionSectionRequest{
			DetectionSection: &pb.DetectionSection{
				Id:             "1",
				DataSourceName: "Updated Source",
				Description:    "Updated description",
				StartLatitude:  12.3456,
				StartLongitude: 78.9101,
				EndLatitude:    34.5678,
				EndLongitude:   90.1234,
				Direction:      2,
				Shape:          nil,
				RoadNetworkId:  "RN123",
				CreatedAt:      timestamppb.New(now.Add(-24 * time.Hour)),
				ModifiedAt:     timestamppb.New(now),
			},
		}

		_, err = controller.UpdateDetectionSection(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("Update Detection Section Failure", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)
		ctx := context.Background()
		req := &pb.UpdateDetectionSectionRequest{
			DetectionSection: &pb.DetectionSection{
				Id:             "1",
				DataSourceName: "Updated Source",
				Description:    "Updated description",
				StartLatitude:  12.3456,
				StartLongitude: 78.9101,
				EndLatitude:    34.5678,
				EndLongitude:   90.1234,
				Direction:      2,
				Shape:          nil,
				RoadNetworkId:  "RN123",
				CreatedAt:      timestamppb.New(time.Now().Add(-24 * time.Hour)),
				ModifiedAt:     timestamppb.New(time.Now()),
			},
		}

		service.On("Update", mock.Anything, mock.Anything).Return(errors.New("service error"))

		_, err = controller.UpdateDetectionSection(ctx, req)
		assert.Error(t, err)
	})

}

func TestDetectionSection_Delete(t *testing.T) {
	req := &pb.DeleteDetectionSectionRequest{
		Id: "1",
	}

	t.Run("Delete Detection Section Success", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)
		service.On("Delete", mock.Anything, mock.Anything).Return(nil)

		_, err = controller.DeleteDetectionSection(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("Delete Detection Section Failure", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)
		service.On("Delete", mock.Anything, mock.Anything).Return(errors.New("service error"))

		_, err = controller.DeleteDetectionSection(context.Background(), req)
		assert.Error(t, err)
	})

}

func TestDetectionSection_Get(t *testing.T) {

	now := time.Now().UTC()

	t.Run("Get Detection Section Success", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)

		service.On("Get", context.Background(), "1").Return(&models.DetectionSection{
			Id:             "1",
			DataSourceName: "Source1",
			Description:    "Test description",
			StartLatitude:  12.3456,
			StartLongitude: 78.9101,
			EndLatitude:    34.5678,
			EndLongitude:   90.1234,
			Direction:      1,
			Shape:          nil,
			RoadNetworkId:  "RN123",
			CreatedAt:      now.Add(-24 * time.Hour),
			ModifiedAt:     now,
		}, nil)

		req := &pb.GetDetectionSectionRequest{
			Id: "1",
		}

		_, err = controller.GetDetectionSection(context.Background(), req)
		assert.NoError(t, err)
	})

	t.Run("Get Detection Section Failure", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)
		service.On("Get", context.Background(), "1").Return(nil, errors.New("service error"))

		req := &pb.GetDetectionSectionRequest{
			Id: "1",
		}

		_, err = controller.GetDetectionSection(context.Background(), req)
		assert.Error(t, err)
	})
}

func TestDetectionSection_List(t *testing.T) {
	now := time.Now().UTC()

	t.Run("List Detection Section Success", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)

		service.On("List", context.Background()).Return([]*models.DetectionSection{
			{
				Id:             "1",
				DataSourceName: "Source1",
				Description:    "Test description",
				StartLatitude:  12.3456,
				StartLongitude: 78.9101,
				EndLatitude:    34.5678,
				EndLongitude:   90.1234,
				Direction:      1,
				Shape:          nil,
				RoadNetworkId:  "RN123",
				CreatedAt:      now.Add(-24 * time.Hour),
				ModifiedAt:     now,
			},
		}, nil)

		req := &pb.ListDetectionSectionsRequest{}

		resp, err := controller.ListDetectionSections(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.DetectionSections, 1)
		assert.Equal(t, "1", resp.DetectionSections[0].Id)
	})

	t.Run("List Detection Section Failure", func(t *testing.T) {
		service := &mocks.MockDetectionSectionService{}
		fileService := &mocks.MockUploadFilesService{}
		logger, err := logger.NewZapLogger(config.AppConfig.Log)
		if err != nil {
			log.Fatalf("Create logger failed. Error:%v\n", err)
		}
		defer logger.Instance.Sync()
		controller := NewDetectionSectionController(service, fileService, logger)
		service.On("List", context.Background()).Return(nil, errors.New("service error"))

		req := &pb.ListDetectionSectionsRequest{}

		_, err = controller.ListDetectionSections(context.Background(), req)
		assert.Error(t, err)
	})

}
