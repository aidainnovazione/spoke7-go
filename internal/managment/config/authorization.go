package config

import "fmt"

type Authorization struct {
	ModelPath    string `json:"modelPath" yaml:"modelPath"`
	InitRulePath string `json:"initRule" yaml:"initRule"`
}

func (c Authorization) Validate() error {
	if c.ModelPath == "" {
		return fmt.Errorf("casbin model path is required")
	}

	return nil
}
