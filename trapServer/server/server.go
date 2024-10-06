package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func Server(port, host, certFile, keyFile string) error {

	if port == "" {
		port = "1234"
	}
	if host == "" {
		host = "localhost"
	}

	mux := Handle()

	if certFile == "" && keyFile == "" {
		fmt.Println("Listening HTTP on " + host + ":" + port)
		return httpSrv(port, mux)
	}

	fmt.Printf("Listening HTTPS on port %s\n", port)

	return httpsSrv(port, host, certFile, keyFile, mux)
}

func httpSrv(port string, mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func httpsSrv(port, host, certFile, keyFile string, mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequestClientCert,
			ServerName: host,
		},
	}

	return srv.ListenAndServeTLS(certFile, keyFile)
}
