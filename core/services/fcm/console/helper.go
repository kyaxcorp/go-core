package console

import (
	"github.com/kyaxcorp/go-core/core/helpers/array"
	"log"
	"strings"
)

func getInstanceName(args []string) string {

	var instanceName string
	if array.IndexExistsString(0, args) {
		instanceName = args[0]
	}

	if instanceName == "" {
		log.Println("taking default configuration instance name")
		//instanceName = config.GetConfig().Clients.Broker.DefaultInstanceName
	}

	if instanceName == "" {
		// If still empty... raise an error!
		log.Fatal("instance name is empty...")
	}

	return strings.ToLower(instanceName)
}
