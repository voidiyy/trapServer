package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"trapServer/trap/utils"
)

func Handle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/trap/", trp)
	mux.HandleFunc("/fingerprint", fingerprint)

	return mux
}

func root(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("root"))
	if err != nil {
		fmt.Printf("trap write error: %v\n", err)
	}
}

func trp(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("server/canvas.html")
	if err != nil {
		fmt.Printf("trap template parse error: %v\n", err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("trap template execute error: %v\n", err)
	}

	ua := utils.ParseUserAgent(r.Header.Get("User-Agent"))
	ip := utils.ParseIP(r)

	str := fmt.Sprintf(
		"IP address: %v\n User Agent: %s\nOS: %s %s\nBrowser: %s %s\nDevice: %s (%s)",
		ip.Value, ua.Raw,
		ua.OS.Name, ua.OS.Version,
		ua.Browser.Name, ua.Browser.Version,
		ua.Device.Name, ua.Device.Type,
	)

	fmt.Println(str)
}

func fingerprint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	var data utils.FingerprintData

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
Location: %s
Hardware Concurrency: %d
Viewport Size: %s
Device Memory: %s GB
Platform: %s
CPU Class: %s
`, data.DeviceInfo.Timezone, data.DeviceInfo.MonitorResolution, data.DeviceInfo.Language, data.DeviceInfo.Location, data.DeviceInfo.HardwareConcurrency, data.DeviceInfo.ViewportSize, data.DeviceInfo.DeviceMemory, data.DeviceInfo.Platform, data.DeviceInfo.CpuClass)

	fmt.Printf("%s\n", str)
}
