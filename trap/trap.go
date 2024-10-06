package trap

import "trapServer/trap/utils"

type Client struct {
	Fingerprint *utils.FingerprintData
	IPAddr      *utils.IpAddress
	UserAgent   *utils.UserAgent
}

type VPN struct {
	IsEnabled bool
	Provider  string
	Country   string
}

type Proxy struct {
	IsEnabled bool
	Forwarder string
}
