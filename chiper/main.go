package main

import (
    "crypto/tls"
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
        w.Write([]byte("This is an example server.\n"))
    })
    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
	MaxVersion:		  tls.VersionTLS12,
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
	    tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
        },
    }
    srv := &http.Server{
        Addr:         ":8443",
        Handler:      mux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("../certs/mtls.furikuri.net.crt", "../certs/mtls.furikuri.net.key"))
}
