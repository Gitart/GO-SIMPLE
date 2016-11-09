// This systemd runs iptables-restore on boot:
[Unit]
Description=Packet Filtering Framework
DefaultDependencies=no
After=systemd-sysctl.service
Before=sysinit.target

[Service]
Type=oneshot
ExecStart=/usr/sbin/iptables-restore /opt/docker/scripts/iptables/iptables.rules
ExecReload=/usr/sbin/iptables-restore /opt/docker/scripts/iptables/iptables.rules
ExecStop=/usr/sbin/iptables --flush
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target


// This is my iptables.rules file
# Adapted from here: http://wiki.centos.org/HowTos/OS_Protection
*filter

:INPUT DROP [0:0]
:FORWARD DROP [0:0]
:OUTPUT ACCEPT [0:0]
:RH-Firewall-1-INPUT - [0:0]

-A INPUT -j RH-Firewall-1-INPUT
-A FORWARD -j RH-Firewall-1-INPUT
-A RH-Firewall-1-INPUT -i lo -j ACCEPT
-A RH-Firewall-1-INPUT -p icmp --icmp-type echo-reply -j ACCEPT
-A RH-Firewall-1-INPUT -p icmp --icmp-type destination-unreachable -j ACCEPT
-A RH-Firewall-1-INPUT -p icmp --icmp-type time-exceeded -j ACCEPT

# Block Spoofing IP Addresses
-A INPUT -i eth0 -s 10.0.0.0/8 -j DROP
-A INPUT -i eth0 -s 172.16.0.0/12 -j DROP
-A INPUT -i eth0 -s 192.168.0.0/16 -j DROP
-A INPUT -i eth0 -s 224.0.0.0/4 -j DROP
-A INPUT -i eth0 -s 240.0.0.0/5 -j DROP
-A INPUT -i eth0 -d 127.0.0.0/8 -j DROP

# Accept Pings
-A RH-Firewall-1-INPUT -p icmp --icmp-type echo-request -j ACCEPT

# Accept any established connections
-A RH-Firewall-1-INPUT -m conntrack --ctstate  ESTABLISHED,RELATED -j ACCEPT

# Accept ssh, http, https - add other tcp traffic ports here
-A RH-Firewall-1-INPUT -m conntrack --ctstate NEW -m multiport -p tcp --dports 22,80,443 -j ACCEPT

#Log and drop everything else
-A RH-Firewall-1-INPUT -j LOG
-A RH-Firewall-1-INPUT -j REJECT --reject-with icmp-host-prohibited

COMMIT

// After the machine has rebooted and a couple of docker containers also started, this is the output of iptables -L
Chain INPUT (policy DROP)
target     prot opt source               destination
RH-Firewall-1-INPUT  all  --  anywhere             anywhere
DROP       all  --  10.0.0.0/8           anywhere
DROP       all  --  172.16.0.0/12        anywhere
DROP       all  --  192.168.0.0/16       anywhere
DROP       all  --  base-address.mcast.net/4  anywhere
DROP       all  --  240.0.0.0/5          anywhere
DROP       all  --  anywhere             loopback/8

Chain FORWARD (policy DROP)
target     prot opt source               destination
ACCEPT     udp  --  anywhere             172.17.0.3           udp dpt:domain
ACCEPT     tcp  --  anywhere             172.17.0.2           tcp dpt:5000
ACCEPT     all  --  anywhere             anywhere             ctstate RELATED,ESTABLISHED
ACCEPT     all  --  anywhere             anywhere
ACCEPT     all  --  anywhere             anywhere
RH-Firewall-1-INPUT  all  --  anywhere             anywhere

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination

Chain RH-Firewall-1-INPUT (2 references)
target     prot opt source               destination
ACCEPT     all  --  anywhere             anywhere
ACCEPT     icmp --  anywhere             anywhere             icmp echo-reply
ACCEPT     icmp --  anywhere             anywhere             icmp destination-unreachable
ACCEPT     icmp --  anywhere             anywhere             icmp time-exceeded
ACCEPT     icmp --  anywhere             anywhere             icmp echo-request
ACCEPT     all  --  anywhere             anywhere             ctstate RELATED,ESTABLISHED
ACCEPT     tcp  --  anywhere             anywhere             ctstate NEW multiport dports ssh,http,https
LOG        all  --  anywhere             anywhere             LOG level warning
REJECT     all  --  anywhere             anywhere             reject-with icmp-host-prohibited
