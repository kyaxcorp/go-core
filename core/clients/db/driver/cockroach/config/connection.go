package config

import (
	"github.com/kyaxcorp/go-core/core/clients/db/driver"
	"github.com/kyaxcorp/go-core/core/logger/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	CredentialsOverrides CredentialsOverrides `yaml:"credentials_overrides"`
	Credentials          `yaml:"credentials"`

	// This is only for this connection!
	ReconnectOptions `yaml:"reconnect_options"`

	//
	logger       *model.Logger
	masterConfig *Config
}

func (c *Connection) SetLogger(logger *model.Logger) {
	c.logger = logger
}

func (c *Connection) SetMasterConfig(config interface{}) {
	c.masterConfig = config.(*Config)
}

func (c *Connection) GetDialector() gorm.Dialector {
	dsn := c.GenerateDSN()
	c.logger.Info().
		Str("type", "postgresql").
		Str("dsn", dsn.Secured).
		Msg("generating PostgreSQL DSN")
	return postgres.Open(dsn.Plain)
}

func (c *Connection) GetReconnectOptions() driver.ReconnectOptions {
	return &c.ReconnectOptions
}
