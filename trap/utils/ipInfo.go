package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type IPInformation struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Reverse     string  `json:"reverse"`
	Mobile      bool    `json:"mobile"`
	Proxy       bool    `json:"proxy"`
	Hosting     bool    `json:"hosting"`
}

func IPInfo(ip string) (*IPInformation, error) {

	info := &IPInformation{}

	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,countryCode,region,regionName,city,zip,lat,lon,timezone,isp,org,as,reverse,mobile,proxy,hosting", ip)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}

	return info, nil

}

func (ip IPInformation) String(ipAddr string) string {
	return fmt.Sprintf(
		"IP Information: %s\n"+
			"Status: %s\nCountry: %s (%s)\nRegion: %s, %s\nCity: %s\nZip: %s\n"+
			"Coordinates: %.2f, %.2f\nTimezone: %s\nISP: %s\nOrganization: %s\n"+
			"AS: %s\nReverse DNS: %s\nMobile: %t\nProxy/VPN: %t\nHosting: %t\n",
		ipAddr, ip.Status, ip.Country, ip.CountryCode, ip.Region, ip.RegionName, ip.City, ip.Zip,
		ip.Lat, ip.Lon, ip.Timezone, ip.Isp, ip.Org, ip.As, ip.Reverse, ip.Mobile, ip.Proxy, ip.Hosting,
	)
}
