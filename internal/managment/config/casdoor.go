package config

/*
endpoint	Yes	Casdoor server URL, such as http://localhost:8000
clientId	Yes	Application.clientId
clientSecret	Yes	Application.clientSecret
certificate	Yes	x509 certificate content of Application.cert
organizationName	Yes	Application.organization
applicationName	Yes	Application.applicationName
*/

type CasdoorConfig struct {
	Endpoint         string `toml:"endpoint" yaml:"endpoint"`                 // Casdoor server URL
	ClientID         string `toml:"clientId" yaml:"clientId"`                 // Application.clientId
	ClientSecret     string `toml:"clientSecret" yaml:"clientSecret"`         // Application.clientSecret
	Certificate      string `toml:"certificate" yaml:"certificate"`           // x509 certificate content of Application.cert
	OrganizationName string `toml:"organizationName" yaml:"organizationName"` // Application.organization
	ApplicationName  string `toml:"applicationName" yaml:"applicationName"`   // Application.applicationName
}
