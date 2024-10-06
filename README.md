# trapServer

Collects information about a client:
- **IP info**
- **Analyze the User-Agent**
- **Check headers**
- **Grab device data via JavaScript**

## Usage

```bash
trapServer -host <hostname> -port <portnumber> -cert <serverCertFile> -key <serverPrivateKeyFile>
```

-help Prints this message.
-host Required, a DNS-resolvable host name or 'localhost'.
-cert Required, the server's certificate file.
-key Required, the server's private key file.
-port Optional, the HTTPS port for the server to listen on, defaults to 443.

# Result example:

VPN moment:
193.19.207.217 Country: Spain (ES)
Timezone: Europe/Kyiv

IP: 193.19.207.217
Status: success
Country: Spain (ES)
Region: VC, Valencia
City: Valencia
Zip: 46002
Coordinates: 39.49, -0.40
Timezone: Europe/Madrid
ISP: HostRoyale Technologies Pvt Ltd
Organization: 
AS: AS203020 HostRoyale Technologies Pvt Ltd
Reverse DNS: 
Mobile: false
Proxy/VPN: true
Hosting: false

Raw: Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0
OS: Linux (Version: 115.0)
Browser: Firefox (Version: 115.0)
Device:  (Type: Desktop)

Timezone: Europe/Kyiv
Monitor Resolution: 1920x1080
Language: en-US
Hardware Concurrency: 8
Viewport Size: 1920x995
Device Memory: N/A GB
Platform: Linux x86_64
CPU Class: N/A

