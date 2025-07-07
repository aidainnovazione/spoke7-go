package services

import (
	"context"
	"fmt"
	"spoke7-go/internal/managment/models"
	"strings"

	"github.com/casbin/casbin/v2"
)

type RoleService interface {
	ListPermissions() ([]models.PermissionFeatures, error)
	ListRoles() ([]string, error)
	GetRole(name string) (*models.Role, error)

	CreateRole(ctx context.Context, roleName string, permissions []models.Permission) (*models.Role, error)
	UpdateRole(ctx context.Context, roleName string, permissions []models.Permission) (*models.Role, error)
	// DeleteRole
	// AddUserToRole
	// RemoveUserFromRole
	// GetUsersByRole
}

type roleService struct {
	organization string
	application  string
	enforcer     casbin.IEnforcer
}

func NewRoleService(enforcer casbin.IEnforcer) RoleService {
	return &roleService{enforcer: enforcer}
}

func (s *roleService) ListRoles() ([]string, error) {
	//GetAllRoles gets the list of roles that show up in the current policy.
	return s.enforcer.GetAllRoles()
}

func (s *roleService) GetRole(name string) (*models.Role, error) {
	completeRoleName := fmt.Sprintf("role:%s", name)
	filteredGroupingPolicy, err := s.enforcer.GetFilteredGroupingPolicy(0, completeRoleName)
	if err != nil {
		return nil, err
	}

	role := *&models.Role{
		Name: name,
	}
	for _, gPolicy := range filteredGroupingPolicy {
		if len(gPolicy) < 2 {
			continue
		}
		filteredPolicy, err := s.enforcer.GetFilteredPolicy(0, gPolicy[1])
		if err != nil {
			return nil, err
		}
		for _, policy := range filteredPolicy {
			if len(policy) < 3 {
				continue
			}

			role.Permissions = append(role.Permissions, models.Permission{
				PermissionName: policy[0],
				Type:           policy[1],
				Action:         policy[2],
			})
		}
	}
	return &role, nil
}

func (s *roleService) ListPermissions() ([]models.PermissionFeatures, error) {
	policies, err := s.enforcer.GetPolicy()
	if err != nil {
		return nil, err
	}

	permissionSet := make(map[string]models.PermissionFeatures)

	for _, policy := range policies {
		if len(policy) < 3 {
			continue
		}
		featureName := policy[0]
		resourceType := policy[1] // e.g., "user", "group"
		action := policy[2]       // e.g., "read", "create"

		if strings.HasPrefix(featureName, "feature:") {
			permission := models.Permission{
				Type:           resourceType,
				Action:         action,
				PermissionName: featureName,
			}
			splitFeature := strings.Split(featureName, ":")
			//save feature under feature:object
			featureSubjectName := fmt.Sprintf("%s:%s", splitFeature[0], splitFeature[1])

			permissionSet[featureSubjectName] = models.PermissionFeatures{
				Feature:     featureSubjectName,
				Permissions: append(permissionSet[featureSubjectName].Permissions, permission),
			}
		}

	}

	permissions := make([]models.PermissionFeatures, 0, len(permissionSet))
	for _, perm := range permissionSet {
		permissions = append(permissions, perm)
	}

	return permissions, nil
}

func (u *roleService) CreateRole(ctx context.Context, roleName string, permissions []models.Permission) (*models.Role, error) {
	completeRoleName := fmt.Sprintf("role:%s", roleName)
	// Add each permission to the Casbin policy
	for _, p := range permissions {
		// Add policy: roleName can perform action on type
		ok, err := u.enforcer.AddGroupingPolicy(completeRoleName, p.PermissionName)
		if err != nil {
			return nil, fmt.Errorf("failed to add policy for role %s: %w", completeRoleName, err)
		}
		if !ok {
			// Policy already exists
			continue
		}
	}

	// Optionally persist to storage (if not using AutoSave)
	if err := u.enforcer.SavePolicy(); err != nil {
		return nil, fmt.Errorf("failed to save policy: %w", err)
	}

	// Return the Role struct
	role := &models.Role{
		Name:        completeRoleName,
		Permissions: permissions,
	}

	return role, nil
}

func (u *roleService) UpdateRole(ctx context.Context, roleName string, permissions []models.Permission) (*models.Role, error) {
	completeRoleName := fmt.Sprintf("role:%s", roleName)
	// Get all policies where role is the subject
	currentPermissions, err := u.enforcer.GetFilteredGroupingPolicy(0, completeRoleName)

	if err != nil {
		return nil, fmt.Errorf("failed to get current policies for role %s: %w", completeRoleName, err)
	}

	permissionSet := make(map[string]struct{}, len(permissions))
	for _, p := range permissions {
		permissionSet[p.PermissionName] = struct{}{}
	}

	for _, p := range currentPermissions {
		if len(p) < 3 {
			continue // malformed policy
		}
		key := fmt.Sprintf("%s", p[0])
		if _, ok := permissionSet[key]; !ok {
			// Remove policy: completeRoleName can no longer perform action on type
			_, err = u.enforcer.RemoveGroupingPolicy(completeRoleName, p[0])
			if err != nil {
				return nil, fmt.Errorf("failed to remove policy for role %s: %w", completeRoleName, err)
			}
		}
	}

	for _, p := range permissions {
		// Add policy: completeRoleName can perform action on type
		ok, err := u.enforcer.AddGroupingPolicy(completeRoleName, p.PermissionName)
		if err != nil {
			return nil, fmt.Errorf("failed to add policy for role %s: %w", completeRoleName, err)
		}
		if !ok {
			// Policy already exists
			continue
		}
	}

	// Optionally persist to storage (if not using AutoSave)
	if err := u.enforcer.SavePolicy(); err != nil {
		return nil, fmt.Errorf("failed to save policy: %w", err)
	}

	// Return the Role struct
	role := &models.Role{
		Name:        completeRoleName,
		Permissions: permissions,
	}

	return role, nil
}

func (u *roleService) GetRoleBySubject(ctx context.Context, subject string) (*models.Role, error) {
	// Get all policies where role is the subject
	policies, err := u.enforcer.GetFilteredPolicy(0, subject)

	if err != nil {
		return nil, fmt.Errorf("failed to save policy: %w", err)
	}

	permissions := make([]models.Permission, 0, len(policies))

	for _, p := range policies {
		if len(p) < 3 {
			continue // malformed policy
		}
		permissions = append(permissions, models.Permission{
			Type:   p[1], // obj
			Action: p[2], // act
		})
	}

	return &models.Role{
		Name:        subject,
		Permissions: permissions,
	}, nil
}
