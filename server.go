package golib

import (
    "crypto/tls"
    "net"
    "net/http"
)

func ServeHTTP(hostname, port string, handler http.Handler, tlsConfig *tls.Config) error {
    server := &http.Server{
        Addr:      net.JoinHostPort(hostname, port),
        Handler:   handler,
        TLSConfig: tlsConfig,
    }
    if tlsConfig == nil {
        return server.ListenAndServe()
    } else {
        return server.ListenAndServeTLS("", "")
    }
}