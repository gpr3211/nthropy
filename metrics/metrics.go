package metrics

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

type CpuData struct{}
type MemData struct{}

type Metrics struct {
	cpu    CpuData
	memory MemData
}

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	// Get per-core usage over 1 second
	percents, err := cpu.Percent(time.Second, true)
	if err != nil {
		panic(err)
	}

	for i, percent := range percents {
		fmt.Printf("Core %d: %.2f%%\n", i, percent)
	}
}
