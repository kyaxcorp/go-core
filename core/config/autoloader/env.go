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
	// TODO: try adding the tags into the model, so the info will be more generalized...
	// 		create a function which will read the tags as a suffix!

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

		if logLevel := os.Getenv(getInstanceEnvVarName(connectionName, "CRDB_LOG_LEVEL")); logLevel != "" {
			dbInstance.Logger.Level = conv.StrToInt(logLevel)
		}

		cfgData.MainConfig.Clients.Cockroach.Instances[connectionName] = dbInstance
	}
	//---------------------------------------------------------------------------------\\
	//-----------------------\\    COCKROACHDB CLIENT    //----------------------------\\
	//---------------------------------------------------------------------------------\\

	//

	//---------------------------------------------------------------------------------\\
	//-----------------------\\    HTTP SERVER/LISTENER    //--------------------------\\
	//---------------------------------------------------------------------------------\\
	for instanceName, instance := range cfgData.MainConfig.Listeners.Http.Instances {
		// Comma separated
		if listeningAddresses := os.Getenv(getInstanceEnvVarName(instanceName, "HTTP_LISTENING_ADDRESSES")); listeningAddresses != "" {
			addresses := strings.Split(listeningAddresses, ",")
			instance.ListeningAddresses = addresses
		}
		if listeningSSLAddresses := os.Getenv(getInstanceEnvVarName(instanceName, "HTTP_LISTENING_SSL_ADDRESSES")); listeningSSLAddresses != "" {
			addresses := strings.Split(listeningSSLAddresses, ",")
			instance.ListeningAddressesSSL = addresses
		}

		cfgData.MainConfig.Listeners.Http.Instances[instanceName] = instance
	}
	//---------------------------------------------------------------------------------\\
	//-----------------------\\    HTTP SERVER/LISTENER    //--------------------------\\
	//---------------------------------------------------------------------------------\\

	return _err
}
