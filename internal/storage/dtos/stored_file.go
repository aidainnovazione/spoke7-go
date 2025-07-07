package dtos

import (
	"spoke7-go/internal/storage/models"
	"spoke7-go/internal/storage/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewStoredFileProtoFromModel(model *models.StoredFile) pb.StoredFile {
	return pb.StoredFile{
		Id:             model.ID,
		DataSourceName: model.DataSourceName,
		Description:    model.Description,
		Tag:            model.Tag,
		Owner:          model.Owner,
		Groups:         model.Groups,
		CreatedAt:      timestamppb.New(model.CreatedAt),
		ModifiedAt:     timestamppb.New(model.ModifiedAt),
		FileName:       model.FileName,
		FileSize:       model.FileSize,
		FileType:       pb.FileType(model.FileType),
		FileFormat:     model.FileFormat,
	}
}

func StoredFileProtoToModel(dto *pb.StoredFile) models.StoredFile {
	return models.StoredFile{
		ID:             dto.Id,
		DataSourceName: dto.DataSourceName,
		Description:    dto.Description,
		Tag:            dto.Tag,
		Owner:          dto.Owner,
		Groups:         dto.Groups,
		CreatedAt:      dto.CreatedAt.AsTime(),
		ModifiedAt:     dto.ModifiedAt.AsTime(),
		FileName:       dto.FileName,
		FileSize:       dto.FileSize,
		FileType:       models.FileType(dto.FileType),
		FileFormat:     dto.FileFormat,
	}
}

func UploadStoredFileProtoToModel(dto *pb.StoredFileUploadRequest) models.StoredFileUpload {
	if dto == nil {
		return models.StoredFileUpload{}
	}

	return models.StoredFileUpload{
		DataSourceName: dto.DataSourceName,
		Description:    dto.Description,
		Tag:            dto.Tag,
		FileName:       dto.FileName,
		FileSize:       dto.FileSize,
		FileFormat:     dto.FileFormat,
		FileType:       models.FileType(dto.FileType),
		FileContent:    dto.FileContent,
	}
}

func UpdateStoredFileProtoToModel(dto *pb.StoredFileUpdateRequest) models.StoredFileUpdate {
	if dto == nil {
		return models.StoredFileUpdate{}
	}

	return models.StoredFileUpdate{
		ID:             dto.Id,
		DataSourceName: dto.DataSourceName,
		Description:    dto.Description,
		Tag:            dto.Tag,
		FileName:       dto.FileName,
		FileSize:       dto.FileSize,
		FileFormat:     dto.FileFormat,
		FileType:       models.FileType(dto.FileType),
		FileContent:    dto.FileContent,
	}
}
