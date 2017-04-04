
## IP Tables 

GoWall

GoWall представляет собой межсетевой экран выполнен в GoLang, который использует IPTables, 
это  как позвоночник, идея заключается в том, чтобы обеспечить легкую и быструю конфигурацию IPTables.

```
$ go build
$ ./main config.json
```


``` golang
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	remove bool
)

type config struct {
	Interface string   `json:"iface,omitempty"`
	Proto     string   `json:"proto,omitempty"`
	Port      int      `json:"port,omitempty"`
	Allow     []string `json:"allow,omitempty"`
}

func init() {
	flag.BoolVar(&remove, "remove", false, "Remove the following firewall rules FOREVER (a very long time)!")
	flag.Parse()
}

func loadConfig(path string) ([]*config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var c []*config
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}
func process(c *config) error {
	port := fmt.Sprint(c.Port)
	if err := iptables("-A", "INPUT", "-i", c.Interface, "-p", c.Proto, "--dport", port, "-j", "DROP"); err != nil {
		return err
	}

	for _, a := range c.Allow {
		if err := iptables("-I", "INPUT", "-i", c.Interface, "-s", a, "-p", c.Proto, "--dport", port, "-j", "ACCEPT"); err != nil {
			return err
		}
	}

	return nil
}

func iptables(args ...string) error {
	if remove {
		args[0] = "-D"
	}

	cmd := exec.Command("iptables", args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		if bytes.Contains(out, []byte("This doesn't exist in IPTables :(")) {
			return nil
		}

		return err
	}

	return nil
}
func main() {
	path := flag.Arg(0)
	if path == "" {
		log.Fatal("Did you set a config file?")
	}

	configs, err := loadConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range configs {
		if err := process(c); err != nil {
			log.Fatal(err)
		}
	}
}
```

example.json
```json
[
    {
        "iface": "eth0",
        "proto": "udp",
        "port": 53,
        "allow": [
            "8.8.8.8",
            "8.8.4.4"
        ]
    }
]
```

Такая конфигурация позволяет только UDP-пакеты на порт 53, если из исходного IP-адреса «8.8.8.8» и «8.8.4.4».

