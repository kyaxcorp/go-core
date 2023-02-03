package config

type CredentialsOverrides struct {
	Host              string `yaml:"host" default:"no"`
	Port              string `yaml:"port" default:"no"`
	User              string `yaml:"user" default:"no"`
	Password          string `yaml:"password" default:"no"`
	DbName            string `yaml:"db_name" default:"no"`
	Schema            string `yaml:"schema" default:"no"`
	TimeZone          string `yaml:"time_zone" default:"no"`
	SSLMode           string `yaml:"ssl_mode" default:"no"`
	SSLFactory        string `yaml:"ssl_factory" default:"no"`
	CACertificate     string `yaml:"ca_certificate" default:"no"`
	ClientCertificate string `yaml:"client_certificate" default:"no"`
	ClientPrivateKey  string `yaml:"client_private_key" default:"no"`
}
