package main

import (
	"fmt"
	"github.com/kyaxcorp/go-core/core/bootstrap"
	"github.com/kyaxcorp/go-core/core/clients/db"
	"github.com/kyaxcorp/go-core/core/clients/db/dbresolver"
	config "github.com/kyaxcorp/go-core/core/clients/db/driver/mysql/config"
	configLoader "github.com/kyaxcorp/go-core/core/config/autoloader"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/errors2"
	loggerConfig "github.com/kyaxcorp/go-core/core/logger/config"
	"log"
	"time"
)

/*
1. Create the client
*/

type test struct {
	ID   int
	Name string
}

var Data ConfigModel

type ConfigModel struct {
}

type DBExampleTable struct {
	ID        int
	Name      string
	FirstName string
}

func LoadConfigs() bool {
	//if !config.AutoLoad() {
	if _err := configLoader.StartAutoLoader(configLoader.Config{
		CustomConfig:              &Data,
		CustomConfigModel:         &ConfigModel{},
		AdditionalLoggingChannels: nil,
	}); _err != nil {
		msg := "Failed to load Main configuration file!" + " -> " + _err.Error()
		errors2.New(0, msg)
		fmt.Println(msg)
		return false
	}
	return true
}

func main() {
	// Define version
	//version.Init(version.AppVersion{
	//	ProjectName: "db-resolver",
	//	Version:     "1.0.0",
	//	BuildBy:     "MrDiffer",
	//})

	// Define error reporting and capture panics
	//error_reporting.New("https://fc46abe043fb4273b67437eba3ef909c@error.kyax.de/2", false)
	//defer error_reporting.BeforeShutdown()
	// Bootstrap
	bootstrap.StartForProcess()
	// Load configs
	if !LoadConfigs() {
		return
	}

	// that's where we will receive the data...
	var tests []test

	port := 3310
	// Create the on connect options
	onConnectOptions := config.OnConnectOptions{}
	_struct.SetDefaultValues(&onConnectOptions)

	// Create the connections
	mainResolverSource1 := config.Connection{
		CredentialsOverrides: config.CredentialsOverrides{
			Host: "yes",
			Port: "yes",
		},
		Credentials: config.Credentials{
			Port: port,
			Host: "localhost",
		},
	}
	mainResolverSource2 := config.Connection{
		CredentialsOverrides: config.CredentialsOverrides{
			Host: "yes",
			Port: "yes",
		},
		Credentials: config.Credentials{
			Port: port,
			Host: "localhost",
		},
	}

	// Set the Connections defaults
	_struct.SetDefaultValues(&mainResolverSource1)
	_struct.SetDefaultValues(&mainResolverSource2)

	/*policy1, _err := dbresolver.NewRoundRobinPolicy()
	if _err != nil {
		log.Println("POLICY OPTIONS ERRRRRORRR 11")
	}

	policy2, _err := dbresolver.NewRoundRobinPolicy()
	if _err != nil {
		log.Println("POLICY OPTIONS ERRRRRORRR 22")
	}*/

	// create the first resolver
	mainResolver := config.Resolver{
		Sources:    []config.Connection{mainResolverSource1, mainResolverSource2},
		Replicas:   nil,
		PolicyName: dbresolver.TRoundRobin,
		//PolicyOptions: policy1,
	}

	// set default values for  the Main Resolver
	_struct.SetDefaultValues(&mainResolver)

	// create the second resolver

	secondResolver := config.Resolver{
		Sources:    []config.Connection{mainResolverSource1, mainResolverSource2},
		Replicas:   nil,
		PolicyName: dbresolver.TRoundRobin,
		//PolicyOptions: policy2,
	}
	// set default values for the second resolver
	_struct.SetDefaultValues(&secondResolver)
	// aggregate resolvers
	//resolvers := []config.Resolver{mainResolver, secondResolver}
	//resolvers := []config.Resolver{mainResolver, secondResolver}

	resolvers := []config.Resolver{
		mainResolver,
		secondResolver,
	}

	// Create the Log config
	logConfig, _err := loggerConfig.DefaultConfig(&loggerConfig.Config{
		Name:      "dbresolver",
		IsEnabled: "yes",
	})

	if _err != nil {
		log.Println("failed to set settings", _err)
	}

	// create the main config db config
	_config := config.Config{
		IsEnabled:   "yes",
		Description: "",
		// These params are global...they are referred as to a cluster
		Credentials: config.Credentials{
			User:     "root",
			Password: "1w3r5y13245!@!!",
			DbName:   "gocore",
		},
		OnConnectOptions: onConnectOptions,
		Resolvers:        resolvers,
		Logger:           logConfig,
	}

	// set default values only for main config level
	_struct.SetDefaultValues(&_config)

	// create the client
	//dbClient, _err := mysql.New(_context.GetDefaultContext(), &_config)

	// YOu can use this case
	//dbClient, _err := db.New(_context.GetDefaultContext(), &_config)
	// Or use this case... same thing...
	dbClient, _err := db.MySQLNew(_context.GetDefaultContext(), _config)
	if _err != nil {
		log.Println("error", _err)
		return
	}

	// in this case we will use same client by multiple goroutines
	execDB := func(goroutineID string) {
		result := dbClient.Find(&tests)

		log.Println("")
		log.Println("")
		log.Println("")
		log.Println("")
		log.Println(goroutineID, "rows affected", result.RowsAffected)
		log.Println(goroutineID, "result error", result.Error)
		log.Println(goroutineID, "data", tests)
		log.Println("")
		log.Println("")
		log.Println("")
		log.Println("")
		newTest := &test{
			Name: "random name",
		}
		result = dbClient.Create(newTest)
		log.Println(goroutineID, "create rows affected", result.RowsAffected)
		log.Println(goroutineID, "create result error", result.Error)
		log.Println(goroutineID, "create data", newTest)

		/*result = dbClient.Table("test").Find(&tests)
		log.Println(goroutineID, "rows affected", result.RowsAffected)
		log.Println(goroutineID, "result error", result.Error)
		log.Println(goroutineID, "data", tests)

		result = dbClient.Table("test").Find(&tests)
		log.Println(goroutineID, "rows affected", result.RowsAffected)
		log.Println(goroutineID, "result error", result.Error)
		log.Println(goroutineID, "data", tests)*/
	}

	/*	go execDB("1")
		go execDB("2")*/

	dbClient.AutoMigrate(&DBExampleTable{})

	startTime := time.Now().Nanosecond()
	dbClient.Exec("TRUNCATE tests")
	endTime := time.Now().Nanosecond()
	log.Println("truncate time = ", endTime-startTime)
	log.Println("truncated...")

	for {
		execDB("3")
		log.Println("FINISHED....")
		time.Sleep(time.Second * 5)
		//go execDB("3")

	}

	// print the data
}
