package models

type PermissionFeatures struct {
	Feature     string
	Permissions []Permission
}

type Permission struct {
	PermissionName string
	Type           string
	Action         string
}

type Role struct {
	Name        string
	Permissions []Permission
}
