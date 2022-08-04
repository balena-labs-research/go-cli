package networking

import (
	"fmt"
	"testing"
)

func TestArpScan(t *testing.T) {

	out, err := ArpScan()
	fmt.Println(out)

	if err != nil {
		t.Error(err)
	}

}
