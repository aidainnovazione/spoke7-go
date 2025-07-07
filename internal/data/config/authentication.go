package config

import "fmt"

type AuthenticationConfig struct {
	JwksUrl          string   `toml:"jwksUrl" yaml:"jwksUrl"`
	UsernameClaim    string   `toml:"usernameClaim" yaml:"usernameClaim"`
	RolesClaim       string   `toml:"rolesClaim" yaml:"rolesClaim"`
	GroupsClaim      string   `toml:"groupsClaim" yaml:"groupsClaim"`
	OrganizationName string   `toml:"organizationName" yaml:"organizationName"`
	SkipUrls         []string `toml:"skipUrls" yaml:"skipUrls"`
}

func (c *AuthenticationConfig) Validate() error {
	if c.JwksUrl == "" {
		return fmt.Errorf("Authentication.JwksUrl is required")
	}
	if c.UsernameClaim == "" {
		return fmt.Errorf("Authentication.UsernameClaim is required")
	}
	if c.RolesClaim == "" {
		return fmt.Errorf("Authentication.RolesClaim is required")
	}
	return nil
}
