package constructor

import (
	"context"
	"github.com/KyaXTeam/go-core/v2/core/clients/db/dbinstance"
	"sync"
)

const MySQLDriver = "mysql"
const CockroachDriver = "cockroach"
const SQLiteDriver = "sqlite"
const PostGRESDriver = "postgres"

// Here we store the mysql instances

var driverInstancesLock sync.RWMutex
var driverInstances = make(map[string]*dbinstance.Instance)

// var instances = dbinstance.NewInstance()

type DBClient struct {
	DriverType  string
	instanceRef *dbinstance.Instance
	Ctx         context.Context
}
