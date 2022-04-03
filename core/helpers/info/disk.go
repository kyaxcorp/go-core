package info

import (
	"github.com/fatih/structs"
	"github.com/shirou/gopsutil/disk"
)

func GetDisksInfo() map[string]interface{} {
	stat, _ := disk.Partitions(false)

	devices := make(map[string]interface{})
	for _, device := range stat {
		usage, _ := disk.Usage(device.Mountpoint)
		deviceMap := structs.Map(device)
		deviceMap["Usage"] = usage
		devices[device.Device] = deviceMap
	}

	return devices
}
