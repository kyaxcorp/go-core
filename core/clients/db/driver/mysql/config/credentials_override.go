package config

type CredentialsOverrides struct {
	Host     string `yaml:"host" default:"no"`
	Port     string `yaml:"port" default:"no"`
	User     string `yaml:"user" default:"no"`
	Password string `yaml:"password" default:"no"`
	DbName   string `yaml:"db_name" default:"no"`
}
