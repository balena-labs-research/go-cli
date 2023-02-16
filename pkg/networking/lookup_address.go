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

func createHostRange(netw string) ([]string, error) {
	_, ipv4Net, err := net.ParseCIDR(netw)
	if err != nil {
		return nil, err
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

	return hosts, err
}

func getLocalRange() string {
	defaultAddr := "192.168.1.0/24"
	addrs, err := net.InterfaceAddrs()

	// If no interface address available, return the default address
	if err == nil {
		for _, address := range addrs {
			// check the address type and if it is not a loopback then display it
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

func LookupAddresses(ipRange string) ([]string, error) {
	var localRange string

	if ipRange != "" {
		localRange = ipRange
	} else {
		localRange = getLocalRange()
	}

	hostRange, err := createHostRange(localRange)

	if err != nil {
		return nil, err
	}

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

	return results, err
}

func lookupIP(ip string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	const timeout = 3000 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel() // Important to avoid a resource leak

	var r net.Resolver
	names, err := r.LookupAddr(ctx, ip)

	if err == nil && len(names) > 0 {

		ch <- strings.TrimSuffix(names[0], ".")
	}

}
