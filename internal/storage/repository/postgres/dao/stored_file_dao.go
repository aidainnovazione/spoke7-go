package dao

import (
	"spoke7-go/internal/storage/models"
	"time"

	"github.com/lib/pq"
)

// DAO for StoredFile

type FileTypeDao int

const (
	FileTypeUnknown FileTypeDao = iota
	FileTypeRealTimeTrafficByLane
	FileTypeRealTimeTrafficByDetectionSection
	FileTypeAggregatedTraffic5MinByLane
	FileTypeAggregatedTraffic5MinByDetectionPoint
	FileTypeAggregatedTraffic5MinByDetectionSection
	FileTypeAggregatedTraffic1HourByLane
	FileTypeAggregatedTraffic1HourByDetectionPoint
	FileTypeAggregatedTraffic1HourByDetectionSection
	FileTypeAggregatedTrafficDayByLane
	FileTypeAggregatedTrafficDayByDetectionPoint
	FileTypeAggregatedTrafficDayByDetectionSection
	FileTypeSumoNetwork
	FileTypeSumoAdditional
	FileTypeSumoRoutes
	FileTypeDetectionSections
	FileTypeDetectionPoints
	FileTypeNetwork
)

type StoredFile struct {
	ID             string `gorm:"primaryKey;autoIncrement"`
	DataSourceName string `gorm:"type:text"`
	Description    string `gorm:"type:text"`
	Tag            string `gorm:"type:text"`

	FileName   string `gorm:"type:text"`
	FileSize   uint32
	FileType   FileTypeDao
	FileFormat string `gorm:"type:text"`

	FileContent []byte `gorm:"type:bytea"`

	Owner  string         `gorm:"type:text"`
	Groups pq.StringArray `gorm:"type:text[]"`

	CreatedAt  time.Time `gorm:"autoCreateTime:milli"`
	ModifiedAt time.Time `gorm:"autoUpdateTime:milli"`
}

// TableName sets the table name for the StoredFile struct
func (StoredFile) TableName() string {
	return "stored_files"
}

func (dao *StoredFile) ToModel() models.StoredFile {

	return models.StoredFile{
		ID:             dao.ID,
		DataSourceName: dao.DataSourceName,
		Description:    dao.Description,
		Tag:            dao.Tag,

		Owner:      dao.Owner,
		Groups:     dao.Groups,
		CreatedAt:  dao.CreatedAt,
		ModifiedAt: dao.ModifiedAt,

		FileName:   dao.FileName,
		FileSize:   dao.FileSize,
		FileType:   models.FileType(dao.FileType),
		FileFormat: dao.FileFormat,

		FileContent: dao.FileContent,
	}
}

func NewStoredFileDaoFromModel(model models.StoredFile) StoredFile {
	return StoredFile{
		ID:             model.ID,
		DataSourceName: model.DataSourceName,
		Description:    model.Description,
		Tag:            model.Tag,

		Owner:      model.Owner,
		Groups:     model.Groups,
		CreatedAt:  model.CreatedAt,
		ModifiedAt: model.ModifiedAt,

		FileName:   model.FileName,
		FileSize:   model.FileSize,
		FileType:   FileTypeDao(model.FileType),
		FileFormat: model.FileFormat,

		FileContent: model.FileContent,
	}
}

func NewStoredFileUploadDaoFromModel(model models.StoredFileUpload) StoredFile {
	return StoredFile{
		DataSourceName: model.DataSourceName,
		Description:    model.Description,
		Tag:            model.Tag,

		Owner:  model.Owner,
		Groups: model.Groups,

		FileName:   model.FileName,
		FileSize:   model.FileSize,
		FileFormat: model.FileFormat,

		FileContent: model.FileContent,
		FileType:    FileTypeDao(model.FileType),
	}
}

func NewStoredFileUpdateDaoFromModel(model models.StoredFileUpdate) StoredFile {
	return StoredFile{
		ID:             model.ID,
		DataSourceName: model.DataSourceName,
		Description:    model.Description,
		Tag:            model.Tag,

		Owner:  model.Owner,
		Groups: model.Groups,

		FileName:   model.FileName,
		FileSize:   model.FileSize,
		FileFormat: model.FileFormat,

		FileContent: model.FileContent,
		FileType:    FileTypeDao(model.FileType),
	}
}
