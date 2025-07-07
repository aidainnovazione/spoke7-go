package config

import "fmt"

type DatabaseConfig struct {
	Type         string `toml:"type" yaml:"type"`
	Host         string `toml:"host" yaml:"host"`
	Port         int    `toml:"port" yaml:"port"`
	Timeout      int    `toml:"timeout" yaml:"timeout"`
	DatabaseName string `toml:"databaseName" yaml:"databaseName"`
	Username     string `toml:"username" yaml:"username"`
	Password     string `toml:"password" yaml:"password"`
	BatchSize    int    `toml:"batchSize" yaml:"batchSize"`
	SSLMode      bool   `toml:"sslMode" yaml:"sslMode"`
	SSLRootCert  string `toml:"sslRootCert" yaml:"sslRootCert"`
}

func (p DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		p.Host, p.Username, p.Password, p.DatabaseName, p.Port,
	)
}

func (c *DatabaseConfig) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("Database.Type is required")
	}
	if c.Host == "" {
		return fmt.Errorf("Database.Host is required")
	}
	if c.Port == 0 {
		return fmt.Errorf("Database.Port is required")
	}
	if c.DatabaseName == "" {
		return fmt.Errorf("Database.DatabaseName is required")
	}
	if c.Username == "" {
		return fmt.Errorf("Database.Username is required")
	}
	if c.Password == "" {
		return fmt.Errorf("Database.Password is required")
	}

	if c.SSLMode {
		if c.SSLRootCert == "" {
			return fmt.Errorf("Database.SSLRootCert is required")
		}
	}
	return nil
}
