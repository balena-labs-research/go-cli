package networking

import (
	"testing"
)

func TestArpScan(t *testing.T) {
	_, err := ArpScan()

	if err != nil {
		t.Error(err)
	}

}
