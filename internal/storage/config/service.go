package config

import "fmt"

type ServiceConfig struct {
	Host     string `toml:"host" yaml:"host"`
	Port     int64  `toml:"port" yaml:"port"`
	GrpcPort int64  `toml:"grpcPort" yaml:"grpcPort"`
	OpenMsg  string `toml:"openMsg" yaml:"openMsg"`
}

func (c *ServiceConfig) Validate() error {
	if c.Port == 0 {
		return fmt.Errorf("Service.Port is required")
	}
	if c.GrpcPort == 0 {
		return fmt.Errorf("Service.GrpcPort is required")
	}
	return nil
}
