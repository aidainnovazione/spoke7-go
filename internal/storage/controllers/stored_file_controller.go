package controllers

import (
	"context"
	"errors"
	"spoke7-go/internal/storage/dtos"
	"spoke7-go/internal/storage/pb"
	"spoke7-go/internal/storage/services"

	errorsInternal "spoke7-go/internal/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoredFileHttpController struct {
	service services.StoredFileService
	pb.UnimplementedStoredFileServiceServer
}

// func NewStoredFileController(service services.StoredFileService) pb.StoredFileServiceServer {
// 	return &StoredFileHttpController{service: service}
// }

func NewStoredFileController(service services.StoredFileService) *StoredFileHttpController {
	return &StoredFileHttpController{service: service}
}

func (sc *StoredFileHttpController) List(ctx context.Context, req *pb.StoredFileListParams) (*pb.StoredFileListResponse, error) {
	list, err := sc.service.List(ctx, req.DataSourceName, req.Tag)
	if err != nil {
		return nil, err
	}

	response := make([]*pb.StoredFile, 0)
	for _, storedFile := range list {
		dto := dtos.NewStoredFileProtoFromModel(storedFile)
		response = append(response, &dto)
	}

	return &pb.StoredFileListResponse{Storage: response}, nil
}

func (sc *StoredFileHttpController) Get(ctx context.Context, req *pb.StoredFileGetRequest) (*pb.StoredFile, error) {
	Id := req.Id
	storedFile, err := sc.service.Get(ctx, Id)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewStoredFileProtoFromModel(storedFile)
	return &response, nil
}

func (sc *StoredFileHttpController) Download(ctx context.Context, req *pb.StoredFileDownloadRequest) (*pb.StoredFileDownloadResponse, error) {
	Id := req.Id
	storedFile, err := sc.service.Get(ctx, Id)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &pb.StoredFileDownloadResponse{
		FileName:   storedFile.FileName,
		FileFormat: storedFile.FileFormat,
		Content:    storedFile.FileContent,
	}, nil
}

func (sc *StoredFileHttpController) Upload(ctx context.Context, req *pb.StoredFileUploadRequest) (*pb.StoredFile, error) {
	storedFile := dtos.UploadStoredFileProtoToModel(req)
	storedFileCreated, err := sc.service.Create(ctx, &storedFile)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewStoredFileProtoFromModel(storedFileCreated)
	return &response, nil
}

func (sc *StoredFileHttpController) Update(ctx context.Context, req *pb.StoredFileUpdateRequest) (*pb.StoredFile, error) {
	storedFile := dtos.UpdateStoredFileProtoToModel(req)
	storedFileUpdated, err := sc.service.Update(ctx, &storedFile)
	if err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	response := dtos.NewStoredFileProtoFromModel(storedFileUpdated)
	return &response, nil
}

func (sc *StoredFileHttpController) Delete(ctx context.Context, req *pb.StoredFileDeleteRequest) (*emptypb.Empty, error) {

	if err := sc.service.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, errorsInternal.ErrForbidden) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
