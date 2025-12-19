# Concurrent Port Scanner (Go)

A fast, concurrent **TCP port scanner written in Go (Golang)**. This tool demonstrates Go’s powerful concurrency model using **goroutines and wait groups**, while applying real-world **networking and cybersecurity concepts**.

This project is designed for **learning, portfolio building, and interview discussion**.

---

## Features

* Concurrent TCP port scanning
* Customizable target host and port range
* Fast scanning using goroutines
* Timeout handling to avoid hangs
* Simple, clean CLI output
* Cross-platform (macOS, Linux, Windows)

---

## Concepts Demonstrated

* Go concurrency (goroutines & `sync.WaitGroup`)
* Networking with the `net` package
* TCP connection handling
* Timeouts and performance tuning
* Modular Go project structure

---

## Project Structure

```
portscanner/
│── main.go       # Entry point
│── scanner.go    # Port scanning logic
│── README.md
```

---

## Requirements

* Go 1.20+ (recommended)

Check your version:

```bash
go version
```

---

## Build & Run

### Run directly

```bash
go run main.go scanner.go
```

### Build executable

```bash
go build -o portscanner
./portscanner
```

---

## Example Output

```
Scanning 127.0.0.1...
Port 22 is open
Port 80 is open
Scan complete.
```

---

## 🔍 How It Works

* Each port scan runs in its own **goroutine**
* `net.DialTimeout` attempts a TCP connection
* Open ports successfully establish a connection
* A `WaitGroup` ensures all scans finish before exit

---

## Disclaimer

This tool is for **educational and authorized testing only**. Do not scan systems you do not own or have explicit permission to test.

---

## Future Improvements

* CLI flags for host and port range
* Adjustable concurrency limits
* UDP scanning support
* Colored output
* Export results to file (JSON/CSV)
* Multi-host scanning

---

## Author

**F.K Madagu**
Cybersecurity & IT | Go (Golang) | Network Security Tools

---

