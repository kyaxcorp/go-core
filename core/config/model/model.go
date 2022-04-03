package model

import (
	brokerClientConfig "github.com/kyaxcorp/go-core/core/clients/broker/config"
	// cassandraConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/cassandra/config"
	cockroachConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/cockroach/config"
	mysqlConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/mysql/config"
	//sqliteConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/sqlite/config"
	websocketClientConfig "github.com/kyaxcorp/go-core/core/clients/websocket/config"
	httpConfig "github.com/kyaxcorp/go-core/core/listeners/http/config"
	websocketConfig "github.com/kyaxcorp/go-core/core/listeners/websocket/config"
	loggingConfig "github.com/kyaxcorp/go-core/core/logger/config"
	brokerConfig "github.com/kyaxcorp/go-core/core/services/broker/config"
)

// Here we will be storing default configuration key from the map!
//var DefaultConfigName string

// Here we will store different configurations with different models!
//var MainConfigMap map[string]interface{}

//var MainConfig = &NonPtrObj{}

// Are used for the main purpose of the Programmer... he/she can add there custom structures
// which are reflected to their APP model!

type CustomConfigParams struct {
	ConfigModel interface{}
	Config      interface{}
}

const CustomConfigFileName = "config_custom"
const CustomConfigFileType = "yaml"

const ConfigFileName = "config"
const ConfigTmpFileName = "config.tmp"
const ConfigFileType = "yaml"

// Tag mapstructure - how it's being read from file!
// yaml and after that name, it's how it's saved!
// default is the default value of the element, check -> https://github.com/creasty/defaults for docs

// TODO: should i set the entire config to work on Pointers?!
// TODO: So when somethings changes, it will be instantly available?!
// TODO: but still, it's a bit dangerous...

type Model struct {
	// Application params (Which are custom defined by the user
	Application struct {
		// App Data Path (all other subfolders will be stored there...)
		AppDataPath string `yaml:"app_data_path" mapstructure:"app_data_path" default:".appdata"`
		// Where the PIDs are saved!
		PIDsPath string `yaml:"pids_path" mapstructure:"pids_path" default:"pids"`
		// Where the FileLocks are saved
		LocksPath string `yaml:"locks_path" mapstructure:"locks_path" default:"locks"`
		// Temporary folder...
		TempPath string `yaml:"temp_path" mapstructure:"temp_path" default:"tmp"`
		// Storage Path -> where the data will be located
		StoragePath string `yaml:"storage_path" mapstructure:"storage_path" default:"storage"`
		// Cache Path -> where the caching files are being stored
		CachePath string `yaml:"cache_path" mapstructure:"cache_path" default:"cache"`
		// Certs Path -> where we store certificates for different use cases
		CertsPath string `yaml:"certs_path" mapstructure:"certs_path" default:"certs"`
		// ConfigsBackupPath -> we store here the backups of the configs... the app will automatically backup the
		// current config before launching the autoloader! Each time the app starts... it will do this backup!
		// This saves everyone in case of programming incorrection
		ConfigsBackupPath string `yaml:"configs_backup_path" mapstructure:"configs_backup_path" default:"configs_backup"`
		// How many seconds should wait until the app gracefully shuts down all the services before os.exit(0)
		OnShutdownWaitSeconds int `yaml:"on_shutdown_wait_seconds" mapstructure:"on_shutdown_wait_seconds" default:"2"`
		// This is the default timezone... wherever it runs, when the app starts, it will override the timezone with this value
		TimeZone string `yaml:"time_zone" mapstructure:"time_zone" default:"UTC"`
		// This param disables the override function on bootstrap only if necessary!
		DisableTimezoneOverride string `yaml:"disable_timezone_override" mapstructure:"disable_timezone_override" default:"no"`
	}
	// Different connections to different services
	Clients struct {
		DefaultDBClient string `yaml:"default_db_client" mapstructure:"" default:"cockroach"`
		MySQL           struct {
			DefaultConn struct {
				// This is the default Instance Name
				InstanceId string `yaml:"instance_id" mapstructure:"instance_id" default:"default"`
			}
			Instances map[string]mysqlConfig.Config
			//Instances2 map[string]mysqlConfig.Config
		}
		Cockroach struct {
			// This is the default connection name -> from which we
			DefaultConn struct {
				// This is the default Instance Name
				InstanceId string `yaml:"instance_id" mapstructure:"instance_id" default:"default"`
			}
			Instances map[string]cockroachConfig.Config
		}
		/*SQLite struct {
			DefaultConn struct {
				// This is the default Instance Name
				InstanceId string `yaml:"instance_id" mapstructure:"instance_id" default:"default"`
			}
			Instances map[string]sqliteConfig.Config
		}
		*/
		/*Cassandra struct {
			DefaultConn struct {
				// This is the default Instance Name
				InstanceId string `yaml:"instance_id" mapstructure:"instance_id" default:"default"`
			}
			Instances map[string]cassandraConfig.Config
		}*/
		Redis struct {
			Instances struct {
			}
		}
		ElasticSearch struct {
			Instances struct {
			}
		}
		// Remote Procedure Call -> Through HTTP (For interprocess Communication)
		RPC struct {
			// Channel Name
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           struct {
			}
		}
		Broker struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           map[string]brokerClientConfig.Config
		}
		WebSocket struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           map[string]websocketClientConfig.Config
		}
	}
	// The main Console!
	Console struct {
	}
	// Specific Autonomous services which are isolated and separately!
	Services struct {
		Broker struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           map[string]brokerConfig.Config
		}
	}

	// The one's that are listening on some ports (Servers)
	Listeners struct {
		Http struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           map[string]httpConfig.Config
		}
		WebSocket struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           map[string]websocketConfig.Config
		}
		UDP struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           struct {
			}
		}
		UNIX struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           struct {
			}
		}
		TCP struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           struct {
			}
		}
		// Remote Procedure Call -> Through HTTP (For interprocess Communication)
		RPC struct {
			DefaultInstanceName string `yaml:"default_instance_name" mapstructure:"default_instance_name" default:"default"`
			Instances           struct {
			}
		}
	}

	Logging struct {
		LogsPath string `yaml:"logs_path" mapstructure:"logs_path" default:"logs"`
		// This is the default channel
		DefaultChannel string `yaml:"default_channel" mapstructure:"default_channel" default:"default"`
		// A default channel will always be and will be created automatically
		Channels map[string]loggingConfig.Config
	}
}
