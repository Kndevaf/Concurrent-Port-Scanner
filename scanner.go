package main

import (
    "fmt"
    "net"
    "sync"
    "time"
)

func ScanPorts(host string, start, end int) {
    var wg sync.WaitGroup

    for port := start; port <= end; port++ {
        wg.Add(1)
        go func(p int) {
            defer wg.Done()
            address := fmt.Sprintf("%s:%d", host, p)
            conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
            if err != nil {
                return
            }
            conn.Close()
            fmt.Printf("Port %d is open\n", p)
        }(port)
    }

    wg.Wait()
    fmt.Println("Scan complete.")
}
