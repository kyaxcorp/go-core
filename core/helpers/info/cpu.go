package info

import cpustat "github.com/shirou/gopsutil/cpu"
import . "github.com/klauspost/cpuid/v2"

type CPUInfoStr struct {
	General  CPUInfo
	Detailed []cpustat.InfoStat
}

func GetCPUInfo() CPUInfoStr {
	//stat, _ := cpustat.Info()

	return CPUInfoStr{
		General: CPU,
		// The stat is slow...
		//Detailed: stat,
	}
}
