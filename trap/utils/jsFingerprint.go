package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type deviceInfo struct {
	Timezone            string `json:"timezone"`
	MonitorResolution   string `json:"monitorResolution"`
	Language            string `json:"language"`
	HardwareConcurrency int    `json:"hardwareConcurrency"`
	ViewportSize        string `json:"viewportSize"`
	DeviceMemory        string `json:"deviceMemory"`
	Platform            string `json:"platform"`
	CpuClass            string `json:"cpuClass"`
}

type FingerprintData struct {
	DeviceInfo deviceInfo `json:"deviceInfo"`
}

func JSReader(r *http.Request) string {
	var data FingerprintData

	rbody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(rbody, &data)
	if err != nil {
		fmt.Println(err)
	}

	str := fmt.Sprintf(`Received Fingerprint Data:
==========================
Timezone: %s
Monitor Resolution: %s
Language: %s
Hardware Concurrency: %d
Viewport Size: %s
Device Memory: %s GB
Platform: %s
CPU Class: %s
`, data.DeviceInfo.Timezone, data.DeviceInfo.MonitorResolution, data.DeviceInfo.Language, data.DeviceInfo.HardwareConcurrency, data.DeviceInfo.ViewportSize, data.DeviceInfo.DeviceMemory, data.DeviceInfo.Platform, data.DeviceInfo.CpuClass)

	return str

}
