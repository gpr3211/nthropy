package main

import (
	"fmt"

	"github.com/gpr3211/nthropy/metrics"
	"github.com/shirou/gopsutil/v3/host"
)

func main() {
	info, _ := host.Info()
	fmt.Printf("Hostname: %s, OS: %s, Platform: %s %s\n",
		info.Hostname, info.OS, info.Platform, info.PlatformVersion)
	metrics.GetUptime()
	metrics.PrintMetrics()
}
