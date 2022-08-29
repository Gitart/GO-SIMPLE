## Find DNS records programmatically

DNS records are mapping files that associate with DNS server whichever IP addresses each domain is associated with, and they handle requests sent to each domain. The net package contains various methods to find general details of DNS records. Let's run some examples, to collect information about the DNS servers and the corresponding records of a target domain:

---

## A and AAAA

The net.LookupIP() function accepts a string(domain-name) and returns a slice of net.IP objects that contains host's IPv4 and IPv6 addresses.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupIP("facebook.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
}
```

The output of above program lists the A records for facebook.com that were returned in IPv4 and IPv6 formats.

```jsx
2a03:2880:f12f:83:face:b00c:0:25de
31.13.79.35
```

---

## Canonical Name (CNAME)

This is the abbreviation for canonical name. CNAMEs are essentially domain and subdomain text aliases to bind traffic. The net.LookupCNAME() function accepts a host-name(m.facebook.com) as a string and returns a single canonical domain name for the given host.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	cname, _ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)
}
```

The CNAME record that was returned for the m.facebook.com domain is shown below:

```jsx
star-mini.c10r.facebook.com.
```

---

## PTR (pointer)

These records provide the reverse binding from addresses to names. PTR records should exactly match the forward maps. The net.LookupAddr() function performs a reverse finding for the address and returns a list of names that map to the given address.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	ptr, _ := net.LookupAddr("6.8.8.8")
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}
}
```

For the given address the above program returns a single reverse record as shown below:

```jsx
tms_server.yuma.army.mil
```

---

## Name Server (NS)

The NS records describe the authorized name servers for the zone. The NS also delegates subdomains to other organizations on zone files. The net.LookupNS() function takes a domain name(facebook.com) as a string and returns DNS-NS records as a slice of NS structs.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	nameserver, _ := net.LookupNS("facebook.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
}
```

```jsx
&{a.ns.facebook.com.} &{b.ns.facebook.com.}
```

---

## MX records

These records identify the servers that can exchange emails. The net.LookupMX() function takes a domain name as a string and returns a slice of MX structs sorted by preference. An MX struct is made up of a Host as a string and Pref as a uint16.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	mxrecords, _ := net.LookupMX("facebook.com")
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}
```

The output list MX record for the domain(facebook.com) followed by preference.

```jsx
msgin.vvv.facebook.com. 10
```

---

## SRV service record

The LookupSRV function tries to resolve an SRV query of the given service, protocol, and domain name. The second parameter is "tcp" or "udp". The returned records are sorted by priority and randomized by weight within a priority.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ncname: %s \n\n", cname)

	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}
```

The output below demonstrates the CNAME return, followed by the SRV record target, port, priority, and weight separated by a colon.

```jsx
cname: _xmpp-server._tcp.golang.org.
```

---

## TXT records

This text record stores information about the SPF that can identify the authorized server to send email on behalf of your organization. The net.LookupTXT() function takes a domain name(facebook.com) as a string and returns a list of DNS TXT records as a slice of strings.

### Example

```jsx
package main

import (
	"fmt"
	"net"
)

func main() {
	txtrecords, _ := net.LookupTXT("facebook.com")

	for _, txt := range txtrecords {
		fmt.Println(txt)
	}
}
```

The single TXT record for gmail.com is shown below:

```jsx
v=spf1 redirect=_spf.facebook.com
```
