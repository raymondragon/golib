package golib

import (
    "crypto/tls"
    "net/http"
)

func ServeHTTP(host string, handler http.Handler) error {
    return http.ListenAndServe(host, handler)
}

func ServeHTTPS(host string, handler http.Handler, tlsConfig *tls.Config) error {
    server := &http.Server{
        Addr:      host,
        Handler:   handler,
        TLSConfig: tlsConfig,
    }
    return server.ListenAndServeTLS("", "")
}