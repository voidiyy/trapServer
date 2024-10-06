# trapServer
collect info about client
-ip info
-analythe user-agent 
-check headers
-grab device data via JS

#Usage:

trapServer -host <hostname> -port <portnumber> -cert <serverCertFile> -key <serverPrivateKeyFile>
	
Options:
  -help       Prints this message
  -host       Required, a DNS resolvable host name or 'localhost'
  -cert    Required, the name the server's certificate file
  -key     Required, the name the server's key certificate file
  -port       Optional, the https port for the server to listen on, defaults to 443

  if cert & key files not specified, run as HTTP


  Result example:

  (VPN moment):
  193.19.207.217 Country: Spain (ES)
  Timezone: Europe/Kyiv


  IP Information: 193.19.207.217
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

User Agent Information:
Raw: Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0
OS: Linux (Version: 115.0)
Browser: Firefox (Version: 115.0)
Device:  (Type: Desktop)

Received Fingerprint Data:
==========================
Timezone: Europe/Kyiv
Monitor Resolution: 1920x1080
Language: en-US
Hardware Concurrency: 8
Viewport Size: 1920x995
Device Memory: N/A GB
Platform: Linux x86_64
CPU Class: N/A
