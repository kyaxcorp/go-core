package info

import "github.com/fatih/structs"
import "github.com/shirou/gopsutil/net"

func GetNetInfo() map[string]interface{} {
	interfaces, _ := net.Interfaces()
	interfacesMap := make(map[string]interface{})
	for _, interfaceStat := range interfaces {
		intMap := structs.Map(interfaceStat)
		interfacesMap[interfaceStat.Name] = intMap
	}

	return interfacesMap
}
