package golib

import (
    "bufio"
    "os"
    "strings"
)

func IsInFile(str, filepath string) bool {
    file, err := os.Open(filepath)
    if err != nil {
        return false
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if strings.TrimSpace(scanner.Text()) == str {
            return true
        }
    }
    if err := scanner.Err(); err != nil {
        return false
    }
    return false
}