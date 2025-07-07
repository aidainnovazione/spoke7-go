package authz

import (
	"errors"
	"fmt"
	"spoke7-go/internal/managment/pb"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// IsRestricted checks if a method is marked as restricted in its MethodSecurityScope.
func IsRestricted(fullMethod string) (bool, error) {
	serviceName, methodName, err := ResolveMethod(fullMethod)
	if err != nil {
		return false, errors.New("method authorization scope extension missing or invalid")
	}
	sd, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(serviceName))

	// Convert to ServiceDescriptor
	svcDesc, ok := sd.(protoreflect.ServiceDescriptor)
	if !ok {
		return false, errors.New("descriptor is not a service")
	}

	// Find method descriptor
	md := svcDesc.Methods().ByName(protoreflect.Name(methodName))
	if md == nil {
		return false, fmt.Errorf("method %s not found in service %s", methodName, serviceName)
	}

	// Extract extension
	ext := proto.GetExtension(md.Options(), pb.E_MethodAuthorizationScope)
	optValue, ok := ext.(*pb.MethodSecurityScope)
	if !ok {
		return false, errors.New("method authorization scope extension missing or invalid")
	}

	return optValue.Restricted, nil
}

// GenerateObjectAction builds authz.Object and authz.Action from fullMethod and optional owner/groups.

func GenerateObjectWithOwnerActionFromProto(fullMethod string, owner string, groups []string) (*Object, string, error) {
	objType, action, err := resolveObjectAndAction(fullMethod)
	if err != nil {
		return nil, "", err
	}

	if objType != nil {
		return &Object{Type: *objType,
			Owner:  owner,
			Groups: groups}, *action, nil
	}

	return nil, "", nil
}

func GenerateObjectActionFromProto(fullMethod string) (*Object, string, error) {
	objType, action, err := resolveObjectAndAction(fullMethod)
	if err != nil {
		return nil, "", err
	}
	if objType != nil {
		return &Object{Type: *objType, Owner: "*"}, *action, nil
	}

	return nil, "", nil

}

func resolveObjectAndAction(fullMethod string) (objectType *string, action *string, err error) {
	serviceName, methodName, err := ResolveMethod(fullMethod)
	if err != nil {
		return nil, nil, errors.New("method authorization scope extension missing or invalid")
	}

	sd, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(serviceName))
	if err != nil {
		return nil, nil, err
	}

	svcDesc, ok := sd.(protoreflect.ServiceDescriptor)
	if !ok {
		return nil, nil, errors.New("descriptor is not a service")
	}

	md := svcDesc.Methods().ByName(protoreflect.Name(methodName))
	if md == nil {
		return nil, nil, fmt.Errorf("method %s not found in service %s", methodName, serviceName)
	}

	//no method specification
	ext := proto.GetExtension(md.Options(), pb.E_MethodAuthorizationScope)

	optValue, ok := ext.(*pb.MethodSecurityScope)
	if !ok {
		return nil, nil, errors.New("method authorization scope extension missing or invalid")
	}

	if optValue == nil {
		return nil, nil, nil
	}

	parts := strings.SplitN(optValue.RequiredPermissions, ":", 2)
	if len(parts) < 2 {
		return nil, nil, errors.New("required_permissions format invalid, expected 'type:action'")
	}

	return &parts[0], &parts[1], nil
}

func ResolveMethod(fullMethod string) (string, string, error) {
	// Remove the leading slash
	fullMethod = strings.TrimPrefix(fullMethod, "/")

	// Split "pb.UserService/ListUsers"
	parts := strings.Split(fullMethod, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid full method name: %s", fullMethod)
	}

	// parts[0] is "pb.UserService", parts[1] is "ListUsers"
	serviceName := parts[0]
	methodName := parts[1]

	return serviceName, methodName, nil
}
