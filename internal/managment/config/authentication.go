package config

import "fmt"

type AuthenticationConfig struct {
	Endpoint         string   `toml:"endpoint" yaml:"endpoint"`
	ClientId         string   `toml:"clientId" yaml:"clientId"`
	ClientSecret     string   `toml:"clientSecret" yaml:"clientSecret"`
	Certificate      string   `toml:"certificate" yaml:"certificate"`
	OrganizationName string   `toml:"organizationName" yaml:"organizationName"`
	ApplicationName  string   `toml:"applicationName" yaml:"applicationName"`
	JwksUrl          string   `toml:"jwksUrl" yaml:"jwksUrl"`
	UsernameClaim    string   `toml:"usernameClaim" yaml:"usernameClaim"`
	RolesClaim       string   `toml:"rolesClaim" yaml:"rolesClaim"`
	GroupsClaim      string   `toml:"groupsClaim" yaml:"groupsClaim"`
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
	if c.GroupsClaim == "" {
		return fmt.Errorf("Authentication.GroupsClaim is required")
	}
	if c.Endpoint == "" {
		return fmt.Errorf("Authentication.Endpoint is required")
	}
	if c.ClientId == "" {
		return fmt.Errorf("Authentication.ClientId is required")
	}
	if c.ClientSecret == "" {
		return fmt.Errorf("Authentication.ClientSecret is required")
	}
	if c.Certificate == "" {
		return fmt.Errorf("Authentication.Certificate is required")
	}
	if c.OrganizationName == "" {
		return fmt.Errorf("Authentication.Organization is required")
	}
	if c.ApplicationName == "" {
		return fmt.Errorf("Authentication.Application is required")
	}

	return nil
}
