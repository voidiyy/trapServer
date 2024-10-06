package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func Server(port, host, certFile, keyFile string) {

	if port == "" {
		port = "1234"
	}
	if host == "" {
		host = "localhost"
	}

	mux := Handle()

	if certFile == "" && keyFile == "" {
		fmt.Println("Listening HTTP on " + host + ":" + port)
		err := httpSrv(port, mux)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Listening HTTPS on port %s\n", port)

	err := httpsSrv(port, host, certFile, keyFile, mux)
	if err != nil {
		fmt.Println(err)
	}
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
