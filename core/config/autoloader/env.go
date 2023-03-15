package autoloader

import (
	cfgData "github.com/kyaxcorp/go-core/core/config/data"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"os"
	"strings"
)

func getInstanceEnvName(instanceName string) string {
	return strings.ReplaceAll(strings.ToUpper(instanceName), "-", "_")
}

func getInstanceEnvVarName(instanceName, varName string) string {
	return getInstanceEnvName(instanceName) + "_" + varName
}

func setEnv() error {
	var _err error
	//---------------------------------------------------------------------------------\\
	//-----------------------\\    COCKROACHDB CLIENT    //----------------------------\\
	//---------------------------------------------------------------------------------\\
	for connectionName, dbInstance := range cfgData.MainConfig.Clients.Cockroach.Instances {

		if host := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_HOST")); host != "" {
			dbInstance.Credentials.Host = host
		}

		if port := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_PORT")); port != "" {
			dbInstance.Credentials.Port = conv.StrToInt(port)
		}

		if user := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_USERNAME")); user != "" {
			dbInstance.Credentials.User = user
		}

		if password := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_PASSWORD")); password != "" {
			dbInstance.Credentials.Password = password
		}

		if dbName := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_DB_NAME")); dbName != "" {
			dbInstance.Credentials.DbName = dbName
		}

		if schema := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_SCHEMA")); schema != "" {
			dbInstance.Credentials.Schema = schema
		}

		if sslMode := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_SSL_MODE")); sslMode != "" {
			dbInstance.Credentials.SSLMode = sslMode
		}

		cfgData.MainConfig.Clients.Cockroach.Instances[connectionName] = dbInstance
	}
	//---------------------------------------------------------------------------------\\
	//-----------------------\\    COCKROACHDB CLIENT    //----------------------------\\
	//---------------------------------------------------------------------------------\\

	return _err
}
