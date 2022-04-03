package constructor

import (
	"github.com/KyaXTeam/go-core/v2/core/clients/db/codes"
	"github.com/KyaXTeam/go-core/v2/core/clients/db/dbinstance"
	dbDriver "github.com/KyaXTeam/go-core/v2/core/clients/db/driver"
	cockroachConfig "github.com/KyaXTeam/go-core/v2/core/clients/db/driver/cockroach/config"
	mysqlConfig "github.com/KyaXTeam/go-core/v2/core/clients/db/driver/mysql/config"
	mainConfig "github.com/KyaXTeam/go-core/v2/core/config"
	"github.com/KyaXTeam/go-core/v2/core/helpers/err"
	// mysqlConfig "github.com/KyaXTeam/go-core/v2/core/clients/db/driver/mysql/config"
	"gorm.io/gorm"
)

func (dbc *DBClient) GetDriverInstance() *dbinstance.Instance {
	// This is for reading...
	driverInstancesLock.RLock()
	if instance, ok := driverInstances[dbc.DriverType]; ok {
		driverInstancesLock.RUnlock()
		return instance
	}
	driverInstancesLock.RUnlock()

	// This is for creating
	driverInstancesLock.Lock()
	driverInstances[dbc.DriverType] = dbinstance.NewInstance()
	// Define the auto creator!
	driverInstances[dbc.DriverType].OnMissingAutoCreate = func(instanceName string) (*gorm.DB, error) {
		c, _e := dbc.NewClient(
			instanceName,
		)

		if _e != nil {
			return nil, _e
		}
		return c, nil
	}
	// Set just as reference for faster access
	dbc.instanceRef = driverInstances[dbc.DriverType]
	driverInstancesLock.Unlock()
	return dbc.instanceRef
	// That's for responding...
	//driverInstancesLock.RLock()
	//defer driverInstancesLock.RUnlock()
	//return driverInstances[dbc.DriverType]
}

// New -> Creates totally a new DB Client, this is the root function!
func (dbc *DBClient) New(
	config dbDriver.Config,
) (*gorm.DB, error) {
	return NewClient(dbc.Ctx, config)
}

// GetDefaultClient -> it returns the default client based on the existing app configuration
func (dbc *DBClient) GetDefaultClient() (*gorm.DB, error) {
	conf := mainConfig.GetConfig()
	//defaultConnName := conf.Clients.MySQL.DefaultConn.Name

	// TODO: get driverType and get from config

	var defaultInstanceName string
	switch dbc.DriverType {
	case MySQLDriver:
		defaultInstanceName = conf.Clients.MySQL.DefaultConn.InstanceId
	case CockroachDriver:
		defaultInstanceName = conf.Clients.Cockroach.DefaultConn.InstanceId
	case SQLiteDriver:

	}

	if defaultInstanceName == "" {
		return nil, codes.ErrDefaultDbInstanceNameIsEmpty
	}

	return dbc.GetDriverInstance().GetClientByInstanceId(defaultInstanceName)
}

// GetClient -> it returns an existing client
func (dbc *DBClient) GetClient(
	instanceName string,
) (*gorm.DB, error) {
	if instanceName == "" {
		return nil, err.New(0, "mysql client instance name is empty")
	}
	// check if there is an existing instance
	srv, _err := dbc.GetDriverInstance().GetClientByInstanceId(instanceName)
	if _err == nil && srv != nil {
		return srv, nil
	}
	// If the client is missing, then we will create it automatically

	//if _err == codes.ErrDbInstanceIsMissing {
	//	srv, _err = dbc.NewClient(instanceName)
	//	if _err != nil {
	//		return nil, _err
	//	}
	//	// Save the instance
	//	dbc.GetDriverInstance().SaveClientToInstances(instanceName, srv)
	//}

	return nil, nil
}

// NewClient -> it creates a new DB Client based on the existing app configuration
func (dbc *DBClient) NewClient(
	instanceName string,
) (*gorm.DB, error) {
	// Check if instance name is ok
	if instanceName == "" {
		return nil, err.NewDefined(codes.ErrDbClientInstanceNameEmpty)
	}

	// Get configuration
	if cfg, _err := dbc.GetInstanceConfig(instanceName); _err == nil {
		// If config is found... then create a new client
		return dbc.New(cfg)
	}
	return nil, err.NewDefined(codes.ErrDbInstanceConfigurationIsMissing)
}

//
func (dbc *DBClient) GetDefaultGeneratedConfig() (dbDriver.Config, error) {
	switch dbc.DriverType {
	case MySQLDriver:
		return mysqlConfig.SetDefaults(nil)
	case CockroachDriver:
		return cockroachConfig.SetDefaults(nil)
	}
	return nil, err.NewDefined(codes.ErrDriverNoDefaultConfigFound)
}

//
func (dbc *DBClient) GetInstanceConfig(instanceName string) (dbDriver.Config, error) {

	switch dbc.DriverType {
	case MySQLDriver:
		if cfg, ok := mainConfig.GetConfig().Clients.MySQL.Instances[instanceName]; ok {
			return &cfg, nil
		}
	case CockroachDriver:
		if cfg, ok := mainConfig.GetConfig().Clients.Cockroach.Instances[instanceName]; ok {
			return &cfg, nil
		}
	}

	return nil, err.NewDefined(codes.ErrDbInstanceConfigurationIsMissing)
}
