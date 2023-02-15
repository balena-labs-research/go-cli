package networking

import (
	"context"
	"encoding/binary"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func createHostRange(netw string) []string {
	_, ipv4Net, err := net.ParseCIDR(netw)
	if err != nil {
		log.Fatal(err)
	}

	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	finish := (start & mask) | (mask ^ 0xffffffff)

	var hosts []string
	for i := start + 1; i <= finish-1; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}

	return hosts
}

func getLocalRange() string {
	defaultAddr := "192.168.1.0/24"
	addrs, err := net.InterfaceAddrs()

	// If no interface address available, return the default address
	if err == nil {
		for _, address := range addrs {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					split := strings.Split(ipnet.IP.String(), ".")
					return split[0] + "." + split[1] + "." + split[2] + ".0/24"
				}
			}
		}
	}

	// Return the default address if no address is found
	log.Print("No interface address found. Using default address: " + defaultAddr)
	return defaultAddr
}

// ScanRange scans every address on a CIDR for open ports
func LookupAddresses(ipRange string) []string {

	var localRange string

	if ipRange != "" {
		localRange = ipRange
	} else {
		localRange = getLocalRange()
	}

	hostRange := createHostRange(localRange)

	var wg sync.WaitGroup
	var results []string
	ch := make(chan string)
	for _, ip := range hostRange {
		// Store the output in a channel
		wg.Add(1)
		go lookupIP(ip, ch, &wg)
	}

	// Collect results:
	go func() {
		for v := range ch {
			results = append(results, v)
		}
	}()

	wg.Wait()
	close(ch)

	return results
}

func lookupIP(ip string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	const timeout = 1000 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel() // important to avoid a resource leak

	var r net.Resolver
	names, err := r.LookupAddr(ctx, ip)

	if err == nil && len(names) > 0 {

		ch <- strings.TrimSuffix(names[0], ".")
	}

}
