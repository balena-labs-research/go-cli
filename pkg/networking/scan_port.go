package networking

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(ip string, port int, timeout time.Duration) error {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		return err
	}

	conn.Close()
	return nil
}
