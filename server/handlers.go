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

	//render html with js inside
	tmpl, err := template.ParseFiles("server/canvas.html")
	if err != nil {
		fmt.Printf("trap template parse error: %v\n", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Printf("trap template execute error: %v\n", err)
	}

	//headers
	interestingHeaders := utils.CheckHeaders(r)

	//user agent and ip
	ua := utils.ParseUserAgent(r.Header.Get("User-Agent"))
	ip := utils.ParseIP(r)

	ipInfo, err := utils.IPInfo(ip.Value)
	if err != nil {
		fmt.Printf("trap ip parse error: %v\n", err)
	}

	fmt.Println(ipInfo.String(ip.Value))
	fmt.Println(ua.String())
	fmt.Println(interestingHeaders)

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
Hardware Concurrency: %d
Viewport Size: %s
Device Memory: %s GB
Platform: %s
CPU Class: %s
`, data.DeviceInfo.Timezone, data.DeviceInfo.MonitorResolution, data.DeviceInfo.Language, data.DeviceInfo.HardwareConcurrency, data.DeviceInfo.ViewportSize, data.DeviceInfo.DeviceMemory, data.DeviceInfo.Platform, data.DeviceInfo.CpuClass)

	fmt.Printf("%s\n", str)
}
