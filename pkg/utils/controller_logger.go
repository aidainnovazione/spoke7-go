package utils

import (
	"context"
	"fmt"
	"spoke7-go/pkg/authz"
	"spoke7-go/pkg/logger"
	"spoke7-go/pkg/request_id"
	"strings"

	"go.uber.org/zap"
)

const MicroserviceString string = "microservice"
const ResourceString string = "resource"
const MethodString string = "method"
const DatasourceNameString string = "datasource_name"
const UsernameString string = "username"
const GroupsString string = "groups"

func AddControllerRequestLogging(logger logger.Logger, ctx context.Context, logText string, microservice string, resource string, method string, name string, additionalFields ...zap.Field) error {
	var requestID string

	if ctx.Value(request_id.RequestIdCtxKey) != nil {
		requestID = ctx.Value(request_id.RequestIdCtxKey).(string)
	}
	user, err := authz.GetUserFromContext(ctx)

	if err != nil {
		errorText := fmt.Sprintf("%v: impossible retrieve user from context", logText)
		loggerFields := []zap.Field{
			zap.String(MicroserviceString, microservice),
			zap.String(request_id.RequestIdCtxKey, requestID),
			zap.String(ResourceString, resource),
			zap.String(MethodString, method),
			zap.String(DatasourceNameString, name),
			zap.String(UsernameString, "NONE"),
			zap.String(GroupsString, "NONE"),
		}
		loggerFields = append(loggerFields, additionalFields...)
		logger.ZapError(errorText, loggerFields...)
		return err
	}

	loggerFields := []zap.Field{
		zap.String(MicroserviceString, microservice),
		zap.String(request_id.RequestIdCtxKey, requestID),
		zap.String(ResourceString, resource),
		zap.String(MethodString, method),
		zap.String(DatasourceNameString, name),
		zap.String(UsernameString, user.Username),
		zap.String(GroupsString, strings.Join(user.Groups, ",")),
	}
	loggerFields = append(loggerFields, additionalFields...)
	logger.ZapInfo(logText, loggerFields...)

	return nil
}

func AddControllerErrorLogging(logger logger.Logger, ctx context.Context, logText string, microservice string, resource string, method string, name string, additionalFields ...zap.Field) error {
	var requestID string

	if ctx.Value(request_id.RequestIdCtxKey) != nil {
		requestID = ctx.Value(request_id.RequestIdCtxKey).(string)
	}
	user, err := authz.GetUserFromContext(ctx)
	if err != nil {
		errorText := fmt.Sprintf("%v: impossible retrieve user from context", logText)
		loggerFields := []zap.Field{
			zap.String(MicroserviceString, microservice),
			zap.String(request_id.RequestIdCtxKey, requestID),
			zap.String(ResourceString, resource),
			zap.String(MethodString, method),
			zap.String(DatasourceNameString, name),
			zap.String(UsernameString, "NONE"),
			zap.String(GroupsString, "NONE"),
		}
		loggerFields = append(loggerFields, additionalFields...)
		logger.ZapError(errorText, loggerFields...)
		return err
	}

	loggerFields := []zap.Field{
		zap.String(MicroserviceString, microservice),
		zap.String(request_id.RequestIdCtxKey, requestID),
		zap.String(ResourceString, resource),
		zap.String(MethodString, method),
		zap.String(DatasourceNameString, name),
		zap.String(UsernameString, user.Username),
		zap.String(GroupsString, strings.Join(user.Groups, ",")),
	}
	loggerFields = append(loggerFields, additionalFields...)
	logger.ZapError(logText, loggerFields...)

	return nil
}

func fieldsToInterfaces(fields []zap.Field) []interface{} {
	result := make([]interface{}, len(fields))
	for i, f := range fields {
		result[i] = f
	}
	return result
}
