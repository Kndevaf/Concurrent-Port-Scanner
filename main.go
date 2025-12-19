package main

import "fmt"

func main() {
    host := "127.0.0.1"
    startPort := 1
    endPort := 1024

    fmt.Printf("Scanning %s...\n", host)
    ScanPorts(host, startPort, endPort)
}
