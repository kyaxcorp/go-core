package config

type Credentials struct {
	// Username,Password, DbName are the same for the entire cluster, but it depends if you are using some kind of
	// temporary password solution which changes every specific time... that's when you will need set the credentials
	// per host/node/server
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"26257"`
	User     string `yaml:"user" default:""`
	Password string `yaml:"password" default:""`
	DbName   string `yaml:"db_name" default:""`
	Schema   string `yaml:"schema" default:""`

	TimeZone string `yaml:"time_zone" default:""`

	SSLMode string `yaml:"ssl_mode" default:"require"`
	// SSLFactory - Optional
	SSLFactory string `yaml:"ssl_factory" default:""` // TODO: what's this
	// CACertificate - Optional
	CACertificate string `yaml:"ca_certificate" default:""`
	// ClientCertificate - Optional
	ClientCertificate string `yaml:"client_certificate" default:""`
	// ClientPrivateKey - Optional
	ClientPrivateKey string `yaml:"client_private_key" default:""`
}
