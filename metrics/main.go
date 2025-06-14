package metrics

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
)

type CPUInfo struct {
	Core    int     `json:"core"`
	Percent float64 `json:"percent"`
}

type TempInfo struct {
	SensorKey   string  `json:"sensor_key"`
	Temperature float64 `json:"temperature"`
	High        float64 `json:"high,omitempty"`
	Critical    float64 `json:"critical,omitempty"`
}

type DiagnosticData struct {
	Timestamp    time.Time  `json:"timestamp"`
	CPUUsage     []CPUInfo  `json:"cpu_usage"`
	Temperatures []TempInfo `json:"temperatures"`
}

func PrintMetrics() {

	cpu, err := collectCPUUsage()
	if err != nil {

		fmt.Println("failed to collect cpu usage")
		log.Fatalln(err)
	}
	temps, err := collectTemperatures()
	if err != nil {
		fmt.Println("failed to collect temps")
		log.Fatalln(err)
	}

	data := createTempTable(cpu, temps)
	for k, v := range data {
		fmt.Printf("Core: %s Temp: %f \n", k, v)
	}

}

func collectCPUUsage() ([]CPUInfo, error) {
	// Get per-core usage over 1 second
	percents, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}

	var cpuInfo []CPUInfo
	for i, percent := range percents {
		cpuInfo = append(cpuInfo, CPUInfo{
			Core:    i,
			Percent: percent,
		})
	}
	return cpuInfo, nil
}

func collectTemperatures() ([]TempInfo, error) {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		//	return nil, err
	}

	var tempInfo []TempInfo
	for _, temp := range temps {
		tempInfo = append(tempInfo, TempInfo{
			SensorKey:   temp.SensorKey,
			Temperature: temp.Temperature,
			High:        temp.High,
			Critical:    temp.Critical,
		})
	}
	return tempInfo, nil
}
func createTempTable(cpuUsage []CPUInfo, temperatures []TempInfo) map[string]float64 {

	// Create temperature map for easy lookup
	tempMap := make(map[string]float64)
	for _, temp := range temperatures {
		// Try different key patterns since sensor keys vary by system
		tempMap[temp.SensorKey] = temp.Temperature
		tempMap[fmt.Sprintf("Core %d", len(tempMap))] = temp.Temperature
	}
	// Temperature column

	return tempMap
}

func GetUptime() {
	uptime, _ := host.Uptime()
	fmt.Printf("Uptime: %s\n", time.Duration(uptime)*time.Second)

	bootTime, _ := host.BootTime()
	fmt.Println("Boot time:", time.Unix(int64(bootTime), 0))

}
