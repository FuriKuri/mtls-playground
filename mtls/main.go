package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/", helloHandler)

	caCert, err := ioutil.ReadFile("../certs/rootCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	log.Fatal(server.ListenAndServeTLS("../certs/mtls.furikuri.net.crt", "../certs/mtls.furikuri.net.key"))
}