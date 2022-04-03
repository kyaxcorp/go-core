package info

import (
	"github.com/fatih/structs"
	"github.com/shirou/gopsutil/mem"
	"runtime"
)

func GetMemoryInfo() *mem.VirtualMemoryStat {
	vmStat, _ := mem.VirtualMemory()
	return vmStat
}

func GetProcessMemoryInfo() map[string]interface{} {
	bToMb := func(b uint64) uint64 {
		return b / 1024 / 1024
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memStats := structs.Map(m)
	delete(memStats, "PauseNs")
	delete(memStats, "PauseEnd")
	delete(memStats, "BySize")
	memStats["AllocMB"] = bToMb(memStats["Alloc"].(uint64))
	memStats["TotalAllocMB"] = bToMb(memStats["TotalAlloc"].(uint64))
	memStats["SysMB"] = bToMb(memStats["Sys"].(uint64))
	return memStats
}
