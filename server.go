package golib

import (
    "crypto/tls"
    "net"
    "net/http"
)

func ServeHTTP(hostname, port string, handler http.Handler) error {
    return http.ListenAndServe(net.JoinHostPort(hostname, port), handler)
}

func ServeHTTPS(hostname, port string, handler http.Handler, tlsConfig *tls.Config) error {
    server := &http.Server{
        Addr:      net.JoinHostPort(hostname, port),
        Handler:   handler,
        TLSConfig: tlsConfig,
    }
    return server.ListenAndServeTLS("", "")
}