package golib

import (
    "bufio"
    "os"
    "strings"
)

func IsInFile(strName, filePath string) bool {
    file, err := os.Open(filePath)
    if err != nil {
        return false
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.TrimSpace(scanner.Text()) == strName {
            return true
        }
    }
    if err := scanner.Err(); err != nil {
        return false
    }
    return false
}