package randomclb

import (
	"fmt"
	"log"
	"testing"

	"github.com/pratikju/dns-clb-go/dns"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

//strconv.FormatInt(int64(srv.Port), 10)
func TestRandomLookup(t *testing.T) {
	// given
	srvName := "foo.service.fligl.io"
	lib := dns.NewLookupLib("8.8.8.8:53")
	c := NewRandomClb(lib)

	// when
	address, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}

	if address.Port == 8001 && address.Address == "0.1.2.3" {
		return
	} else if address.Port == 8002 && address.Address == "4.5.6.7" {
		return
	} else {
		t.Errorf("port '%d' not expected with address: '%s'", address.Port, address.Address)
	}
}
