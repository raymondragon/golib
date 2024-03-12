package golib

import (
    "crypto/tls"
    "net/http"
)

func ServeHTTP(parsedURL ParsedURL, handler http.Handler, tlsConfig *tls.Config) error {
    server := &http.Server{
        Addr:      parsedURL.Hostname + ":" + parsedURL.Port,
        Handler:   handler,
        TLSConfig: tlsConfig,
    }
    if tlsConfig == nil {
        return server.ListenAndServe()
    } else {
        return server.ListenAndServeTLS("", "")
    }
}
