package db

import (
	"context"
	dbClient "github.com/kyaxcorp/go-core/core/clients/db/constructor"
	dbDriver "github.com/kyaxcorp/go-core/core/clients/db/driver"
	cockroachConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/cockroach/config"
	mysqlConfig "github.com/kyaxcorp/go-core/core/clients/db/driver/mysql/config"
	"github.com/kyaxcorp/go-core/core/config"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"gorm.io/gorm"
)

// here we store the default clients for specific drivers... (for default instances only)

var DefaultInstanceClient *gorm.DB
var CRDBDefaultInstanceClient *gorm.DB
var MySQLDefaultInstanceClient *gorm.DB

// If we want to use other instances or multiple ones... maybe we should define these instances at bootstrap
// and then use them there as predefined...
// or try using cmap

//------------------HELPERS------------------------\\

func GetDefaultClientDriverName() string {
	var driver string
	switch config.GetConfig().Clients.DefaultDBClient {
	case dbClient.CockroachDriver:
		driver = dbClient.CockroachDriver
	case dbClient.SQLiteDriver:
		driver = dbClient.SQLiteDriver
	case dbClient.MySQLDriver:
		driver = dbClient.MySQLDriver
	case dbClient.PostGRESDriver:
		driver = dbClient.PostGRESDriver
	}
	return driver
}

func getDefaultDBClient() *dbClient.DBClient {
	// get the object, and we will set the driver type in each function separately
	cl := &dbClient.DBClient{
		Ctx:        _context.GetDefaultContext(),
		DriverType: GetDefaultClientDriverName(),
	}
	return cl
}

//------------------HELPERS------------------------\\

//

//

//----------------------DEFAULT--------------------------\\

// New -> Creates totally a new DB Client, this is the root function!
func New(
	ctx context.Context,
	config dbDriver.Config,
) (*gorm.DB, error) {
	return dbClient.NewClient(ctx, config)
}

// GetDefaultClient -> this is a common function which helps getting the default Database client
// based on the current configuration!
func GetDefaultClient() (*gorm.DB, error) {
	// TODO: do we need a lock?!
	if DefaultInstanceClient != nil {
		return DefaultInstanceClient, nil
	}
	cl := getDefaultDBClient()
	// TODO: add function and variable which can override the get default client!...
	c, _err := cl.GetDefaultClient()
	if _err == nil {
		DefaultInstanceClient = c
	}
	return c, _err
}

// DB -> will get the default client without any errors
func DB() *gorm.DB {
	_db, _err := GetDefaultClient()
	if _err != nil {
		// todo: should we throw a panic?
		// 		 should we retry?!
		panic("failed to get default db client -> " + _err.Error())
	}
	return _db
}

func DBCtx(ctx context.Context) *gorm.DB {
	return DB().WithContext(ctx)
}

// NewClient ->  Get the default driver and generate a new client...
func NewClient(instanceName string) (*gorm.DB, error) {
	cl := getDefaultDBClient()
	return cl.NewClient(instanceName)
}

func GetClient(
	instanceName string,
) (*gorm.DB, error) {
	cl := getDefaultDBClient()
	return cl.GetClient(instanceName)
}

//----------------------DEFAULT--------------------------\\

//

//

//----------------------MYSQL--------------------------\\

// MySQLNew -> Creates totally a new DB Client, this is the root function!
func MySQLNew(
	ctx context.Context,
	config mysqlConfig.Config,
) (*gorm.DB, error) {
	return New(ctx, &config)
}

// MySQLGetDefaultClient -> this is a common function which helps getting the default Database client
// based on the current configuration!
func MySQLGetDefaultClient() (*gorm.DB, error) {
	if MySQLDefaultInstanceClient != nil {
		return MySQLDefaultInstanceClient, nil
	}
	cl := getDefaultDBClient()
	cl.DriverType = dbClient.MySQLDriver
	c, _err := cl.GetDefaultClient()
	if _err == nil {
		MySQLDefaultInstanceClient = c
	}
	return c, _err
}

// MySQLNewClient ->  Get the default driver and generate a new client...
func MySQLNewClient(instanceName string) (*gorm.DB, error) {
	cl := getDefaultDBClient()
	cl.DriverType = dbClient.MySQLDriver
	return cl.NewClient(instanceName)
}

func MySQLGetClient(
	instanceName string,
) (*gorm.DB, error) {
	cl := getDefaultDBClient()
	cl.DriverType = dbClient.MySQLDriver
	return cl.GetClient(instanceName)
}

//----------------------MYSQL--------------------------\\

//

//

//----------------------COCKROACHDB--------------------------\\

// CRDBNew -> Creates totally a new DB Client, this is the root function!
func CRDBNew(
	ctx context.Context,
	config cockroachConfig.Config,
) (*gorm.DB, error) {
	return New(ctx, &config)
}

// CRDBGetDefaultClient -> this is a common function which helps getting the default Database client
// based on the current configuration!
func CRDBGetDefaultClient() (*gorm.DB, error) {
	if CRDBDefaultInstanceClient != nil {
		return CRDBDefaultInstanceClient, nil
	}
	cl := getDefaultDBClient()
	cl.DriverType = dbClient.CockroachDriver
	c, _err := cl.GetDefaultClient()
	// Let's cache
	if _err == nil {
		CRDBDefaultInstanceClient = c
	}
	return c, _err
}

// CRDBNewClient ->  Get the default driver and generate a new client...
func CRDBNewClient(instanceName string) (*gorm.DB, error) {
	cl := getDefaultDBClient()
	cl.DriverType = dbClient.CockroachDriver
	return cl.NewClient(instanceName)
}

func CRDBGetClient(
	instanceName string,
) (*gorm.DB, error) {
	cl := getDefaultDBClient()
	cl.DriverType = dbClient.CockroachDriver
	return cl.GetClient(instanceName)
}

//----------------------COCKROACHDB--------------------------\\
