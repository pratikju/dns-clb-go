# DNS Client Load Balancer for Go

Selects a `SRV` record answer according to specified load balancer algorithm, then resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port    uint16
	}


## Example:

	// uses dns server configured in /etc/resolv.conf
	srvName := "my-svc.service.consul"
	c := clb.New()
	address, err := c.GetAddress(srvName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

### or configure explicitely

	srvName := "my-svc.service.consul"
	c := clb.NewClb("127.0.0.1", "8600", clb.RoundRobin)
	address, err := c.GetAddress(srvName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

### or use an `AddressProvider`

	// uses dns server configured in /etc/resolv.conf
	ap := NewAddressProvider("my-svc.service.consul")
	address, err := ap.GetAddress()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001


## Development
tests are run against some fixture dns entries I set up on fligl.io (`dig foo.service.fligl.io SRV`).


- `make deps` install deps
- `make test` run all tests
