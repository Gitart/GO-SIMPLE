# Find DNS records programmatically

DNS records are mapping files that associate with DNS server whichever IP addresses each domain is associated with, and they handle requests sent to each domain. The net package contains various methods to find general details of DNS records. Let's run some examples, to collect information about the DNS servers and the corresponding records of a target domain:

---

## Go program to find Forward(A) record of a domain

The net.LookupIP() function accepts a string(domain\-name) and returns a slice of net.IP objects that contains host's IPv4 and IPv6 addresses.

package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, \_ := net.LookupIP("facebook.com")
	for \_, ip := range iprecords {
		fmt.Println(ip)
	}
}

The output of above program lists the A records for facebook.com that were returned in IPv4 and IPv6 formats.

C:\\golang\\dns>go run example1.go 2a03:2880:f12f:83:face:b00c:0:25de 31.13.79.35

---

## Go program to find CNAME record of a domain

This is the abbreviation for canonical name. CNAMEs are essentially domain and subdomain text aliases to bind traffic. The net.LookupCNAME() function accepts a host\-name(m.facebook.com) as a string and returns a single canonical domain name for the given host.

package main

import (
	"fmt"
	"net"
)

func main() {
	cname, \_ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)
}

The CNAME record that was returned for the m.facebook.com domain is shown below:

C:\\golang\\dns>go run example2.go star\-mini.c10r.facebook.com.

---

## Go program to find PTR pointer record of a domain

These records provide the reverse binding from addresses to names. PTR records should exactly match the forward maps. The net.LookupAddr() function performs a reverse finding for the address and returns a list of names that map to the given address.

package main

import (
	"fmt"
	"net"
)

func main() {
	ptr, \_ := net.LookupAddr("6.8.8.8")
	for \_, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}
}

For the given address the above program returns a single reverse record as shown below:

C:\\golang\\dns>go run example3.go tms\_server.yuma.army.mil.

---

## Go program to find Name Server (NS) record of a domain

The NS records describe the authorized name servers for the zone. The NS also delegates subdomains to other organizations on zone files. The net.LookupNS() function takes a domain name(facebook.com) as a string and returns DNS\-NS records as a slice of NS structs.

package main

import (
	"fmt"
	"net"
)

func main() {
	nameserver, \_ := net.LookupNS("facebook.com")
	for \_, ns := range nameserver {
		fmt.Println(ns)
	}
}

The NS records that support the domain are shown below:

C:\\golang\\dns>go run example4.go &{a.ns.facebook.com.} &{b.ns.facebook.com.}

---

## Go program to find MX records record of a domain

These records identify the servers that can exchange emails. The net.LookupMX() function takes a domain name as a string and returns a slice of MX structs sorted by preference. An MX struct is made up of a Host as a string and Pref as a uint16.

package main

import (
	"fmt"
	"net"
)

func main() {
	mxrecords, \_ := net.LookupMX("facebook.com")
	for \_, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}

The output list MX record for the domain(facebook.com) followed by preference.

C:\\golang\\dns>go run example5.go msgin.vvv.facebook.com. 10

---

## Go program to find SRV service record of a domain

The LookupSRV function tries to resolve an SRV query of the given service, protocol, and domain name. The second parameter is "tcp" or "udp". The returned records are sorted by priority and randomized by weight within a priority.

package main

import (
	"fmt"
	"net"
)

func main() {
	cname, srvs, err := net.LookupSRV("xmpp\-server", "tcp", "golang.org")
	if err !\= nil {
		panic(err)
	}

	fmt.Printf("\\ncname: %s \\n\\n", cname)

	for \_, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}

The output below demonstrates the CNAME return, followed by the SRV record target, port, priority, and weight separated by a colon.

C:\\golang\\dns>go run example6.go cname: \_xmpp\-server.\_tcp.golang.org.

---

## Go program to find TXT records of a domain

This text record stores information about the SPF that can identify the authorized server to send email on behalf of your organization. The net.LookupTXT() function takes a domain name(facebook.com) as a string and returns a list of DNS TXT records as a slice of strings.

package main

import (
	"fmt"
	"net"
)

func main() {
	txtrecords, \_ := net.LookupTXT("facebook.com")

	for \_, txt := range txtrecords {
		fmt.Println(txt)
	}
}

The single TXT record for gmail.com is shown below.

C:\\golang\\dns>go run example7.go v=spf1 redirect=\_spf.facebook.com
