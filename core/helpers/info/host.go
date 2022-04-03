package info

import "github.com/shirou/gopsutil/host"

func GetHostInfo() *host.InfoStat {
	stat, _ := host.Info()
	return stat
}
