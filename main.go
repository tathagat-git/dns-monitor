package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

// Result holds the health check output for one domain
type Result struct {
	Domain      string
	IPs         []string
	StatusCode  int
	Latency     time.Duration
	Error       string
}

// checkDNS resolves a domain to its IP addresses
func checkDNS(domain string) ([]string, error) {
	ips, err := net.LookupHost(domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

// checkHTTP makes an HTTPS GET request and returns status code + latency
func checkHTTP(domain string) (int, time.Duration, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get("https://" + domain)
	latency := time.Since(start)

	if err != nil {
		return 0, latency, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, latency, nil
}

// monitor runs DNS + HTTP check for a single domain
func monitor(domain string) Result {
	result := Result{Domain: domain}

	// DNS check
	ips, err := checkDNS(domain)
	if err != nil {
		result.Error = fmt.Sprintf("DNS failed: %v", err)
		return result
	}
	result.IPs = ips

	// HTTP check
	status, latency, err := checkHTTP(domain)
	if err != nil {
		result.Error = fmt.Sprintf("HTTP failed: %v", err)
		return result
	}
	result.StatusCode = status
	result.Latency = latency

	return result
}

// printReport prints a clean terminal report
func printReport(results []Result) {
	fmt.Println("\n========================================")
	fmt.Println("   DNS & HTTP Health Monitor Report")
	fmt.Println("========================================")

	for _, r := range results {
		fmt.Printf("\nDomain  : %s\n", r.Domain)

		if r.Error != "" {
			fmt.Printf("Status  : ❌ FAILED\n")
			fmt.Printf("Error   : %s\n", r.Error)
			continue
		}

		fmt.Printf("IPs     : %v\n", r.IPs)
		fmt.Printf("HTTP    : %d\n", r.StatusCode)
		fmt.Printf("Latency : %v\n", r.Latency.Round(time.Millisecond))

		if r.StatusCode >= 200 && r.StatusCode < 300 {
			fmt.Printf("Health  : ✅ OK\n")
		} else {
			fmt.Printf("Health  : ⚠️  WARNING (non-2xx response)\n")
		}
	}

	fmt.Println("\n========================================\n")
}

func main() {
	// Default domains if none are passed as arguments
	domains := []string{
		"google.com",
		"github.com",
		"cloudflare.com",
	}

	// Allow custom domains via CLI args: ./dns-monitor example.com akamai.com
	if len(os.Args) > 1 {
		domains = os.Args[1:]
	}

	fmt.Printf("Checking %d domain(s)...\n", len(domains))

	var results []Result
	for _, domain := range domains {
		fmt.Printf("  → %s\n", domain)
		result := monitor(domain)
		results = append(results, result)
	}

	printReport(results)
}
