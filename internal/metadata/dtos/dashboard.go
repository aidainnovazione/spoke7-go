package dtos

import (
	"spoke7-go/internal/metadata/models"
	"spoke7-go/internal/metadata/pb"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewDashboardProtoFromModel(model *models.Dashboard) pb.Dashboard {

	return pb.Dashboard{
		Name:           model.Name,
		Description:    model.Description,
		Id:             model.ID,
		DataSourceName: model.DataSourceName,
		Sections:       DashboardSectionModelsToProtos(model.Sections),
		CreatedAt:      timestamppb.New(model.CreatedAt),
		ModifiedAt:     timestamppb.New(model.ModifiedAt),
		Owner:          model.Owner,
		Groups:         model.Groups,
		Sidebar:        DashboardSidebarModelToProto(model.Sidebar),
		Bottombar:      DashboardBottombarModelToProto(model.Bottombar),
		BaseMap:        DashboardBaseMapModelToProto(model.BaseMap),
	}
}

func DashboardProtoToModel(dto *pb.Dashboard) models.Dashboard {

	model := models.Dashboard{
		ID:             dto.Id,
		Name:           dto.Name,
		Description:    dto.Description,
		DataSourceName: dto.DataSourceName,
		Sections:       DashboardSectionProtosToModels(dto.Sections),
		CreatedAt:      dto.CreatedAt.AsTime(),
		ModifiedAt:     dto.ModifiedAt.AsTime(),
		Owner:          dto.Owner,
		Groups:         dto.Groups,
	}

	if dto.Bottombar == nil {
		model.Bottombar = models.Bottombar{}
	} else {
		model.Bottombar = DashboardBottombarProtoToModel(dto.Bottombar)
	}

	if dto.Sidebar == nil {
		model.Sidebar = models.Sidebar{}
	} else {
		model.Sidebar = DashboardSidebarProtoToModel(dto.Sidebar)
	}

	if dto.BaseMap == nil {
		model.BaseMap = models.BaseMap{}
	} else {
		model.BaseMap = DashboardBaseMapProtoToModel(dto.BaseMap)
	}

	return model
}

func DashboardSectionProtosToModels(list []*pb.Section) []models.Section {
	protoStats := make([]models.Section, len(list))
	for i, section := range list {
		protoStatsList := DashboardSectionProtoToModel(section)
		protoStats[i] = protoStatsList
	}
	return protoStats
}

func DashboardSectionProtoToModel(dto *pb.Section) models.Section {
	return models.Section{
		Name:       dto.Name,
		IsVisible:  dto.IsVisible,
		IsExpanded: dto.IsExpanded,
		LayerType:  models.LayerType(dto.LayerType),
		Params:     dto.GetParams().AsMap(),
	}
}

func DashboardSectionModelsToProtos(list []models.Section) []*pb.Section {
	proto := make([]*pb.Section, len(list))
	for i, model := range list {
		protoList := DashboardSectionModelToProto(model)
		proto[i] = protoList
	}
	return proto
}

func DashboardSectionModelToProto(model models.Section) *pb.Section {
	paramsStruct, err := structpb.NewStruct(model.Params)
	if err != nil {
		return nil
	}
	return &pb.Section{
		Name:       model.Name,
		IsVisible:  model.IsVisible,
		IsExpanded: model.IsExpanded,
		LayerType:  pb.LayerType(model.LayerType),
		Params:     paramsStruct,
	}
}

func DashboardBottombarProtoToModel(dto *pb.Bottombar) models.Bottombar {
	return models.Bottombar{
		IsOpen:       dto.IsOpen,
		SelectedTime: dto.SelectedTime.AsTime(),
		StartTime:    dto.StartTime.AsTime(),
		EndTime:      dto.EndTime.AsTime(),
		Interval:     dto.Interval,
	}
}

func DashboardBottombarModelToProto(model models.Bottombar) *pb.Bottombar {
	return &pb.Bottombar{
		IsOpen:       model.IsOpen,
		SelectedTime: timestamppb.New(model.SelectedTime),
		StartTime:    timestamppb.New(model.StartTime),
		EndTime:      timestamppb.New(model.EndTime),
		Interval:     model.Interval,
	}
}

func DashboardSidebarProtoToModel(dto *pb.Sidebar) models.Sidebar {
	return models.Sidebar{
		IsOpen:            dto.IsOpen,
		InjectedComponent: dto.InjectedComponent,
	}
}

func DashboardSidebarModelToProto(model models.Sidebar) *pb.Sidebar {
	return &pb.Sidebar{
		IsOpen:            model.IsOpen,
		InjectedComponent: model.InjectedComponent,
	}
}

func DashboardBaseMapProtoToModel(dto *pb.BaseMap) models.BaseMap {
	return models.BaseMap{
		MapTheme: models.MapTheme(dto.MapTheme),
		Center:   dto.Center,
		Zoom:     dto.Zoom,
	}
}

func DashboardBaseMapModelToProto(model models.BaseMap) *pb.BaseMap {
	return &pb.BaseMap{
		MapTheme: pb.MapTheme(model.MapTheme),
		Center:   model.Center,
		Zoom:     model.Zoom,
	}
}
