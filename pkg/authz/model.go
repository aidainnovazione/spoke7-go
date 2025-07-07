package authz

import (
	"context"

	"google.golang.org/grpc"
)

/*ABAC Model*/
type User struct {
	Username string
	Roles    []string
	Groups   []string
}

type Object struct {
	Type   string
	Owner  string
	Groups []string
}

// AuthResolver interface must be implemented by each controller.
type AuthResolver interface {
	GetObjectAndActionFromRequest(ctx context.Context, req any, info *grpc.UnaryServerInfo) (*Object, string, error)
}

type AuthzConfig struct {
	ModelPath     string // Path to Casbin model.conf file
	InitRulePath  string // Optional: Path to initial policy file (CSV or Casbin format)
	JwksUrl       string // If you want to validate JWTs here (optional)
	AdminRoleName string // Role name that will get full permissions, e.g. "admin"
}
