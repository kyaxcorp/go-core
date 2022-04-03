package info

import (
	"github.com/KyaXTeam/go-core/v2/core/bootstrap"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"runtime"
	"time"
)

// TODO: should be moved in helpers or smth
type SystemStatus struct {
	PID                   int
	ExecutableLocation    string
	UserID                int
	GroupID               int
	Started               time.Time
	RunningTimeSeconds    float64
	NrOfRunningGoroutines int
	MemStats              map[string]interface{}
	NrOfCPUUsed           int
	NrOfCGOCalls          int64
	CPUInfo               CPUInfoStr

	DisksInfo         map[string]interface{}
	NetInfo           map[string]interface{}
	HostInfo          *host.InfoStat
	VirtualMemoryStat *mem.VirtualMemoryStat
}

func GetSystemStatus() SystemStatus {
	b := bootstrap.GetProcessBootstrap()

	executablePath, _ := os.Executable()

	cpuInfoChan := make(chan CPUInfoStr)
	processMemInfoChan := make(chan map[string]interface{})
	memInfoChan := make(chan *mem.VirtualMemoryStat)
	hostInfoChan := make(chan *host.InfoStat)
	disksInfoChan := make(chan map[string]interface{})
	netInfoChan := make(chan map[string]interface{})

	go func() {
		cpuInfoChan <- GetCPUInfo()
	}()
	go func() {
		processMemInfoChan <- GetProcessMemoryInfo()
	}()
	go func() {
		memInfoChan <- GetMemoryInfo()
	}()
	go func() {
		hostInfoChan <- GetHostInfo()
	}()
	go func() {
		disksInfoChan <- GetDisksInfo()
	}()
	go func() {
		netInfoChan <- GetNetInfo()
	}()

	return SystemStatus{
		PID:                   os.Getpid(),
		ExecutableLocation:    executablePath,
		UserID:                os.Getuid(),
		GroupID:               os.Getgid(),
		Started:               b.GetStartTime(),
		RunningTimeSeconds:    b.GetRunningTime().Seconds(),
		NrOfRunningGoroutines: runtime.NumGoroutine(),
		NrOfCPUUsed:           runtime.NumCPU(),
		NrOfCGOCalls:          runtime.NumCgoCall(),

		MemStats:          <-processMemInfoChan,
		CPUInfo:           <-cpuInfoChan,
		VirtualMemoryStat: <-memInfoChan,
		HostInfo:          <-hostInfoChan,
		DisksInfo:         <-disksInfoChan,
		NetInfo:           <-netInfoChan,
	}
}
