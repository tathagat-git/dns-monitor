# DNS & HTTP Health Monitor

A lightweight CLI tool written in Go that performs real-time DNS resolution and HTTP health checks on a list of domains. Built to demonstrate network monitoring fundamentals relevant to Zero Trust and edge security systems.

## What it does

- Resolves domain → IP addresses via DNS lookup
- Makes HTTPS requests and measures response latency
- Reports HTTP status codes and flags non-2xx responses
- Handles timeouts and network errors gracefully

## Why it's relevant

Zero Trust security platforms like Akamai's Enterprise Threat Protector monitor DNS queries and HTTP traffic to detect threats like DNS exfiltration, phishing, and ransomware C2 communication. This tool demonstrates the same underlying primitives: DNS resolution, HTTP inspection, and latency measurement.

## Tech

- **Language:** Go 1.22
- **Packages used:** `net` (DNS), `net/http` (HTTP client), `time`, `os` — all standard library, no external dependencies

## Run

```bash
# Install Go: https://go.dev/dl/

# Clone and run
git clone https://github.com/tathagat-git/dns-monitor
cd dns-monitor

# Default domains (google.com, github.com, cloudflare.com)
go run main.go

# Custom domains
go run main.go akamai.com example.com
```

## Sample Output

```
Checking 3 domain(s)...
  → google.com
  → github.com
  → cloudflare.com

========================================
   DNS & HTTP Health Monitor Report
========================================

Domain  : google.com
IPs     : [142.250.67.46 2404:6800:4009:821::200e]
HTTP    : 200
Latency : 312ms
Health  : ✅ OK

Domain  : github.com
IPs     : [140.82.112.4]
HTTP    : 200
Latency : 198ms
Health  : ✅ OK

Domain  : cloudflare.com
IPs     : [104.16.132.229 104.16.133.229]
HTTP    : 200
Latency : 89ms
Health  : ✅ OK

========================================
```

## Concepts demonstrated

| Concept | Go code |
|---|---|
| DNS resolution | `net.LookupHost()` |
| HTTP client with timeout | `http.Client{Timeout}` |
| Latency measurement | `time.Since(start)` |
| Struct-based data modeling | `type Result struct` |
| CLI argument parsing | `os.Args` |
| Error handling | Go-style `if err != nil` |
