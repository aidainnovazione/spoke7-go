// pkg/configloader/config_loader.go
package configloader

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var once sync.Once

type Configuration interface {
	Validate() error
}

type ConfigLoader struct {
	v           *viper.Viper
	filePath    string
	defaultPort int
}

func NewConfigLoader(defaultFilePath string, defaultPort int) *ConfigLoader {
	v := viper.New()
	return &ConfigLoader{
		filePath:    defaultFilePath,
		defaultPort: defaultPort,
		v:           v,
	}
}

// LoadConfig loads the configuration from file and environment variables.
func (loader *ConfigLoader) LoadConfig(config Configuration) error {

	// Parse flags
	err := loader.ParseFlags()
	if err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	loader.v.SetConfigFile(loader.filePath)
	log.Printf("Loading configuration from: %s\n", loader.filePath)

	loader.v.AutomaticEnv() // override with environment variables

	// Load configuration file
	if err := loader.v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read configuration file: %w", err)
	}
	log.Printf("Config file used: %s\n", loader.v.ConfigFileUsed())

	// Unmarshal into the provided config struct
	if err := loader.v.Unmarshal(config); err != nil {
		return fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return config.Validate()
}

func (loader *ConfigLoader) ParseFlags() error {

	once.Do(func() {
		pflag.StringVarP(&loader.filePath, "conf", "c", loader.filePath, fmt.Sprintf("Specify config file path (default: %s)", loader.filePath))
		pflag.Int16("port", int16(loader.defaultPort), "Specify the port to listen on.")

		pflag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
			pflag.PrintDefaults()
		}
	})

	pflag.Parse()

	// Bind flags to Viper keys
	if err := loader.v.BindPFlag("Service.Port", pflag.Lookup("port")); err != nil {
		return fmt.Errorf("failed to bind flag port: %w", err)
	}

	return nil
}
