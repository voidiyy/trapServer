package main

import (
	"flag"
	"fmt"
	"trapServer/server"
)

func main() {
	help := flag.Bool("help", false, "Optional, prints usage info")
	host := flag.String("host", "", "Required flag, must be the hostname that is resolvable via DNS, or 'localhost'")
	port := flag.String("port", "443", "The https port, defaults to 443")
	certFile := flag.String("cert", "", "Required, the name of the server's certificate file")
	keyFile := flag.String("key", "", "Required, the file name of the server's private key file")
	flag.Parse()

	usage := `usage:
	
trapServer -host <hostname> -cert <serverCertFile> -key <serverPrivateKeyFile> [-port <port> -help]
	
Options:
  -help       Prints this message
  -host       Required, a DNS resolvable host name or 'localhost'
  -cert    Required, the name the server's certificate file
  -key     Required, the name the server's key certificate file
  -port       Optional, the https port for the server to listen on, defaults to 443

   if -cert -key not provided, server run on HTTP
  `

	if *help == true {
		fmt.Println(usage)
		return
	}

	server.Server(*port, *host, *certFile, *keyFile)
}
