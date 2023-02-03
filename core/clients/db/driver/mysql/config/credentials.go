package config

type Credentials struct {
	// Username,Password, DbName are the same for the entire cluster, but it depends if you are using some kind of
	// temporary password solution which changes every specific time... that's when you will need set the credentials
	// per host/node/server
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"3306"`
	User     string `yaml:"user" default:""`
	Password string `yaml:"password" default:""`
	DbName   string `yaml:"db_name" default:""`
}
