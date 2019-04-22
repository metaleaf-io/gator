package conf

import (
	"github.com/metaleaf-io/config"
	"github.com/metaleaf-io/log"
)

type GatorConfig struct {
	// UDP port Gator listens for log messages on.
	Port int16 `json:"port"`
}

// Configuration for the web server.
type ServerConfig struct {
	// Port the web server listens on.
	Port int16 `json:"port"`
}

// Configuration for the database server.
type DatabaseConfig struct {
	// Driver passed to the Go SQL library.
	Driver string `json:"driver"`

	// Hostname the database server listens on.
	Host string `json:"host"`

	// Port the database server listens on.
	Port int16 `json:"port"`

	// Name of the database to use.
	Database string `json:"database"`

	// Username to access the database server
	Username string `json:"username"`

	// Password to access the database server
	Password string `json:"password"`
}


type AppConfig struct {
	Gator       GatorConfig     `json:"gator"`
	Server      ServerConfig    `json:"server"`
	Database    DatabaseConfig  `json:"database"`
}

func LoadConfig(path string) *AppConfig {
	log.Info("Loading configuration", log.String("path", path))
	appConfig := new(AppConfig)
	if err := config.FromFile(path, appConfig); err != nil {
		panic(err)
	}
	return appConfig
}
