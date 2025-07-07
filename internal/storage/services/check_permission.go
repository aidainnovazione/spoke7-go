package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/errors"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/grpc_client"
	"strings"
)

func CheckDataSourceUserPermission(ctx context.Context, dataSourceName string, organizationName string, grpcMetadataService grpc_client.GrpcMetadataClient) error {
	if dataSourceName == "" {
		return errors.ErrMissingDataSourceName
	}

	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		return errors.ErrUnauthorized
	}

	datasources, err := grpcMetadataService.ListDataSource(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve data source information: %w", err)
	}

	var groups []string
	for _, groupStr := range user.Groups {
		groupStr = strings.TrimPrefix(groupStr, organizationName+"/")
		groups = append(groups, groupStr)
	}

	var found bool
	for _, datasource := range datasources {
		if datasource.Name == dataSourceName {
			found = true
			if datasource.Groups != nil && !groupsContains(groups, datasource.Groups) && datasource.Owner != user.Username {
				return errors.ErrForbidden
			}
		}
	}

	if !found {
		return errors.ErrForbidden
	}

	return nil
}

func sliceContains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func groupsContains(originalGroups []string, groupsToCheck []string) bool {
	for _, group := range groupsToCheck {
		if sliceContains(originalGroups, group) {
			return true
		}
	}
	return false
}
