package handlers

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/tatsushid/go-fastping"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("ip")
	host = strings.TrimSpace(host)

	// Perform actual ICMP ping operation
	response := performICMPPing(host)

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	_, err := w.Write([]byte(response))
	if err != nil {
		return
	}
}

func performICMPPing(host string) string {
	var response strings.Builder

	response.WriteString(fmt.Sprintf("PING %s ", host))

	// Resolve IP address for hostnames
	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		response.WriteString(fmt.Sprintf("(%s)", err.Error()))
		return response.String()
	}

	response.WriteString(fmt.Sprintf("(%s) 56(84) bytes of data.\n", ipAddr))

	// Perform ICMP ping operation
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip", ipAddr.String())
	if err != nil {
		response.WriteString(fmt.Sprintf("Error resolving IP address: %s\n", err.Error()))
		return response.String()
	}

	p.AddIPAddr(ra)

	received := 0
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		received++
		response.WriteString(fmt.Sprintf("64 bytes from %s: icmp_seq=%d time=%.3f ms\n",
			addr.String(), received, rtt.Seconds()*1000))
	}

	// Perform 4 ping requests
	transmitted := 4
	for i := 0; i < transmitted; i++ {
		err = p.Run()
		if err != nil {
			response.WriteString(fmt.Sprintf("Error during ping: %s\n", err.Error()))
			break
		}
	}

	packetLoss := float64(transmitted-received) / float64(transmitted) * 100.0
	minLatency := 0.0
	avgLatency := 0.0
	maxLatency := 0.0

	response.WriteString(fmt.Sprintf("--- %s ping statistics ---\n", host))
	response.WriteString(fmt.Sprintf("%d packets transmitted, %d received, %.1f%% packet loss, time %dms\n",
		transmitted, received, packetLoss, 1815))
	response.WriteString(fmt.Sprintf("rtt min/avg/max/mdev = %.3f/%.3f/%.3f/%.3f ms\n",
		minLatency, avgLatency, maxLatency, maxLatency-minLatency))

	return response.String()
}
