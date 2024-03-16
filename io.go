package golib

import (
    "io"
    "log"
    "net"
)

func HandleConn(localConn net.Conn, filePath, remoteAddr string) {
    defer localConn.Close()
    clientIP := localConn.RemoteAddr().(*net.TCPAddr).IP.String()
    if filePath != "" && !IsInFile(clientIP, filePath) {
        log.Printf("[WARN] %v", clientIP)
        return
    }
    remoteConn, err := net.Dial("tcp", remoteAddr)
    if err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
    defer remoteConn.Close()
    go io.Copy(remoteConn, localConn)
    io.Copy(localConn, remoteConn)
}