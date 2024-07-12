package golib

import (
    "log"
    "net"
    "net/http"
    "os"

    "golang.org/x/net/webdav"
    "github.com/elazarl/goproxy"
    "github.com/elazarl/goproxy/ext/auth"
)

func WebdavHandler(dir, prefix string) http.Handler {
    return &webdav.Handler{
        FileSystem: webdav.Dir(dir),
        Prefix:     prefix,
        LockSystem: webdav.NewMemLS(),
    }
}

func ProxyHandler(hostname, username, password string, custom http.Handler) http.Handler {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    auth.ProxyBasic(proxy, hostname, func(uname, passwd string) bool {
        return uname == username && passwd == password
    })
    if custom != nil {
        proxy.NonproxyHandler = custom
    } else {
        proxy.NonproxyHandler = http.HandlerFunc(IPDisplayHandler)
    }
    return proxy
}

func IPDisplayHandler(w http.ResponseWriter, r *http.Request) {
    clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        log.Printf("[WARN] %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if _, err := w.Write([]byte(clientIP + "\n")); err != nil {
        log.Printf("[WARN] %v", err)
        return
    }
}

func IPRecordHandler(fileName string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            log.Printf("[WARN] %v", err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
            log.Printf("[WARN] %v", err)
            return
        }
        defer file.Close()
        if IsInFile(clientIP, fileName) {
            return
        }
        if _, err := file.WriteString(clientIP + "\n"); err != nil {
            log.Printf("[WARN] %v", err)
            return
        }
    }
}
